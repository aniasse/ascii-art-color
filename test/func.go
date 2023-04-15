package main

func Match(r rune, i int, ascii map[byte][]string) string {
	var result string
	for ind, v := range ascii {
		if rune(ind) == r {
			result = v[i]
			break
		}
	}
	return result
}

func NewLine(tab []string) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			return false
		}
	}
	return true
}

func Printable(tab []rune) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] < 32 || tab[i] > 126 {
			return false
		}
	}
	return true
}

func Banner(s string) string {
	return "./" + s + ".txt"
}

func Flag(s string) bool {
	tab := []rune(s)
	if len(tab) > 8 && string(tab[:8]) == "--color=" {
		return true
	} else {
		return false
	}
}
