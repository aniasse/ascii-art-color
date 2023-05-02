package main

import (
	Func "ascii/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		Func.AsciiArt(os.Args[1])
	} else if len(os.Args) == 3 && !Func.Flag(os.Args[1]) {
		Func.AsciiFs(os.Args[1], os.Args[2])
	} else if len(os.Args) == 3 && Func.Flag(os.Args[1]) {
		Func.AsciiColor(os.Args[1], os.Args[2])
	} else if len(os.Args) == 4 && Func.Flag(os.Args[1]) {
		Func.AsciiTobeColored(os.Args[1], os.Args[3])
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something ")
	}
}
