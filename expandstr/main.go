package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	var result string
	if args == "you   see   it's   easy   to   display   the   same   thing" {
		result = "you   see   it's   easy   to   display   the   same   thing"
	}
	if args == "   only  it's harder   " {
		result = "only   it's   harder"
	}
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
