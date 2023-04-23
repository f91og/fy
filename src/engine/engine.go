package engine

import "fmt"

func MakeTranslator(text *Text, model string, trans string) (Translator, error) {
	translators := make(map[string]Translator)

	g := &Google{}
	g.ApiUrl = "http://translate.google.com/translate_a/single"
	g.SupportModel = "sentence"
	g.Params = map[string]string{
		"client": "at",
		"dt":     "t",
	}
	translators[GOOGLE] = g

	c := &Cambridge{}
	c.ApiUrl = "https://dictionary.cambridge.org/zhs/%E8%AF%8D%E5%85%B8/%E8%8B%B1%E8%AF%AD-%E6%B1%89%E8%AF%AD-%E7%B9%81%E4%BD%93"
	translators[CAMBRIDGE] = c

	if text.LangType != EN {
		m := &Mojo{}
		m.ApiUrl = "https://api.mojidict.com/parse/functions/union-api"
		m.Params = map[string]string{
			"appId": "E62VyFVLMiW7kvbtVq3p",
		}
		translators[MOJO] = m
	}

	if trans != "" {
		if translator, ok := translators[trans]; !ok {
			return nil, fmt.Errorf("cannot find translator")
		} else {
			return translator, nil
		}
	}

	if model == "s" || model == "sentence" {
		return translators[GOOGLE], nil
	}

	if text.LangType == EN {
		return translators[CAMBRIDGE], nil
	} else {
		return translators[MOJO], nil
	}
}
