package mapper

func Mapkeys(text string) string {
	text_len := len(text)
	mapped_string := make([]rune, text_len)

	for i, char := range text {
		mapped_string[i] = mapkey(char)
	}

	return string(mapped_string)
}

func mapkey(s rune) rune {
	if val, ok := mapping[s]; ok {
		return val
	}

	return s
}
