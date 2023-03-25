package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseCambridgeHtml(t *testing.T) {
	url := "https://dictionary.cambridge.org/zhs/%E8%AF%8D%E5%85%B8/%E8%8B%B1%E8%AF%AD-%E6%B1%89%E8%AF%AD-%E7%B9%81%E4%BD%93/decommission"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	prons := doc.Find(".pron.dpron .ipa.dipa").First().Text()
	trans := doc.Find(".trans.dtrans")
	wordTrans, examTrans := trans.First().Text(), trans.Next().Text()

	replacer := strings.NewReplacer("\n", "", "\t", "")
	tmp := strings.Split(replacer.Replace(examTrans), ".")
	example := fmt.Sprintf("%s/%s", tmp[0], strings.TrimLeft(tmp[1], " "))
	t.Log("$", prons, "$", wordTrans, example)
}

func TestParseBingHtml(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://cn.bing.com/dict/search?q=broken&qs=n&form=Z9LH5&sp=-1&lq=0&pq=broken&sc=7-6&sk=&cvid=9583152EFAF74A1284E30EB81F5C07FC&ghsh=0&ghacc=0&ghpl=", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)

	//	url := fmt.Sprintf("https://cn.bing.com/dict/search?q=%s", "broken")
	//resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	t.Log(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	trans := doc.Find(".def.b_regtxt").First().Text()
	prons := doc.Find(".hd_pr.b_primtxt").First().Text()

	//hd_pr b_primtxt

	t.Log(prons, trans)
}
