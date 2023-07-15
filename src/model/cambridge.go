package model

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Cambridge struct {
	Engine
}

func (c *Cambridge) Translate(query string) (Record, error) {
	url := fmt.Sprintf("%s/%s", c.ApiUrl, query)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to create request in cambridge translator: %w", err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("cambridge result document parse failed %w", err)
	}

	prons := doc.Find(".pron.dpron .ipa.dipa").First().Text()
	trans := doc.Find("span.trans.dtrans.dtrans-se")
	egs := doc.Find("span.eg.deg")
	wordTrans := strings.TrimSpace(trans.Eq(0).Text())
	eg := strings.TrimSpace(egs.Eq(0).Text())
	egTrans := strings.TrimSpace(trans.Eq(1).Text())
	// replacer := strings.NewReplacer("\n", "", "\t", "")
	// tmp := strings.Split(replacer.Replace(examTrans), ".")
	// example := fmt.Sprintf("%s/%s", tmp[0], strings.TrimLeft(tmp[1], " "))

	return WordRecord{query, prons, wordTrans, fmt.Sprintf("%s/%s", eg, egTrans)}, nil
}
