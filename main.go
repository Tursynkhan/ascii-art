package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
)
func main(){
	if len(os.Args)!=2{
		fmt.Println("Not correct number of arguments")
		return
	}
	args := []rune(os.Args[1])
	err:=Isvalid(args)
	if err{
		return
	}
	symbols:=Read()
	PrintArt(args,symbols)
	Isvalid(args)
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
func PrintArt(args []rune, symbols map[int][]string){
	output:= make([]string,8)
	for  k:=0;k<len(args);k++{
		for i,w:=range symbols[int(args[k])]{
			output[i]+=w
		}
	}
	for _,r:=range output{
		fmt.Println(r)
	}
}

func Isvalid(args []rune)bool{
	for _,r:=range args{
		if (r<32 || r>127){
			fmt.Println("Please, input onli ascii character!")
			return true
		}
	}
	return false
}