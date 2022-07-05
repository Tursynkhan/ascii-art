package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
)
func main(){
	// if len(os.Args)!=2{
	// 	fmt.Println("Not correct number of arguments")
	// 	return
	// }
	// args := []rune(os.Args[1])
	// err:=Isvalid(args)
	// if err{
	// 	return
	// }
	args:="hello"
	symbols:=ReadBanner()
	ReadArgs(args,symbols)
	
}

func ReadBanner()(map[int][]string){
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
func ReadArgs(args []rune, symbols map[int][]string){
	output:= make([]string,8)
	for  k:=0;k<len(args);k++{
		if args[k]=='\n'{
			output=PrintArt(output)
			continue
		}
		if  len(args)!=k+1 && (args[k]=='\\' && args[k+1]=='n'){
			output=PrintArt(output)
			continue
		}
		for i,w:=range symbols[int(args[k])]{
			output[i]+=w
		}
		// if len(args)<k+1{
		// 	output=PrintArt(output)
		// 	break
		// }
	}

}

func PrintArt(output []string)[]string{	
	if output[0]!=""{
		for i,w:=range output{
			fmt.Println(w)
			output[i]=""
		}
		return output
	}
	fmt.Println()
	return output
}

func Isvalid(args []rune)bool{
	for _,r:=range args{
		if (r<32 || r>127)&&r!=10{
			fmt.Println("Please, input only ascii character!")
			return true
		}
	}
	return false
}