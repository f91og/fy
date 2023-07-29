package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/f91og/fy/src/util"
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
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user info when init dict: %w", err)
	}
	homeDir := usr.HomeDir
	os.MkdirAll(fmt.Sprintf("%s/.fy", homeDir), 0755)

	d := &Dict{
		LangType:        langType,
		WordRecords:     make(map[string]WordRecord),
		SentenceRecords: make(map[string]SentenceRecord),
	}

	if langType == EN {
		d.FilePath = fmt.Sprintf("%s/.fy/dict.en", homeDir)
	} else if langType == JA {
		d.FilePath = fmt.Sprintf("%s/.fy/dict.ja", homeDir)
	} else if langType == ZH {
		d.FilePath = fmt.Sprintf("%s/.fy/dict.zh", homeDir)
	}

	file, err := os.OpenFile(d.FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s when init dict: %w", d.FilePath, err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "|")
		key := strs[0]
		if len(strs) > 3 {
			d.WordRecords[key] = WordRecord{strs[0], strs[1], strs[2], strs[3]}
		} else {
			d.SentenceRecords[key] = SentenceRecord{strs[0], strs[1], strs[2]}
		}
	}

	return d, nil
}

func (d *Dict) AddRecord(record interface{}) error {
	file, err := os.OpenFile(d.FilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	var recordStr string
	switch r := record.(type) {
	case WordRecord:
		d.WordRecords[r.Word] = r
		recordStr = fmt.Sprintf("%s|%s|%s|%s\n", r.Word, r.Pronunciation, r.Translation, r.Example)
	case SentenceRecord:
		d.SentenceRecords[r.Sentence] = r
		recordStr = fmt.Sprintf("%s|%s|%s\n", r.Sentence, r.Translation1, r.Translation2)
	default:
		return fmt.Errorf("unsupported record type")
	}

	_, err = file.WriteString(recordStr)
	if err != nil {
		return fmt.Errorf("add record failed: %w", err)
	}

	return nil
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
	_, ok1 := d.WordRecords[query]
	_, ok2 := d.SentenceRecords[query]

	if ok1 || ok2 {
		delete(d.WordRecords, query)
		delete(d.SentenceRecords, query)
		if err := util.DeleteLine(d.FilePath, query); err != nil {
			return fmt.Errorf("delete record failed: %w", err)
		}
	}

	return nil
}

var (
	yellow = color.New(color.Bold, color.FgHiYellow).SprintFunc()
	cyan   = color.New(color.Bold, color.FgCyan).SprintFunc()
	white  = color.New(color.Bold, color.FgWhite).SprintFunc()
)

func (r WordRecord) ColorPrint() {
	fmt.Printf("%s; %s; %s; %s\n", yellow(r.Word), cyan(r.Pronunciation), white(r.Translation), yellow(r.Example))
}

func (r SentenceRecord) ColorPrint() {
	fmt.Printf("%s; %s; %s\n", yellow(r.Sentence), cyan(r.Translation1), white(r.Translation2))
}
