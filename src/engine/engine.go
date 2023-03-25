package engine

func GetTranslators(text *Text) map[string]Translator {
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
	translators["cambridge"] = c

	if text.LangType != EN {
		m := &Mojo{}
		m.ApiUrl = "https://api.mojidict.com/parse/functions/union-api"
		m.Params = map[string]string{
			"appId": "E62VyFVLMiW7kvbtVq3p",
		}
		translators[MOJO] = m
	}

	return translators
}
