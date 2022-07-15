package src

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReadArgs(args []rune, symbols map[int][]string) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	size := strings.Split(string(out), " ")
	width, err := strconv.Atoi(size[1][:len(size[1])-1])
	if err != nil {
		log.Fatal(err)
	}
	counter := 0

	output := make([]string, 8)
	newline := false
	for k := 0; k < len(args); k++ {
		if args[k] == '\n' {
			output = PrintArt(output)
			counter = 0
			continue
		}
		if len(args) != k+1 && (args[k] == '\\' && args[k+1] == 'n') {
			output = PrintArt(output)
			counter = 0
			k++
			continue
		}
		newline = true
		if counter+len(symbols[int(args[k])][0]) >= width {
			output = PrintArt(output)
			counter = 0
		}
		for i, w := range symbols[int(args[k])] {
			output[i] += w
		}
		counter += len(symbols[int(args[k])][0])
		if k+1 >= len(args) {
			newline = false
			PrintArt(output)
			break
		}
	}
	if newline {
		fmt.Println()
		return
	}
}

// func IsvalidArgs() error {
// 	if len(os.Args) != 3 {
// 		errors.New("The arguments should be : go run . [STRING] [BANNER]")
// 		// fmt.Println("Not correct number of arguments")
// 		// fmt.Println("The arguments should be : go run . [STRING] [BANNER]")
// 		return nil
// 	}

// 	if os.Args[2] != "standard" || os.Args[2] != "shadow" || os.Args[2] != "thinkertoy" {
// 		// fmt.Print("[BANNER] should be : standard, shadow, thinkertoy")
// 		errors.New("[BANNER] should be : standard, shadow, thinkertoy")
// 		return nil
// 	}
// 	return nil
// }
