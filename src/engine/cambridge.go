package engine

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Cambridge struct {
	Engine
}

func (c *Cambridge) Translate(text *Text) (string, string, error) {
	url := fmt.Sprintf("%s/%s", c.ApiUrl, text.Value)

	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", "", err
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
	return fmt.Sprintf("%s\t%s\t%s", prons, wordTrans, fmt.Sprintf("%s/%s", eg, egTrans)), "", nil
}
