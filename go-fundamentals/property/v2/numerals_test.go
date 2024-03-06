package v2

import "testing"

func TestConvertToRoman(t *testing.T) {
	tests := []struct {
		name, want string
		arabic     int
	}{
		{name: "1 gets converted to I", arabic: 1, want: "I"},
		{name: "2 gets converted to II", arabic: 2, want: "II"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToRoman(tt.arabic)
			if got != tt.want {
				t.Errorf("ConvertToRoman(%d) = %q, want %q", tt.arabic, got, tt.want)
			}
		})
	}
}
