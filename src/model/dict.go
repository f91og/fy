package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Record interface {
	ColorPrint()
}

type WordRecord struct {
	Word          string
	Pronunciation string
	Translation   string
	Example       string
}

type SentenceRecord struct {
	Sentence     string
	Translation1 string
	Translation2 string
}

type Dict struct {
	LangType        string
	FilePath        string
	WordRecords     map[string]WordRecord
	SentenceRecords map[string]SentenceRecord
}

func InitDict(langType string) (*Dict, error) {
	d := &Dict{}
	if langType == EN {
		d.FilePath = "/Users/xue.a.yu/.fy/dict.en"
	} else if langType == JA {
		d.FilePath = "/Users/xue.a.yu/.fy/dict.ja"
	} else if langType == ZH {
		d.FilePath = "/Users/xue.a.yu/.fy/dict.zh"
	}
	file, err := os.Open(d.FilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "|")
		if len(strs) > 3 {
			d.WordRecords[strs[0]] = WordRecord{strs[0], strs[1], strs[2], strs[3]}
		} else {
			d.SentenceRecords[strs[0]] = SentenceRecord{strs[0], strs[1], strs[2]}
		}
	}

	return d, nil
}

func (d *Dict) AddRecord(record interface{}) error {
	os.MkdirAll(d.FilePath, 0755)
	file, err := os.OpenFile(d.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	var recordStr string
	switch r := record.(type) {
	case WordRecord:
		d.WordRecords[r.Word] = r
		recordStr = fmt.Sprintf("%s | %s | %s | %s\n", r.Word, r.Pronunciation, r.Translation, r.Example)
	case SentenceRecord:
		d.SentenceRecords[r.Sentence] = r
		recordStr = fmt.Sprintf("%s | %s | %s\n", r.Sentence, r.Translation1, r.Translation2)
	default:
		return fmt.Errorf("unsupported record type")
	}

	_, err = file.WriteString(recordStr)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dict) GetRecordByQuery(query string) (Record, error) {

	return nil, fmt.Errorf("cannot find query record")
}

func (d *Dict) GetRecordByLine(line int) (Record, error) {
	rand.Seed(time.Now().UnixNano())

	wordLen, sentenceLen := len(d.WordRecords), len(d.SentenceRecords)
	index := rand.Intn(wordLen + sentenceLen)
	i := 0
	if index >= wordLen {
		index = index - wordLen
		for _, record := range d.SentenceRecords {
			if i == index {
				return &record, nil
			}
		}
	} else {
		for _, record := range d.WordRecords {
			if i == index {
				return &record, nil
			}
		}
	}

	return nil, fmt.Errorf("cannot get random record")
}

func (d *Dict) DeleteRecordByQuery(query string) error {

	return nil
}

var (
	green = color.New(color.Bold, color.FgGreen).SprintFunc()
	cyan  = color.New(color.Bold, color.FgCyan).SprintFunc()
	white = color.New(color.Bold, color.FgWhite).SprintFunc()
)

func (r *WordRecord) ColorPrint() {
	fmt.Printf("%s; %s; %s; %s\n", green(r.Word), cyan(r.Pronunciation), white(r.Translation), green(r.Example))
}

func (r *SentenceRecord) ColorPrint() {
	fmt.Printf("%s; %s; %s\n", green(r.Sentence), cyan(r.Translation1), white(r.Translation2))
}
