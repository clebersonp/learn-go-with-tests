package v3

import "strings"

// ConvertToRoman converts an integer to a roman numeral
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
