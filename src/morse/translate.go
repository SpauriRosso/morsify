package morse

import (
	"strings"
)

func TextToMorse(statusc, text string) (status string, res string) {
	if strings.Contains(statusc, "fail") {
		status, res = statusc, ""
		println(statusc)
		return status, res
	}
	for _, v := range strings.ToUpper(text) {
		res += " " + Translate(v)
	}

	return "success", res
}

func Translate(char rune) string {
	return Codes[string(char)]
}
