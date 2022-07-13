package src

import (
	"bufio"
	"fmt"
	"os"
)

func ReadBanner(banners string) map[int][]string {
	symbols := make(map[int][]string)
	var buf []string
	counter := 0
	key := 31

	file, err := os.Open(banners)
	if err != nil {
		return symbols
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			counter = 0
			key++
			buf = nil
			continue
		}
		buf = append(buf, scanner.Text())
		// fmt.Printf("[%v]\n",buf)
		counter++
		if counter == 8 {
			symbols[key] = buf
		}
	}
	return symbols
}

func Isvalid(args []rune) bool {
	for _, r := range args {
		if (r < 32 || r > 127) && r != 10 {
			fmt.Println("Please, input only ascii character!")
			return true
		}
	}
	return false
}
