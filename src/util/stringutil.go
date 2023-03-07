package util

import (
	"strings"
	"unicode"
)

// 如果出现了一个假名则判定为日文
// 否则如果全是英文字母或者是space则认为是英文
// 否则是中文
func CheckLangType(text string) string {
	lang := "zh"

	for _, v := range text {
		if unicode.Is(unicode.Hiragana, v) || unicode.Is(unicode.Katakana, v) {
			lang = "ja"
		}
		// } else if unicode.IsLetter(v) {
		// 	lang = "en"
		// }
	}

	if enCharOnly(text) {
		lang = "en"
	}

	return lang
}

const alpha = "abcdefghijklmnopqrstuvwxyz, "

func enCharOnly(s string) bool {
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
