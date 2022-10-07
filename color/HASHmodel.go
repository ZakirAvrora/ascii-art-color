package color

import (
	"fmt"
	"log"
	"strconv"
)

func IsColorHash(color string) bool {
	if len(color) == 7 && rune(color[0]) == '#' {
		return true
	}
	return false
}

func HashColor(color string) string {
	var RGB []int
	runes := []rune(color[1:])
	num := 0

	for i := 0; i < len(runes); i++ {
		digit := 0
		if runes[i] == 'a' {
			digit = 10
		} else if runes[i] == 'b' {
			digit = 11
		} else if runes[i] == 'c' {
			digit = 12
		} else if runes[i] == 'd' {
			digit = 13
		} else if runes[i] == 'e' {
			digit = 14
		} else if runes[i] == 'f' {
			digit = 15
		} else {
			digittmp, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				log.Fatalln("Invalid representation of color hash code")
			}
			digit = digittmp
		}
		num = num*16 + digit
		if i == 1 || i == 3 || i == 5 {
			RGB = append(RGB, num)
			num = 0
		}
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", RGB[0], RGB[1], RGB[2])
}
