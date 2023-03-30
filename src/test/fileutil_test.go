package test

import (
	"testing"

	"github.com/f91og/fy/src/util"
)

// https://geektutu.com/post/quick-go-test.html#2-%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E4%BE%8B%E5%AD%90
func TestWriteReadRecord(t *testing.T) {
	var (
		key          = "但是"
		translation1 = "but"
		// text3 = "日本語"
		translation2 = "しかし"
	)

	err := util.WriteRecord(key, translation1, translation2)

	if err != nil {
		t.Fatalf(err.Error())
	}

	res1, res2, err := util.GetRecord(key)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res1 != translation1 || res2 != translation2 {
		t.Errorf("Write record: %s,%s,%s, read result:%s,%s,%s", key, translation1, translation2, key, res1, res2)
	}
}

func TestGetRecordByLineNumber(t *testing.T) {
	line := 1
	res, err := util.GetRecordByLineNumber(line)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDeleteRecord(t *testing.T) {
	err := util.DeleteRecord("deliver")
	if err != nil {
		t.Fatal(err)
	}
}
