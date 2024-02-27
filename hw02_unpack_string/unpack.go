package hw02_unpack_string

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	var prevRune rune
	escapeMode := false

	for i, r := range s {
		if unicode.IsDigit(r) && !escapeMode {
			if i == 0 || prevRune == 0 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(r))
			if count > 0 {
				builder.WriteString(strings.Repeat(string(prevRune), count))
			}
			prevRune = 0
		} else {
			if r == '\\' && !escapeMode {
				escapeMode = true
				continue
			} else {
				if prevRune != 0 {
					builder.WriteRune(prevRune)
				}
				prevRune = r
				escapeMode = false
			}
		}
	}
	if prevRune != 0 && !escapeMode {
		builder.WriteRune(prevRune)
	}
	return builder.String(), nil
}
