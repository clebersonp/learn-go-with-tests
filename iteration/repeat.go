package iteration

import "strings"

func Repeat(character string, count int) string {
	if count < 0 {
		return character
	}
	return strings.Repeat(character, count)
}