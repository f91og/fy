package engine

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"k8s.io/klog"
)

type GoogleTranslator struct {
	Engine
}

func (g *GoogleTranslator) Translate(text, sl string) (string, string, error) {
	if text == "" {
		klog.Fatal("no translate text")
	}
	var tls1, tls2 string
	if sl == "zh" {
		tls1, tls2 = "en", "ja"
	} else if sl == "en" {
		tls1, tls2 = "zh", "ja"
	} else if sl == "ja" {
		tls1, tls2 = "zh", "en"
	}

	params := url.Values{}
	Url, err := url.Parse(g.ApiUrl)
	for key, value := range g.Params {
		params.Set(key, value)
	}
	params.Set("sl", sl)
	params.Set("tl", tls1)
	params.Set("q", text)
	Url.RawQuery = params.Encode()

	resp, err := http.Get(Url.String())
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		klog.Fatal("request failed, google translate may limit your access")
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
