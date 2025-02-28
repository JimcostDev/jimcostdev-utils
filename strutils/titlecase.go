package strutils

import (
	"strings"
	"unicode"
)

// Convierte un string a formato título (primera letra de cada palabra en mayúscula)
func ToTitleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToTitle(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}
