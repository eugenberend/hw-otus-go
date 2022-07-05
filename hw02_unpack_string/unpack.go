package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	str := []rune(input)

	var sb strings.Builder
	var prev rune
	for pos, char := range str {
		if unicode.IsDigit(char) {
			if pos == 0 || unicode.IsDigit(prev) {
				return "", ErrInvalidString
			}
			for i := 1; i <= int(char-'0'); i++ {
				sb.WriteRune(prev)
			}
		} else if unicode.IsLetter(char) {
			if unicode.IsLetter(prev) {
				sb.WriteRune(prev)
			}
			if pos == len(str)-1 {
				sb.WriteRune(char)
			}
		}
		prev = char
	}
	return sb.String(), nil
}
