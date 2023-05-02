package ascii

import (
	"fmt"
	"os"
	"strings"
)

func AsciiColor(arg1 string, arg2 string) {
	ascii := make(map[byte][]string)
	var index byte = 32
	file, err := os.ReadFile("./file/standard.txt")
	if err != nil {
		fmt.Println("Error : Not a ascci file in the repertory")
		return
	}
	Split := strings.Split(string(file), "\n")
	for i := 1; i+8 < len(Split); i += 9 {
		ascii[index] = Split[i : i+8]
		index++
	}
	tabascii := ascii
	if len(arg2) != 0 {
		split := strings.Split(arg2, "\\n")
		if NewLine(split) {
			split = split[:len(split)-1]
		}
		for _, v := range split {
			tabrune := []rune(v)
			if Printable(tabrune) && Supported(ColorFlag(arg1)) {
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
}

func AsciiTobeColored(arg1 string, arg3 string) {
	ascii := make(map[byte][]string)
	var index byte = 32
	file, err := os.ReadFile("./file/standard.txt")
	if err != nil {
		fmt.Println("Error : Not a ascci file in the repertory")
		return
	}
	Split := strings.Split(string(file), "\n")
	for i := 1; i+8 < len(Split); i += 9 {
		ascii[index] = Split[i : i+8]
		index++
	}
	tabascii := ascii
	if len(arg3) != 0 {
		split := strings.Split(arg3, "\\n")
		if NewLine(split) {
			split = split[:len(split)-1]
		}
		for _, v := range split {
			tabrune := []rune(v)
			if Printable(tabrune) && Supported(ColorFlag(arg1)) {
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
}
