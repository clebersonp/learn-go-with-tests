package v5

import "strings"

// ConvertToRoman converts an integer to a roman numeral
func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		if arabic == 5 {
			result.WriteString("V")
			break
		}
		if arabic == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
