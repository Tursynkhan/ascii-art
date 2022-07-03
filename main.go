package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
)
func main(){
	symbols:=Read()
	Print(symbols)
}

func Read()(map[int][]string){
	symbols:=make(map[int][]string)
	var buf []string
	counter:=0
	key:=31

	file,err:=os.Open("standard.txt")
	if err!=nil{
		log.Fatal(err)
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		if scanner.Text()==""{
			counter=0
			key++
			buf=nil
			continue
		}
		buf=append(buf,scanner.Text())
		// fmt.Printf("[%v]\n",buf)
		counter++
		if counter ==8{
			symbols[key]=buf
		}
	}
	return symbols
}
func Print(symbols map[int][]string){
	args:=os.Args[1:]
	runeWord := []rune(args[0])
	res:=""

	for i:=1;i<8;i++{
			for _,w:=range runeWord{
			res+=string(symbols[int(w)][i])
		}
		res+="\n"
	}	
	fmt.Println(res[:len(res)-1])
}
