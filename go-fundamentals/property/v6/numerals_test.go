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
