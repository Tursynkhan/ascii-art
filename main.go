package main

import (
	"log"
	"main/src"
	"os"
)

func main() {
	// src.IsvalidArgs()
	args := []rune(os.Args[1])
	err := src.Isvalid(args)
	if err {
		log.Println(err)
	}
	symbols, err1 := src.ReadBanner()
	if err1 != nil {
		log.Println(err1)
	}
	src.ReadArgs(args, symbols)
}
