package util

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func DeleteLine(filename string, startStr string) error {
	originalFile, err := os.Open(filename)
	if err != nil {
		// 处理打开原始文件的错误
		panic(err)
	}
	defer originalFile.Close()

	usr, _ := user.Current()
	tempDir := fmt.Sprintf("%s/.fy", usr.HomeDir)
	tempFile, err := os.CreateTemp(tempDir, "tempfile")
	if err != nil {
		// 处理创建临时文件的错误
		panic(err)
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(originalFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, strings.TrimSpace(startStr)) {
			continue
		}
		_, err := tempFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := tempFile.Sync(); err != nil {
		return err
	}

	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		return err
	}

	return nil
}
