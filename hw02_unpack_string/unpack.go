package hw02unpackstring

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

	if len(s) == 0 {
		return "", nil
	}

	for _, r := range s {
		switch {
		case unicode.IsDigit(r) && !escapeMode:
			if prevRune == 0 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(r))
			if count > 0 {
				builder.WriteString(strings.Repeat(string(prevRune), count))
			}
			prevRune = 0
		case r == '\\' && !escapeMode:
			escapeMode = true
		case escapeMode && r == 'n':
			escapeMode = false
			continue
		default:
			if prevRune != 0 {
				builder.WriteRune(prevRune)
			}
			prevRune = r
			escapeMode = false
		}
	}

	if escapeMode {
		return "", ErrInvalidString
	}

	if prevRune != 0 && !escapeMode {
		builder.WriteRune(prevRune)
	}
	return builder.String(), nil
}
