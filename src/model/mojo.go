package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Mojo struct {
	Engine
}

type SearchResult struct {
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
}

type Response struct {
	Result struct {
		Code    int `json:"code"`
		Results struct {
			SearchAll struct {
				Code   int `json:"code"`
				Result struct {
					Word struct {
						SearchResult []SearchResult `json:"searchResult"`
					} `json:"word"`
					Grammar struct {
						SearchResult []SearchResult `json:"searchResult"`
					} `json:"grammar"`
					Example struct {
						SearchResult []SearchResult `json:"searchResult"`
					} `json:"example"`
				} `json:"result"`
			} `json:"search-all"`
		} `json:"results"`
	} `json:"result"`
}

func (m *Mojo) Translate(query string) (Record, error) {
	rawStr := fmt.Sprintf(`{"functions":[{"name":"search-all","params":{"text":"%s","types":[102,106,103]}}],"_ApplicationId":"%s"}`, query, m.Params["appId"])
	reqBody := strings.NewReader(rawStr)

	resp, err := http.Post(m.ApiUrl, "text/plain", reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create post request in mojo translator: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get mojo translator result: %w", err)
	}

	data := &Response{}
	if err = json.Unmarshal(body, data); err != nil {
		return nil, fmt.Errorf("failed to parse mojo translator result: %w", err)
	}

	if data.Result.Code != 200 || data.Result.Results.SearchAll.Code != 200 {
		return nil, fmt.Errorf("Mojo translate failed")
	}
	// fmt.Println(data)

	words := data.Result.Results.SearchAll.Result.Word.SearchResult
	examples := data.Result.Results.SearchAll.Result.Example.SearchResult
	// res1 := fmt.Sprintf("%s:%s\teg: %s:%s", words[0].Title, words[0].Excerpt, examples[0].Title, examples[0].Excerpt)

	wordWithPron := strings.Split(words[0].Title, "|")
	word := strings.TrimSpace(wordWithPron[0])
	var pron string
	if len(wordWithPron) >= 2 {
		pron = strings.TrimSpace(wordWithPron[1])
	}

	trans := strings.TrimSpace(words[0].Excerpt)
	example := fmt.Sprintf("%s/%s", examples[0].Title, examples[0].Excerpt)

	return WordRecord{word, pron, trans, example}, nil
}
