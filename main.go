package main

import (
	"bufio"
	"os"
	"fmt"
)
func main(){
	Read()
}

func Read()(map[int][]string,bool){
	symbols:= map[int][]string{}
	var buf []string
	counter:=0
	key:=31

	file,err:=os.Open("standard.txt")
	if err!=nil{
		return symbols,true
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
	fmt.Println(symbols)
	return symbols,false
}

// func PrintArt(buffer []string)[]string{
// 	if buffer[0]!=""{
// 		for i,let:=range buffer{
// 			fmt.Printf("%s\n", let)
// 			buffer[i]=""
// 		}
// 		return buffer
// 	}
// 	fmt.Println()
// 	return buffer
// }