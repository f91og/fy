package test

import (
	"testing"

	"github.com/f91og/fy/src/util"
)

// https://geektutu.com/post/quick-go-test.html#2-%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E4%BE%8B%E5%AD%90
func TestCheckLangType(t *testing.T) {
	var (
		text1 = "english"
		text2 = "中文"
		// text3 = "日本語"
		text4 = "しかし"
	)

	if res := util.CheckLangType(text1); res != "en" {
		t.Errorf("\"english\" expected be en, but %s got", res)
	}
	if res := util.CheckLangType(text2); res != "zh" {
		t.Errorf("\"中文\" expected be zh, but %s got", res)
	}
	if res := util.CheckLangType(text4); res != "ja" {
		t.Errorf("\"しかし\" expected be ja, but %s got", res)
	}
}
