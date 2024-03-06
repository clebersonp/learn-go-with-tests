package v7

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

// 1984 : MCMLXXXIV

var allRomanNumerals = []RomanNumeral{
	{Value: 1000, Symbol: "M"},
	{Value: 900, Symbol: "CM"},
	{Value: 500, Symbol: "D"},
	{Value: 400, Symbol: "CD"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 50, Symbol: "L"},
	{Value: 40, Symbol: "XL"},
	{Value: 10, Symbol: "X"},
	{Value: 9, Symbol: "IX"},
	{Value: 5, Symbol: "V"},
	{Value: 4, Symbol: "IV"},
	{Value: 1, Symbol: "I"},
}

// ConvertToRoman converts an integer to a roman numeral
func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// ConvertToArabic converts a roman numeral to an integer
func ConvertToArabic(roman string) int {

	total := 0
	for range roman {
		total++
	}

	return total
}
