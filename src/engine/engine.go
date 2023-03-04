package engine

// const (
// 	Google = iota
// 	Baidu
// 	DeepL
// 	Hujiang
// 	Mojo
// 	Youdao
// )

type Translator interface {
	Translate(text, sl string) (string, string, error)
}

type Engine struct {
	Params       map[string]string
	SupportModel string
	ApiUrl       string
}

func GetAllTranslators() []Translator {
	translators := make([]Translator, 0)

	g := &GoogleTranslator{}
	g.ApiUrl = "http://translate.google.com/translate_a/single"
	g.SupportModel = "sentence"
	g.Params = map[string]string{
		"client": "at",
		"dt":     "t",
	}

	translators = append(translators, g)

	return translators
}
