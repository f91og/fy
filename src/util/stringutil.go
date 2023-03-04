package util

import (
	"unicode"
)

func CheckLanguageType(text string) string {
	return ""
}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func IsJapanese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is() {
			count++
			break
		}
	}
	return count > 0
}
