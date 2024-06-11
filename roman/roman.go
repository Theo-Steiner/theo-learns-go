package roman

import (
	"fmt"
	"strings"
)

var numerals = []struct {
	value      int
	symbol     string
	subtractor bool
}{
	{
		10, "X", true,
	},
	{
		5, "V", false,
	},
	{
		1, "I", true,
	},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	remainder := arabic
	for _, numeral := range numerals {
		for remainder >= numeral.value {
			result.WriteString(numeral.symbol)
			remainder -= numeral.value
			fmt.Printf("remainder: %d\n", remainder)
		}
	}
	return result.String()
}
