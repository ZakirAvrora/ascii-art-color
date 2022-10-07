package utilities

import (
	"log"
	"strconv"
	"strings"
)

func GetIndexesForColoring(cmd string, str string) []int {
	var resultIndex []int
	start := 0
	end := len(str)
	indexRange := strings.Split(cmd, ":")

	if len(indexRange) == 1 {
		index := StrToInt(indexRange[0])
		resultIndex = append(resultIndex, index)
		return resultIndex
	} else if len(indexRange) == 2 && indexRange[0] == "" && indexRange[1] == "" {
	} else if len(indexRange) == 2 && indexRange[0] != "" && indexRange[1] == "" {
		start = StrToInt(indexRange[0])
	} else if len(indexRange) == 2 && indexRange[0] == "" && indexRange[1] != "" {
		end = StrToInt(indexRange[1])
	} else if len(indexRange) == 2 && indexRange[0] != "" && indexRange[1] != "" {
		start = StrToInt(indexRange[0])
		end = StrToInt(indexRange[1])
		if start >= end {
			log.Fatalln("Invalid representation of index range for coloring: start index > end index")
		}
	} else {
		log.Fatalln("Invalid representation of index range for coloring")
	}

	resultIndex = append(resultIndex, start, end)
	return resultIndex
}

func StrToInt(numStr string) int {
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatalln("Invalid representation of index range for coloring: the invalid index")
	}
	return numInt
}

func GetIndexMapsForColoring(cmd string, newstr string) []int {
	strArr := strings.Split(strings.ReplaceAll(cmd, " ", ""), ",")
	str := strings.ReplaceAll(newstr, "\n", "")
	var indexes []int

	for _, subString := range strArr {
		if len(subString) == 1 {
			for j, r := range str {
				if string(r) == subString {
					indexes = append(indexes, j)
				}
			}
		} else {
			if strings.Contains(str, subString) {
				for i := 0; i <= len(str)-len(subString); i++ {
					if string(str[i:i+len(subString)]) == subString || (i == len(str)-len(subString)-1 && string(str[i:i+len(subString)]) == subString) {
						for j := 0; j < len(subString); j++ {
							indexes = append(indexes, i+j)
						}
						i += len(subString) - 1
					}
				}
			}
		}
	}
	return removeDuplicateInt(indexes)
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
