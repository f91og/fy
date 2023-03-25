package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Google struct {
	Engine
}

func (g *Google) Translate(text *Text) (string, string, error) {
	var tls1, tls2 string
	if text.LangType == ZH {
		tls1, tls2 = "en", "ja"
	} else if text.LangType == EN {
		tls1, tls2 = "zh", "ja"
	} else if text.LangType == JA {
		tls1, tls2 = "zh", "en"
	}

	params := url.Values{}
	Url, err := url.Parse(g.ApiUrl)
	for key, value := range g.Params {
		params.Set(key, value)
	}
	params.Set("sl", text.LangType)
	params.Set("tl", tls1)
	params.Set("q", text.Value)
	Url.RawQuery = params.Encode()

	resp, err := http.Get(Url.String())
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("request failed, google translate may limit your access")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var data []interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", "", err
	}

	res1 := data[0].([]interface{})[0].([]interface{})[0].(string)

	params.Set("tl", tls2)
	Url.RawQuery = params.Encode()
	resp2, err := http.Get(Url.String())
	if err != nil {
		return res1, "", err
	}
	defer resp2.Body.Close()
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return res1, "", err
	}
	var data2 []interface{}
	if err = json.Unmarshal(body2, &data2); err != nil {
		return res1, "", err
	}
	res2 := data2[0].([]interface{})[0].([]interface{})[0].(string)

	return res1, res2, nil
}
