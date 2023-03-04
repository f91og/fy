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

	g := &Google{}
	g.ApiUrl = "http://translate.google.com/translate_a/single"
	g.SupportModel = "sentence"
	g.Params = map[string]string{
		"client": "at",
		"dt":     "t",
	}

	m := &Mojo{}
	m.ApiUrl = "https://api.mojidict.com/parse/functions/union-ap"

	translators = append(translators, g)

	return translators
}
