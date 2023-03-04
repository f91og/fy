package engine

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Mojo struct {
	Engine
}

func (m *Mojo) Translate(text, sl string) (string, string, error) {
	client := &http.Client{}
	m.Params[""] = ""
	data := strings.NewReader(`{"functions":[{"name":"search-all","params":{"text":"訳す","types":[102,106,103]}},{"name":"mojitest-examV2-searchQuestion-v2","params":{"text":"訳す","limit":1,"page":1}}],"_ApplicationId":"E62VyFVLMiW7kvbtVq3p"}`)
	req, err := http.NewRequest("POST", m.ApiUrl, data)
	if err != nil {
		return "", "", nil
	}

	req.Header.Set("content-type", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	return "", "", nil
}
