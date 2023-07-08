package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/f91og/fy/src/util"
)

type Google struct {
	Engine
}

func (g *Google) Translate(query string) (Record, error) {
	var tls1, tls2 string

	langType := util.CheckLangType(query)
	if langType == ZH {
		tls1, tls2 = "en", "ja"
	} else if langType == EN {
		tls1, tls2 = "zh", "ja"
	} else if langType == JA {
		tls1, tls2 = "zh", "en"
	}

	params := url.Values{}
	Url, _ := url.Parse(g.ApiUrl)
	for key, value := range g.Params {
		params.Set(key, value)
	}
	params.Set("sl", langType)
	params.Set("tl", tls1)
	params.Set("q", query)
	Url.RawQuery = params.Encode()

	resp, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed, google translate may limit your access")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	res1 := data[0].([]interface{})[0].([]interface{})[0].(string)

	params.Set("tl", tls2)
	Url.RawQuery = params.Encode()
	resp2, err := http.Get(Url.String())
	if err != nil {
		return &SentenceRecord{query, res1, ""}, err
	}
	defer resp2.Body.Close()
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return &SentenceRecord{query, res1, ""}, err
	}
	var data2 []interface{}
	if err = json.Unmarshal(body2, &data2); err != nil {
		return &SentenceRecord{query, res1, ""}, err
	}
	res2 := data2[0].([]interface{})[0].([]interface{})[0].(string)

	return &SentenceRecord{query, res1, res2}, nil
}
