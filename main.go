package main

import (
	"fmt"
	"main/src"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not correct number of arguments")
		fmt.Println("The arguments should be : go run . [STRING] [BANNER]")
		return
	}
	banners := "banners/" + os.Args[2] + ".txt"

	if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
		fmt.Println("[BANNER] should be : standard, shadow, thinkertoy")
		return
	}
	args := []rune(os.Args[1])
	err := src.Isvalid(args)
	if err {
		return
	}

	symbols := src.ReadBanner(banners)
	src.ReadArgs(args, symbols)
}
