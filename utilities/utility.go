package utilities

import (
	"bufio"
	"crypto/sha256"
	"log"
	"os"
)

var TEMPLATE_STANDARD_HASH_BYTE = []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}

// Function to close the text file
func FileClose(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalf("Problem in closing acsii-art template file")
	}
}

// Function of making maps between rune and printing string
func Initialize() map[rune]string {
	//-----------------------------------------
	readFile, err := os.Open("./templates/standard.txt")
	defer FileClose(readFile)
	if err != nil {
		log.Fatalf("failed to open standard.txt")
	}
	//-----------------------------------------

	if !CheckFileValidity() {
		log.Fatalf("Empty file or Corrupted file to open: file is %s\n", "./templates/standard.txt")
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Scan()

	m := make(map[rune]string) // newMap that stores runes' ascii-art
	var r rune = ' '
	i := 1
	for scanner.Scan() { // Every eight lines are stored in map as ascii-art
		if i == 9 { // with corresponding key rune
			i = 1
			r++ // key rune is updated every eight lines
			continue
		}
		if i != 8 {
			m[r] += scanner.Text() + "\n"
		} else {
			m[r] += scanner.Text()
		}

		i++
	}

	return m
}

// Checks the inporper runes that do not have ascii-art
func CheckRunes(str string) bool {
	for _, r := range str {
		if r >= ' ' && r <= '~' {
			continue
		} else if r == '\n' {
			continue
		} else {
			return false
		}
	}
	return true
}

func CheckFileValidity() bool {
	banner, _ := os.ReadFile("./templates/standard.txt")
	h := sha256.New()
	h.Write(banner)

	return string(h.Sum(nil)) == string(TEMPLATE_STANDARD_HASH_BYTE)
}

func ErrorMsg() string {
	return "Usage: go run . [STRING] [OPTION] \nEX: go run . something --color=<color>"
}
