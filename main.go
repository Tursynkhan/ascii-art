package main

import (
	"fmt"
	"main/src"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not correct number of arguments")
		return
	}
	banners := "banners/" + os.Args[2] + ".txt"

	args := []rune(os.Args[1])
	err := src.Isvalid(args)
	if err {
		return
	}

	symbols := src.ReadBanner(banners)
	src.ReadArgs(args, symbols)
}
