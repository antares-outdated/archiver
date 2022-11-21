package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 10010101

	// split binary by chunks (8): bits to bytes -> '10010101 10010101 10010101'

	// bytes to hex -> '20 30 3C'

	// return hexChunksStr
	return ""
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name isTed -> !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
