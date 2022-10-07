package color

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func IsColorRGB(color string) bool {
	if len(color) < 5 {
		return false
	}
	if color[:4] == "rgb(" && color[len(color)-1] == ')' {
		return true
	}

	return false
}

func RGBColor(color string) string {
	var r, g, b string
	s := strings.Split(color[4:len(color)-1], ",")

	if !CheckRGBCode(s) {
		log.Fatalln("Invalid representation of RGB color\nExample:--color=rgb(255,120,130)")
	} else {
		r, g, b = s[0], s[1], s[2]
	}

	return fmt.Sprintf("\033[38;2;%s;%s;%sm", r, g, b)
}

func CheckRGBCode(code []string) bool {
	if len(code) != 3 {
		return false
	}

	for _, e := range code {
		num, err := strconv.Atoi(e)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}
