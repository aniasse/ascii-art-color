package main

import (
	"fmt"
	"os"
)

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

func Match(r rune, i int, ascii map[byte][]string) {
	for ind, v := range ascii {
		if rune(ind) == r {
			fmt.Print(v[i])
		}
	}
}
func MatchColored(r rune, i int, ascii map[byte][]string) {
	couleur := ColorFlag(os.Args[1])
	for ind, v := range ascii {
		if rune(ind) == r {
			color(couleur, v[i])
		}
	}
}

func ToBeColored(r rune, i int, ascii map[byte][]string) {
	couleur := ColorFlag(os.Args[1])
	for ind, v := range ascii {
		if rune(ind) == r && ToColor(os.Args[2], r) {
			color(couleur, v[i])
		}
		if rune(ind) == r && !ToColor(os.Args[2], r) {
			fmt.Print(v[i])
		}
	}
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
