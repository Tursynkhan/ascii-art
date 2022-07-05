package src

import "fmt"

func PrintArt(output []string) []string {
	if output[0] != "" {
		for i, w := range output {
			fmt.Println(w)
			output[i] = ""
		}
		return output
	}
	fmt.Println()
	return output
}
