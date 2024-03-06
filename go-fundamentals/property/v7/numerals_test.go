package v7

import (
	"fmt"
	"testing"
)

func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Want   string
		Arabic int
	}{
		{Arabic: 1, Want: "I"},
		{Arabic: 2, Want: "II"},
		{Arabic: 3, Want: "III"},
		{Arabic: 4, Want: "IV"},
		{Arabic: 5, Want: "V"},
		{Arabic: 6, Want: "VI"},
		{Arabic: 7, Want: "VII"},
		{Arabic: 8, Want: "VIII"},
		{Arabic: 9, Want: "IX"},
		{Arabic: 10, Want: "X"},
		{Arabic: 11, Want: "XI"},
		{Arabic: 12, Want: "XII"},
		{Arabic: 13, Want: "XIII"},
		{Arabic: 14, Want: "XIV"},
		{Arabic: 15, Want: "XV"},
		{Arabic: 16, Want: "XVI"},
		{Arabic: 17, Want: "XVII"},
		{Arabic: 18, Want: "XVIII"},
		{Arabic: 19, Want: "XIX"},
		{Arabic: 20, Want: "XX"},
		{Arabic: 21, Want: "XXI"},
		{Arabic: 22, Want: "XXII"},
		{Arabic: 23, Want: "XXIII"},
		{Arabic: 24, Want: "XXIV"},
		{Arabic: 25, Want: "XXV"},
		{Arabic: 26, Want: "XXVI"},
		{Arabic: 27, Want: "XXVII"},
		{Arabic: 28, Want: "XXVIII"},
		{Arabic: 29, Want: "XXIX"},
		{Arabic: 30, Want: "XXX"},
		{Arabic: 31, Want: "XXXI"},
		{Arabic: 32, Want: "XXXII"},
		{Arabic: 33, Want: "XXXIII"},
		{Arabic: 34, Want: "XXXIV"},
		{Arabic: 35, Want: "XXXV"},
		{Arabic: 36, Want: "XXXVI"},
		{Arabic: 37, Want: "XXXVII"},
		{Arabic: 38, Want: "XXXVIII"},
		{Arabic: 39, Want: "XXXIX"},
		{Arabic: 40, Want: "XL"},
		{Arabic: 41, Want: "XLI"},
		{Arabic: 42, Want: "XLII"},
		{Arabic: 43, Want: "XLIII"},
		{Arabic: 44, Want: "XLIV"},
		{Arabic: 45, Want: "XLV"},
		{Arabic: 46, Want: "XLVI"},
		{Arabic: 47, Want: "XLVII"},
		{Arabic: 48, Want: "XLVIII"},
		{Arabic: 49, Want: "XLIX"},
		{Arabic: 50, Want: "L"},
		{Arabic: 100, Want: "C"},
		{Arabic: 400, Want: "CD"},
		{Arabic: 500, Want: "D"},
		{Arabic: 900, Want: "CM"},
		{Arabic: 1000, Want: "M"},
		{Arabic: 1984, Want: "MCMLXXXIV"},
		{Arabic: 3999, Want: "MMMCMXCIX"},
		{Arabic: 2014, Want: "MMXIV"},
		{Arabic: 1006, Want: "MVI"},
		{Arabic: 798, Want: "DCCXCVIII"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("ConvertToRoman(%d) = %q,", test.Arabic, test.Want), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("ConvertToRoman(%d) = %q, want %q", test.Arabic, got, test.Want)
			}
		})
	}
}
