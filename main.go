package main

import (
	"log"
	"main/src"
	"os"
)

func main() {
	banners := os.Args[2]
	pathOfBanners := "banners/" + banners + ".txt"
	err := src.IsvalidArgs()
	if err != nil {
		log.Println(err)
		return
	}
	err3 := src.HashValid(banners, pathOfBanners)
	if err3 != nil {
		log.Println(err3)
		return
	}
	args := []rune(os.Args[1])
	err1 := src.Isvalid(args)
	if err1 {
		log.Println(err)
		return
	}
	symbols, err2 := src.ReadBanner(banners, pathOfBanners)
	if err2 != nil {
		log.Println(err1)
		return
	}
	src.ReadArgs(args, symbols)
}
