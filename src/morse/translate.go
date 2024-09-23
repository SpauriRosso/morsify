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
		res += " " + Translate(v, Codes)
	}
	return "success", res
}

func MorseToText(morse string) (status, res string) {
	tabMorse := strings.Split(morse, " ")
	for _, v := range tabMorse {
		res += TranslateMorse(v, MorseTxt)
	}
	return "success", res
}

func TranslateMorse(morse string, m map[string]string) string {
	return m[morse]
}

func Translate(char rune, m map[string]string) string {
	return m[string(char)]
}
