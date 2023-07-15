package test

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/f91og/fy/src/util"
)

// https://geektutu.com/post/quick-go-test.html#2-%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E4%BE%8B%E5%AD%90
func TestDeleteLine(t *testing.T) {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	filename := fmt.Sprintf("%s/.fy/data.txt", homeDir)

	err := util.DeleteLine(filename, "Broadly")
	if err != nil {
		t.Errorf(err.Error())
	}
}
