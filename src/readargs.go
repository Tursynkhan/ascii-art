package src

import "fmt"

func ReadArgs(args []rune, symbols map[int][]string) {
	output := make([]string, 8)
	buf := false
	for k := 0; k < len(args); k++ {
		if args[k] == '\n' {
			output = PrintArt(output)
			continue
		}
		if len(args) != k+1 && (args[k] == '\\' && args[k+1] == 'n') {
			output = PrintArt(output)
			k++
			continue
		}
		buf = true
		for i, w := range symbols[int(args[k])] {
			output[i] += w
		}
		if k+1 >= len(args) {
			buf = false
			PrintArt(output)
			break
		}
	}
	if buf {
		fmt.Println()
		return
	}
}
