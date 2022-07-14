package src

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func ReadBanner() (map[int][]string, error) {
	banners := os.Args[2]
	pathOfBanners := "banners/" + banners + ".txt"

	symbols := make(map[int][]string)
	var buf []string
	counter := 0
	key := 31

	hashesOfBanners := map[string]string{
		"shadow":     "a49d5fcb0d5c59b2e77674aa3ab8bbb1",
		"standard":   "a51f800619146db0c42d26db3114c99f",
		"thinkertoy": "8efd138877a4b281312f6dd1cbe84add",
	}

	if hashesOfBanners[banners] != Md5sum(pathOfBanners) {
		return map[int][]string{}, fmt.Errorf("you have changed a file")
	}
	file, err := os.Open(pathOfBanners)
	if err != nil {
		return map[int][]string{}, nil
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
	return symbols, nil
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

func Md5sum(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
