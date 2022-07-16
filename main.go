package main

import (
	"log"
	"main/src"
	"os"
)

func main() {
	input := os.Args
	switch len(input) {
	case 2:
		banners := "standard"
		pathOfBanners := "banners/" + banners + ".txt"

		symbols, err := src.ReadBanner(banners, pathOfBanners)
		if err != nil {
			log.Println(err)
			return
		}
		args := []rune(os.Args[1])

		if src.Isvalid(args) {
			log.Println(err)
			return
		}

		src.ReadArgs(args, symbols)
	case 3:
		banners := os.Args[2]
		pathOfBanners := "banners/" + banners + ".txt"

		symbols, err := src.ReadBanner(banners, pathOfBanners)
		if err != nil {
			log.Println(err)
			return
		}
		args := []rune(os.Args[1])

		if src.Isvalid(args) {
			log.Println(err)
			return
		}

		src.ReadArgs(args, symbols)
	default:
		if err := src.IsvalidArgs(input); err != nil {
			log.Fatal(err)
		}
	}
}
