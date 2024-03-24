package translation

import "strings"

func Translate(word string, language string) string {
	word = sanitizeInput(word)

	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
