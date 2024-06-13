package roman

import (
	"fmt"
	"testing"
)

type testCase struct {
	arabic int
	roman  string
}

var testCases = []testCase{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{15, "XV"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{99, "XCIX"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{999, "CMXCIX"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, currentCase := range testCases {
		desc := fmt.Sprintf("Test converting %d to roman", currentCase.arabic)
		t.Run(desc, func(t *testing.T) {
			got := ConvertToRoman(currentCase.arabic)
			want := currentCase.roman

			if got != want {
				t.Errorf("expected %q, got %q", want, got)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, currentCase := range testCases {
		desc := fmt.Sprintf("Test converting %q to arabic", currentCase.roman)
		t.Run(desc, func(t *testing.T) {
			got := ConvertToArabic(currentCase.roman)
			want := currentCase.arabic

			if got != want {
				t.Errorf("expected %d, got %d", want, got)
			}
		})
	}
}
