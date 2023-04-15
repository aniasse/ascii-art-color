package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func main() {
	ascii := make(map[byte][]string)
	var index byte = 32
	file, err := os.ReadFile("./standard.txt")
	if err != nil {
		log.Fatal("Error : Not a ascci file in the repertory")
	}
	Split := strings.Split(string(file), "\n")
	for i := 1; i+8 < len(Split); i += 9 {
		ascii[index] = Split[i : i+8]
		index++
	}
	tabascii := ascii
	if len(os.Args) == 2 {
		if len(os.Args[1]) != 0 {
			split := strings.Split(os.Args[1], "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			for _, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							Match(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							fmt.Println()
						} else {
							fmt.Println()
							break
						}
					}
				} else {
					fmt.Println("Error : Non-displayable character !!!")
				}

			}

		}
	} else if len(os.Args) == 3 && Flag(os.Args[1]) {
		if len(os.Args[2]) != 0 {
			split := strings.Split(os.Args[2], "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			for _, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) && Supported(ColorFlag(os.Args[1])) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							MatchColored(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							fmt.Println()
						} else {
							fmt.Println()
							break
						}
					}
				} else {
					fmt.Println("Error : Non-displayable character or Color not supported!!!")
				}
			}
		}
	} else if len(os.Args) == 4 && Flag(os.Args[1]) {
		if len(os.Args[3]) != 0 {
			split := strings.Split(os.Args[3], "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			for _, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) && Supported(ColorFlag(os.Args[1])) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							ToBeColored(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							fmt.Println()
						} else {
							fmt.Println()
							break
						}
					}
				} else {
					fmt.Println("Error : Non-displayable character or Color not supported!!!")
				}

			}

		}
	} else {
		fmt.Println("Error:\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something ")
	}
}
