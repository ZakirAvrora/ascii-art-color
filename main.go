package main

import (
	"log"
	"os"
	"strings"

	col "ascii-art/color/color"
	util "ascii-art/color/utilities"
)

func main() {
	m := util.Initialize() // Initializing the map that store runes' ascii-arts
	var colorcode string = ""
	flagColor := false
	flagColorIndex := false
	flagColorLetters := false
	var indexRange []int
	var indexLetters []int
	colorOption := ""

	if len(os.Args) > 4 || len(os.Args) < 2 {
		log.Fatalf("Invalid number of arguments\n" + util.ErrorMsg())
	}

	if len(os.Args[1:]) >= 2 {
		if len(os.Args[2]) > 8 && os.Args[2][:8] == "--color=" && os.Args[2][8:] != "" {
			colorcode = GetColorEscapeCode(os.Args[2][8:])
			flagColor = true
		} else {
			log.Fatalf(util.ErrorMsg())
		}
	}

	str1 := strings.ReplaceAll(os.Args[1], "\\n", "\n") // Put real newlinevalues
	str := strings.ReplaceAll(str1, "\r\n", "\n")

	if len(os.Args[1:]) == 3 && flagColor {
		colorOption = strings.ReplaceAll(os.Args[3], " ", "")
		if len(colorOption) < 3 {
			log.Fatalln("Invalid representation of index range for coloring")
		}
		if colorOption[0] == '{' && colorOption[len(colorOption)-1] == '}' && colorOption[1:len(colorOption)-1] != "" {
			indexLetters = util.GetIndexMapsForColoring(colorOption[1:len(colorOption)-1], str)
			flagColorLetters = true
		} else if colorOption[0] == '[' && colorOption[len(colorOption)-1] == ']' {
			flagColorIndex = true
			indexRange = util.GetIndexesForColoring(colorOption[1:len(colorOption)-1], str)
		} else {
			log.Fatalln("Invalid representation of index range for coloring")
		}
	}

	if !util.CheckRunes(str) {
		log.Fatalf("String contain invalid rune without possible ascii-art")
	}

	if os.Args[1] == "" {
		return
	}

	strArr := strings.Split(str, "\n") // Split input with newline seperator

	if flagColor && !flagColorIndex && !flagColorLetters {
		col.PrintWithColorAll(strArr, m, colorcode)
	} else if flagColor && flagColorIndex {
		col.PrintWithColorIndexRange(strArr, m, colorcode, indexRange)
	} else if flagColor && flagColorLetters {
		col.PrintWithColorLetterList(strArr, m, colorcode, indexLetters)
	} else {
		col.PrintWihoutColor(strArr, m)
	}
}

func GetColorEscapeCode(inputColor string) (escapeCode string) {
	color := strings.ToLower(strings.ReplaceAll(inputColor, " ", ""))
	if col.IsColorANSI(color) {
		escapeCode = col.ANSIColor(color)
	} else if col.IsColorRGB(color) {
		escapeCode = col.RGBColor(color)
	} else if col.IsColorHSL(color) {
		escapeCode = col.HSLColor(color)
	} else if col.IsColorHash(color) {
		escapeCode = col.HashColor(color)
	} else {
		log.Fatalln("Invalid color representation")
	}
	return
}
