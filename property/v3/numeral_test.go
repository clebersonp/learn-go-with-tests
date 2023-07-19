package property

import (
	"fmt"
	"testing"
)

// "Given input X, we expect Y", it should probably use table based tests.
var cases = []struct {
	Arabic int
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 11, Roman: "XI"},
	{Arabic: 12, Roman: "XII"},
	{Arabic: 13, Roman: "XIII"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 15, Roman: "XV"},
	{Arabic: 16, Roman: "XVI"},
	{Arabic: 17, Roman: "XVII"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 19, Roman: "XIX"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 21, Roman: "XXI"},
	{Arabic: 22, Roman: "XXII"},
	{Arabic: 23, Roman: "XXIII"},
	{Arabic: 24, Roman: "XXIV"},
	{Arabic: 25, Roman: "XXV"},
	{Arabic: 26, Roman: "XXVI"},
	{Arabic: 27, Roman: "XXVII"},
	{Arabic: 28, Roman: "XXVIII"},
	{Arabic: 29, Roman: "XXIX"},
	{Arabic: 30, Roman: "XXX"},
	{Arabic: 31, Roman: "XXXI"},
	{Arabic: 32, Roman: "XXXII"},
	{Arabic: 33, Roman: "XXXIII"},
	{Arabic: 34, Roman: "XXXIV"},
	{Arabic: 35, Roman: "XXXV"},
	{Arabic: 36, Roman: "XXXVI"},
	{Arabic: 37, Roman: "XXXVII"},
	{Arabic: 38, Roman: "XXXVIII"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 41, Roman: "XLI"},
	{Arabic: 42, Roman: "XLII"},
	{Arabic: 43, Roman: "XLIII"},
	{Arabic: 44, Roman: "XLIV"},
	{Arabic: 45, Roman: "XLV"},
	{Arabic: 46, Roman: "XLVI"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 48, Roman: "XLVIII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 51, Roman: "LI"},
	{Arabic: 52, Roman: "LII"},
	{Arabic: 53, Roman: "LIII"},
	{Arabic: 54, Roman: "LIV"},
	{Arabic: 55, Roman: "LV"},
	{Arabic: 56, Roman: "LVI"},
	{Arabic: 57, Roman: "LVII"},
	{Arabic: 58, Roman: "LVIII"},
	{Arabic: 59, Roman: "LIX"},
	{Arabic: 60, Roman: "LX"},
	{Arabic: 61, Roman: "LXI"},
	{Arabic: 62, Roman: "LXII"},
	{Arabic: 63, Roman: "LXIII"},
	{Arabic: 64, Roman: "LXIV"},
	{Arabic: 65, Roman: "LXV"},
	{Arabic: 66, Roman: "LXVI"},
	{Arabic: 67, Roman: "LXVII"},
	{Arabic: 68, Roman: "LXVIII"},
	{Arabic: 69, Roman: "LXIX"},
	{Arabic: 70, Roman: "LXX"},
	{Arabic: 71, Roman: "LXXI"},
	{Arabic: 72, Roman: "LXXII"},
	{Arabic: 73, Roman: "LXXIII"},
	{Arabic: 74, Roman: "LXXIV"},
	{Arabic: 75, Roman: "LXXV"},
	{Arabic: 76, Roman: "LXXVI"},
	{Arabic: 77, Roman: "LXXVII"},
	{Arabic: 78, Roman: "LXXVIII"},
	{Arabic: 79, Roman: "LXXIX"},
	{Arabic: 80, Roman: "LXXX"},
	{Arabic: 81, Roman: "LXXXI"},
	{Arabic: 82, Roman: "LXXXII"},
	{Arabic: 83, Roman: "LXXXIII"},
	{Arabic: 84, Roman: "LXXXIV"},
	{Arabic: 85, Roman: "LXXXV"},
	{Arabic: 86, Roman: "LXXXVI"},
	{Arabic: 87, Roman: "LXXXVII"},
	{Arabic: 88, Roman: "LXXXVIII"},
	{Arabic: 89, Roman: "LXXXIX"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 91, Roman: "XCI"},
	{Arabic: 92, Roman: "XCII"},
	{Arabic: 93, Roman: "XCIII"},
	{Arabic: 94, Roman: "XCIV"},
	{Arabic: 95, Roman: "XCV"},
	{Arabic: 96, Roman: "XCVI"},
	{Arabic: 97, Roman: "XCVII"},
	{Arabic: 98, Roman: "XCVIII"},
	{Arabic: 99, Roman: "XCIX"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 110, Roman: "CX"},
	{Arabic: 140, Roman: "CXL"},
	{Arabic: 199, Roman: "CXCIX"},
	{Arabic: 200, Roman: "CC"},
	{Arabic: 300, Roman: "CCC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 600, Roman: "DC"},
	{Arabic: 700, Roman: "DCC"},
	{Arabic: 800, Roman: "DCCC"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 2023, Roman: "MMXXIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q,", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, but want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, but want %d", got, test.Arabic)
			}
		})
	}
}
