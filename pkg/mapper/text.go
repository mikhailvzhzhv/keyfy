package mapper

import "strings"

func Mapkeys(text string) string {
	return strings.Map(mapkey, text)
}

func mapkey(s rune) rune {
	if val, ok := mapping[s]; ok {
		return val
	}

	return s
}
