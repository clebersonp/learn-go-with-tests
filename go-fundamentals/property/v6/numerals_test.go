package v6

import "testing"

func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Description, Want string
		Arabic            int
	}{
		{Description: "1 gets converted to I", Arabic: 1, Want: "I"},
		{Description: "2 gets converted to II", Arabic: 2, Want: "II"},
		{Description: "3 gets converted to III", Arabic: 3, Want: "III"},
		{Description: "4 gets converted to IV", Arabic: 4, Want: "IV"},
		{Description: "5 gets converted to V", Arabic: 5, Want: "V"},
		{Description: "6 gets converted to VI", Arabic: 6, Want: "VI"},
		{Description: "7 gets converted to VII", Arabic: 7, Want: "VII"},
		{Description: "8 gets converted to VIII", Arabic: 8, Want: "VIII"},
		{Description: "9 gets converted to IX", Arabic: 9, Want: "IX"},
		{Description: "10 gets converted to X", Arabic: 10, Want: "X"},
		{Description: "11 gets converted to XI", Arabic: 11, Want: "XI"},
		{Description: "12 gets converted to XII", Arabic: 12, Want: "XII"},
		{Description: "13 gets converted to XIII", Arabic: 13, Want: "XIII"},
		{Description: "14 gets converted to XIV", Arabic: 14, Want: "XIV"},
		{Description: "15 gets converted to XV", Arabic: 15, Want: "XV"},
		{Description: "16 gets converted to XVI", Arabic: 16, Want: "XVI"},
		{Description: "17 gets converted to XVII", Arabic: 17, Want: "XVII"},
		{Description: "18 gets converted to XVIII", Arabic: 18, Want: "XVIII"},
		{Description: "19 gets converted to XIX", Arabic: 19, Want: "XIX"},
		{Description: "20 gets converted to XX", Arabic: 20, Want: "XX"},
		{Description: "21 gets converted to XXI", Arabic: 21, Want: "XXI"},
		{Description: "22 gets converted to XXII", Arabic: 22, Want: "XXII"},
		{Description: "23 gets converted to XXIII", Arabic: 23, Want: "XXIII"},
		{Description: "24 gets converted to XXIV", Arabic: 24, Want: "XXIV"},
		{Description: "25 gets converted to XXV", Arabic: 25, Want: "XXV"},
		{Description: "26 gets converted to XXVI", Arabic: 26, Want: "XXVI"},
		{Description: "27 gets converted to XXVII", Arabic: 27, Want: "XXVII"},
		{Description: "28 gets converted to XXVIII", Arabic: 28, Want: "XXVIII"},
		{Description: "29 gets converted to XXIX", Arabic: 29, Want: "XXIX"},
		{Description: "30 gets converted to XXX", Arabic: 30, Want: "XXX"},
		{Description: "31 gets converted to XXXI", Arabic: 31, Want: "XXXI"},
		{Description: "32 gets converted to XXXII", Arabic: 32, Want: "XXXII"},
		{Description: "33 gets converted to XXXIII", Arabic: 33, Want: "XXXIII"},
		{Description: "34 gets converted to XXXIV", Arabic: 34, Want: "XXXIV"},
		{Description: "35 gets converted to XXXV", Arabic: 35, Want: "XXXV"},
		{Description: "36 gets converted to XXXVI", Arabic: 36, Want: "XXXVI"},
		{Description: "37 gets converted to XXXVII", Arabic: 37, Want: "XXXVII"},
		{Description: "38 gets converted to XXXVIII", Arabic: 38, Want: "XXXVIII"},
		{Description: "39 gets converted to XXXIX", Arabic: 39, Want: "XXXIX"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("ConvertToRoman(%d) = %q, want %q", test.Arabic, got, test.Want)
			}
		})
	}
}
