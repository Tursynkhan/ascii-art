package main

import (
	"main/src"
	"os"
)

func main() {
	banners := "banners/" + os.Args[2] + ".txt"
	src.IsvalidArgs()
	args := []rune(os.Args[1])
	err := src.Isvalid(args)
	if err {
		return
	}

	symbols := src.ReadBanner(banners)
	src.ReadArgs(args, symbols)
}
