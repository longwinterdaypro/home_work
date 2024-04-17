package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	sb := &strings.Builder{}
	var sub string
	var screen bool

	for _, char := range s {
		switch {
		case (48 <= char && char <= 57 && !screen):
			if sub == "" {
				return "", ErrInvalidString
			}
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			sb.WriteString(strings.Repeat(sub, num))
			sub = ""
		case char == 92 && !screen:
			screen = true
		case screen && (char != 92 && (48 > char || char > 57)):
			return "", ErrInvalidString
		default:
			if sub != "" {
				sb.WriteString(sub)
				screen = false
			}
			sub = string(char)
		}
	}
	sb.WriteString(sub)
	if screen {
		return "", ErrInvalidString
	}
	return sb.String(), nil
}
