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
		if IsDigit(char) {
			if IsFirstRune(str, pos) || IsDigit(prev) {
				return "", ErrInvalidString
			}
			RepeatRune(&sb, prev, int(char-'0'))
		} else {
			if !IsFirstRune(str, pos) && !IsDigit(prev) {
				sb.WriteRune(prev)
			}
			if IsLastRune(str, pos) {
				sb.WriteRune(char)
			}
		}
		prev = char
	}
	return sb.String(), nil
}

func RepeatRune(sb *strings.Builder, r rune, num int) {
	for i := 1; i <= num; i++ {
		sb.WriteRune(r)
	}
}

func IsFirstRune(str []rune, pos int) bool {
	return pos == 0 && len(str) != 0
}

func IsLastRune(str []rune, pos int) bool {
	return pos == len(str)-1 && len(str) != 0
}

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}
