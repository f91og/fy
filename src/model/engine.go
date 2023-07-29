package model

import (
	"fmt"
)

const (
	GOOGLE    string = "google"
	MOJO      string = "mojo"
	BING      string = "bing"
	CAMBRIDGE string = "cambridge"

	EN string = "en"
	JA string = "ja"
	ZH string = "zh"
)

type Translator interface {
	Translate(query string) (Record, error)
}

type Engine struct {
	Params       map[string]string
	SupportModel string
	ApiUrl       string
}

func Translate(query string, langType string, sentence bool, trans string) (Record, error) {
	translator, err := makeTranslator(query, langType, sentence, trans)
	if err != nil {
		return nil, fmt.Errorf("make translator failed: %w", err)
	}

	if record, err := translator.Translate(query); err != nil {
		return nil, fmt.Errorf("translate failed: %w", err)
	} else {
		return record, nil
	}
}

func makeTranslator(query string, langType string, sentence bool, trans string) (Translator, error) {
	translators := make(map[string]Translator)

	g := &Google{}
	g.ApiUrl = "http://translate.google.com/translate_a/single"
	g.SupportModel = "sentence"
	g.Params = map[string]string{
		"client": "at",
		"dt":     "t",
	}
	translators[GOOGLE] = g
	if sentence {
		return g, nil
	}

	c := &Cambridge{}
	c.ApiUrl = "https://dictionary.cambridge.org/zhs/%E8%AF%8D%E5%85%B8/%E8%8B%B1%E8%AF%AD-%E6%B1%89%E8%AF%AD-%E7%B9%81%E4%BD%93"
	translators[CAMBRIDGE] = c

	m := &Mojo{}
	m.ApiUrl = "https://api.mojidict.com/parse/functions/union-api"
	m.Params = map[string]string{
		"appId": "E62VyFVLMiW7kvbtVq3p",
	}
	translators[MOJO] = m

	// return translator if specified, else return default translator
	if trans != "" {
		if translator, ok := translators[trans]; !ok {
			return nil, fmt.Errorf("cannot find translator")
		} else {
			return translator, nil
		}
	}

	if langType == EN {
		return translators[CAMBRIDGE], nil
	} else {
		return translators[MOJO], nil
	}
}
