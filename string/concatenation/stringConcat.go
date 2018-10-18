package concatenation

import (
	"bytes"
	"strings"
)

func Concat(words []string) string {
	var result string
	for _, word := range words {
		result += word
	}
	return result
}

func Join(words []string) string {
	return strings.Join(words, "")
}

func Builder(words []string) string {
	var sBuilder strings.Builder
	for _, word := range words {
		sBuilder.WriteString(word)
	}
	return sBuilder.String()
}

func Buffer(words []string) string {
	var buffer bytes.Buffer
	for _, word := range words {
		buffer.WriteString(word)
	}
	return buffer.String()
}
