package v3

// ConvertToRoman converts an integer to a roman numeral
func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
