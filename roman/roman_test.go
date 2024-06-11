package roman

import (
	"fmt"
	"testing"
)

type testCase struct {
	arabic int
	roman  string
}

func TestConvertToRoman(t *testing.T) {
	testCases := []testCase{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		// {5, "V"},
	}
	for _, currentCase := range testCases {
		desc := fmt.Sprintf("Test converting %d", currentCase.arabic)
		t.Run(desc, func(t *testing.T) {
			got := ConvertToRoman(currentCase.arabic)
			want := currentCase.roman

			if got != want {
				t.Errorf("expected %q, got %q", want, got)
			}
		})
	}
}
