package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetRecord(value string) (string, string, error) {
	file, err := os.Open("/Users/xue.a.yu/.fy/data.txt")
	if err != nil {
		return "", "", err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "|")

		var res1, res2 string

		if value == strs[0] {
			res1 = strs[1]
			if len(strs) > 1 {
				res2 = strs[2]
			}
			return res1, res2, nil
		}
	}

	return "", "", nil
}

func DeleteRecord(key string) error {
	file, err := os.OpenFile("/Users/xue.a.yu/.fy/data.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	records := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, strings.TrimSpace(key)) {
			records = append(records, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil { // Truncate方法把文件截断为0，重新将文件指针定位到文件开头
		return err
	}

	_, err = file.Seek(0, 0) // 将文件的读写位置设置为文件的开头位置，因为在读取或者写入文件之前，需要将文件的读写位置设置到文件开头
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	for _, record := range records {
		if _, err := writer.WriteString(record + "\n"); err != nil {
			writer.Flush()
			return err
		}
	}
	return writer.Flush()
}

func WriteRecord(key, translation1, translation2 string) error {
	os.MkdirAll("/Users/xue.a.yu/.fy", 0755)
	file, err := os.OpenFile("/Users/xue.a.yu/.fy/data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	record := fmt.Sprintf("%s|%s|%s\n", key, translation1, translation2)
	_, err = file.WriteString(record)
	if err != nil {
		return err
	}

	return nil
}

func GetFileLines() (int, error) {
	file, err := os.Open("/Users/xue.a.yu/.fy/data.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	fd := bufio.NewReader(file)
	count := 0
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++

	}
	return count, nil
}

func GetRecordByLineNumber(line int) (string, error) {
	file, err := os.Open("/Users/xue.a.yu/.fy/data.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
		if count == line {
			line := scanner.Text()
			return line, nil
		}
	}
	return "", fmt.Errorf("cannot randomly show record in line %d", line)
}
