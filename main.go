package main

import (
	"fmt"
	"main/src"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not correct number of arguments")
		return
	}
	args := []rune(os.Args[1])
	err := src.Isvalid(args)
	if err {
		return
	}

	symbols := src.ReadBanner()
	src.ReadArgs(args, symbols)
}
