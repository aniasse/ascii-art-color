package ascii

import (
	"fmt"
)

func color(s string, s2 string) {

	code := s
	if code == "black" || code == "rgb(0, 0, 0)" || code == "#000000" || code == "hsl(0, 0%, 0%)" {
		code = "\033[30m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "green" || code == "rgb(0, 255, 0)" || code == "#00ff00" || code == "hsl(120, 100%, 50%)" {
		code = "\033[32m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "red" || code == "rgb(255, 0, 0)" || code == "#ff0000" || code == "hsl(0, 100%, 50%)" {
		code = "\033[31m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "yellow" || code == "rgb(255, 255, 0)" || code == "#ffff00" || code == "hsl(60, 100%, 50%)" {
		code = "\033[33m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "blue" || code == "rgb(0, 0, 255)" || code == "#0000ff" || code == "hsl(240, 100%, 50%)" {
		code = "\033[34m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "cyan" || code == "rgb(0, 255, 255)" || code == "#00ffff" || code == "hsl(180, 100%, 50%)" {
		code = "\033[36m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "white" || code == "rgb(255, 255, 255)" || code == "#ffffff" || code == "hsl(0, 0%, 100%)" {
		code = "\033[37m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "silver" || code == "rgb(192, 192, 192)" || code == "#c0c0c0" || code == "hsl(0, 0%, 75%)" {
		code = "\033[90m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "grey" || code == "rgb(128, 128, 128)" || code == "#808080" || code == "hsl(0, 0%, 50%)" {
		code = "\033[37;90m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "chartreuse" || code == "rgb(127, 255, 0)" || code == "#7fff00" || code == "hsl(90, 100%, 50%)" {
		code = "\033[92m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "brown" || code == "rgb(165, 42, 42)" || code == "#a52a2a" || code == "hsl(0, 60%, 41%)" {
		code = "\033[33;22m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "purple" || code == "rgb(128, 0, 128)" || code == "#800080" || code == "hsl(300, 100%, 25%)" {
		code = "\033[35;22m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "gold" || code == "rgb(255, 215, 0)" || code == "#ffd700" || code == "hsl(51, 100%, 50%)" {
		code = "\033[33;1m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "pink" || code == "rgb(255, 192, 203)" || code == "#ffc0cb" || code == "hsl(350, 100%, 88%)" {
		code = "\033[95m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "salmon" || code == "rgb(250, 128, 114)" || code == "#fa8072" || code == "hsl(6, 93%, 71%)" {
		code = "\033[38;2;255;160;122m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "olive" || code == "rgb(128, 128, 0)" || code == "#808000" || code == "hsl(60, 100%, 25%)" {
		code = "\033[38;2;128;128;0m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "turquoise" || code == "rgb(64, 224, 208)" || code == "#40e0d0" || code == "hsl(174, 72%, 56%)" {
		code = "\033[38;2;64;224;208m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "dark-grey" || code == "rgb(64, 64, 64)" || code == "#404040" || code == "hsl(0, 0%, 25%)" {
		code = "\033[38;2;64;64;64m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "yellow-green" || code == "rgb(154, 205, 50)" || code == "#9acd32" || code == "hsl(80, 61%, 50%)" {
		code = "\033[38;2;154;205;50m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "navy-blue" || code == "rgb(0, 0, 128)" || code == "#000080" || code == "hsl(240, 100%, 25%)" {
		code = "\033[38;2;0;0;128m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "fucshia" || code == "rgb(255, 0, 255)" || code == "#ff00ff" || code == "hsl(300, 100%, 50%)" {
		code = "\033[38;2;255;0;255m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "olive-green" || code == "rgb(85, 107, 47)" || code == "#556b2f" || code == "hsl(82, 39%, 30%)" {
		code = "\033[38;2;85;107;47m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "saffron" || code == "rgb(244, 196, 48)" || code == "#f4c430" || code == "hsl(49, 89%, 59%)" {
		code = "\033[38;2;244;196;48m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "sky-blue" || code == "rgb(135, 206, 235)" || code == "#87ceeb" || code == "hsl(197, 71%, 73%)" {
		code = "\033[38;2;135;206;235m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "dark-red" || code == "rgb(139, 0, 0)" || code == "#8b0000" || code == "hsl(0, 100%, 27%)" {
		code = "\033[38;2;139;0;0m"
		fmt.Print(code + s2 + "\033[0m")
	} else if code == "orange" || code == "rgb(255, 165, 0)" || code == "#ffa500" || code == "hsl(39, 100%, 50%)" {
		code = "\033[38;5;208m"
		fmt.Print(code + s2 + "\033[0m")
	}

}

