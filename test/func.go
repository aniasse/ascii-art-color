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

func ColorFlag(S string) string {
	var color string
	var j int
	for i := 0; i < len(S); i++ {
		if S[i] != '=' {
			j++
		} else {
			break
		}
	}
	for k := j + 1; k < len(S); k++ {
		color += string(S[k])
	}
	return color
}

func ToColor(s string, r rune) bool {
	tab := []rune(s)
	for i := 0; i < len(tab); i++ {
		if tab[i] == r {
			return true
		}
	}
	return false
}
