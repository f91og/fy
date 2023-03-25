package engine

// type TransName int

// const (
// 	GOOGLE TransName = iota
// 	MOJO
// 	BING
// 	CAMBRIDGE
// )

// 翻译输入的文本
type Text struct {
	Value    string
	LangType string
}

type Translator interface {
	Translate(text *Text) (string, string, error)
}

type Engine struct {
	Params       map[string]string
	SupportModel string
	ApiUrl       string
}