func Supported(c string) bool {
	if c == "black" || c == "white" || c == "red" || c == "green" || c == "blue" || c == "yellow" || c == "cyan" || c == "silver" || c == "grey" || c == "chartreuse" || c == "brown" || c == "purple" || c == "gold" || c == "pink" || c == "salmon" || c == "olive" || c == "turquoise" || c == "dark-grey" || c == "yellow-green" || c == "navy-blue" || c == "fucshia" || c == "olive-green" || c == "saffron" || c == "sky-blue" || c == "dark-red" || c == "orange" {
		return true
	} else if c == "rgb(0, 0, 0)" || c == "rgb(0, 255, 0)" || c == "rgb(255, 0, 0)" || c == "rgb(255, 255, 0)" || c == "rgb(0, 0, 255)" || c == "rgb(0, 255, 255)" || c == "rgb(255, 255, 255)" || c == "rgb(192, 192, 192)" || c == "rgb(128, 128, 128)" || c == "rgb(154, 205, 50)" || c == "rgb(165, 42, 42)" || c == "rgb(128, 0, 128)" || c == "rgb(255, 215, 0)" || c == "rgb(255, 192, 203)" || c == "rgb(250, 128, 114)" || c == "rgb(128, 128, 0)" || c == "rgb(64, 224, 208)" || c == "rgb(64, 64, 64)" || c == "rgb(127, 255, 0)" || c == "rgb(0, 0, 128)" || c == "rgb(255, 0, 255)" || c == "rgb(85, 107, 47)" || c == "rgb(244, 196, 48)" || c == "rgb(135, 206, 235)" || c == "rgb(139, 0, 0)" || c == "rgb(255, 165, 0)" {
		return true
	} else if c == "#000000" || c == "#00ff00" || c == "#ff0000" || c == "#ffff00" || c == "#0000ff" || c == "#00ffff" || c == "#ffffff" || c == "#c0c0c0" || c == "#808080" || c == "#7fff00" || c == "#a52a2a" || c == "#800080" || c == "#ffd700" || c == "#ffc0cb" || c == "#fa8072" || c == "#808000" || c == "#40e0d0" || c == "#404040" || c == "#9acd32" || c == "#000080" || c == "#ff00ff" || c == "#556b2f" || c == "#f4c430" || c == "#87ceeb" || c == "#8b0000" || c == "ffa500" {
		return true
	} else if c == "hsl(0, 0%, 0%)" || c == "hsl(120, 100%, 50%)" || c == "hsl(0, 100%, 50%)" || c == "hsl(60, 100%, 50%)" || c == "hsl(240, 100%, 50%)" || c == "hsl(180, 100%, 50%)" || c == "hsl(0, 0%, 100%)" || c == "hsl(0, 0%, 75%)" || c == "hsl(0, 0%, 50%)" || c == "hsl(90, 100%, 50%)" || c == "hsl(0, 60%, 41%)" || c == "hsl(300, 100%, 25%)" || c == "hsl(51, 100%, 50%)" || c == "hsl(350, 100%, 88%)" || c == "hsl(6, 93%, 71%)" || c == "hsl(60, 100%, 25%)" || c == "hsl(174, 72%, 56%)" || c == "hsl(0, 0%, 25%)" || c == "hsl(80, 61%, 50%)" || c == "hsl(240, 100%, 25%)" || c == "hsl(300, 100%, 50%)" || c == "hsl(82, 39%, 30%)" || c == "hsl(49, 89%, 59%)" || c == "hsl(197, 71%, 73%)" || c == "hsl(0, 100%, 27%)" || c == "hsl(39, 100%, 50%)" {
		return true
	} else {
		return false
	}
}
