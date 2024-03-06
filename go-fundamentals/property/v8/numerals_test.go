package v8

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Roman  string
	Arabic int
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
	{Arabic: 100, Roman: "C"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("ConvertToRoman(%d) = %q,", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("ConvertToRoman(%d) = %q, want %q", test.Arabic, got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("ConvertToArabic(%q) = %d,", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("ConvertToArabic(%q) = %d, want %d", test.Roman, got, test.Arabic)
			}
		})
	}
}
