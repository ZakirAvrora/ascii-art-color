package color

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func IsColorHSL(color string) bool {
	if len(color) < 5 {
		return false
	}
	if color[:4] == "hsl(" && color[len(color)-1] == ')' {
		return true
	}

	return false
}

func HSLColor(color string) string {
	var r, g, b string
	str := strings.Split(color[4:len(color)-1], ",")
	if !CheckHSLCode(str) {
		log.Fatalln("Invalid representation of HSL color\nExample:--color=hsl(360,100%,100%)")
	} else {
		h, _ := strconv.ParseFloat(str[0], 64)
		s, _ := strconv.ParseFloat(str[1][:len(str[1])-1], 64)
		l, _ := strconv.ParseFloat(str[2][:len(str[2])-1], 64)
		result := Hsl2Rgb(h, s, l)
		r = strconv.Itoa(result[0])
		g = strconv.Itoa(result[1])
		b = strconv.Itoa(result[2])
	}

	return fmt.Sprintf("\033[38;2;%s;%s;%sm", r, g, b)
}

func CheckHSLCode(code []string) bool {
	if len(code) != 3 {
		return false
	}
	for i, e := range code {
		if i == 0 {
			num, err := strconv.ParseFloat(e, 64)
			if err != nil || (num < 0 || num > 360) {
				return false
			}
		} else {
			if e[len(e)-1] != '%' {
				return false
			}
			num, err := strconv.ParseFloat(e[:len(e)-1], 64)
			if err != nil || (num < 0 || num > 100) {
				return false
			}
		}
	}
	return true
}

func Hsl2Rgb(h float64, s float64, l float64) [3]int {
	h = h / 360
	s = s / 100
	l = l / 100

	var t2 float64
	var t3 float64
	var val float64
	var result [3]int

	if s == 0 {
		val = l * 255
		result[0] = int(math.Round(val))
		result[1] = int(math.Round(val))
		result[2] = int(math.Round(val))

		return result
	}

	if l < 0.5 {
		t2 = l * (1 + s)
	} else {
		t2 = l + s - l*s
	}

	t1 := 2*l - t2

	rgb := [3]float64{0, 0, 0}

	for i := 0; i < 3; i++ {
		t3 = h + 1.0/3.0*(-(float64(i) - 1))

		if t3 < 0 {
			t3++
		}

		if t3 > 1 {
			t3--
		}

		if (6 * t3) < 1 {
			val = t1 + (t2-t1)*6*t3
		} else if (2 * t3) < 1 {
			val = t2
		} else if (3 * t3) < 2 {
			val = t1 + (t2-t1)*(2.0/3.0-t3)*6
		} else {
			val = t1
		}
		rgb[i] = val * 255
	}

	result[0] = int(math.Round(rgb[0]))
	result[1] = int(math.Round(rgb[1]))
	result[2] = int(math.Round(rgb[2]))

	return result
}
