package color

import (
	"fmt"
	"strings"
)

func PrintWihoutColor(strArr []string, m map[rune]string) {
	for j, s := range strArr {

		if j == 0 && s == "" && strArr[1] == "" {
			continue
		} else if s == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			newStr := ""
			for _, r := range s {
				strArr := strings.Split(m[r], "\n")
				newStr += strArr[i]
			}
			fmt.Println(newStr)
		}
	}
}

func PrintWithColorIndexRange(strArr []string, m map[rune]string, colorcode string, indexRange []int) {
	globalIndex := 0
	for j, s := range strArr {

		if j == 0 && s == "" && strArr[1] == "" {
			continue
		} else if s == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			newStr := ""
			for j, r := range s {
				strArr := strings.Split(m[r], "\n")

				if len(indexRange) == 1 && globalIndex == indexRange[0] {
					newStr += colorcode + strArr[i] + "\033[0m"
				} else if len(indexRange) == 2 && globalIndex >= indexRange[0] && globalIndex < indexRange[1] {
					newStr += colorcode + strArr[i] + "\033[0m"
				} else {
					newStr += strArr[i]
				}

				globalIndex++
				if j == len(s)-1 {
					globalIndex = globalIndex - len(s)
				}

			}
			fmt.Println(newStr)
			if i == 7 {
				globalIndex += len(s)
			}
		}
		globalIndex++
	}
}

func PrintWithColorLetterList(strArr []string, m map[rune]string, colorcode string, indexes []int) {
	globalIndex := 0
	for j, s := range strArr {

		if j == 0 && s == "" && strArr[1] == "" {
			continue
		} else if s == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			newStr := ""
			for j, r := range s {
				strArr := strings.Split(m[r], "\n")

				if contains(indexes, globalIndex) {
					newStr += colorcode + strArr[i] + "\033[0m"
				} else {
					newStr += strArr[i]
				}

				globalIndex++
				if j == len(s)-1 {
					globalIndex = globalIndex - len(s)
				}

			}
			fmt.Println(newStr)
			if i == 7 {
				globalIndex += len(s)
			}
		}
	}
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func PrintWithColorAll(strArr []string, m map[rune]string, colorcode string) {
	for j, s := range strArr {

		if j == 0 && s == "" && strArr[1] == "" {
			continue
		} else if s == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			newStr := ""
			for _, r := range s {
				strArr := strings.Split(m[r], "\n")
				newStr += colorcode + strArr[i]
			}
			fmt.Println(newStr)

		}

	}
}
