package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var romanMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"L":    50,
	"C":    100,
}

func romanToArabic(roman string) int {
	return romanMap[roman]
}
func arabicToRoman(arabic int) string {
	var result strings.Builder
	for arabic >= 100 {
		result.WriteString("C")
		arabic -= 100
	}
	if arabic >= 90 {
		result.WriteString("XC")
		arabic -= 90
	}
	for arabic >= 50 {
		result.WriteString("L")
		arabic -= 50
	}
	if arabic >= 40 {
		result.WriteString("XL")
		arabic -= 40
	}
	for arabic >= 10 {
		result.WriteString("X")
		arabic -= 10
	}
	if arabic == 9 {
		result.WriteString("IX")
		arabic -= 9
	}
	if arabic >= 5 {
		result.WriteString("V")
		arabic -= 5
	}
	if arabic == 4 {
		result.WriteString("IV")
		arabic -= 4
	}
	for arabic > 0 {
		result.WriteString("I")
		arabic -= 1
	}
	return result.String()
}
func isRoman(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) || (unicode.ToUpper(char) != 'I' && unicode.ToUpper(char) != 'V' && unicode.ToUpper(char) != 'X') {
			return false
		}
	}
	return true
}
func main() {
	for {
		var a, b string
		var act string
		var err = ""
		fmt.Fscanln(os.Stdin, &a, &act, &b, &err)
		if err != "" {
			panic("Error")

		}
		if isRoman(a) != isRoman(b) {
			panic("Error")

		}
		var X, Y int
		if isRoman(a) {
			X = romanToArabic(a)
		} else {
			X, _ = strconv.Atoi(a)
		}
		if isRoman(b) {
			Y = romanToArabic(b)
		} else {
			Y, _ = strconv.Atoi(b)
		}
		var result int
		switch act {
		case "+":
			if X > 10 || X < 1 || Y > 10 || Y < 1 {
				panic("Error")

			}

			result = X + Y
		case "-":
			if X > 10 || X < 1 || Y > 10 || Y < 1 {
				panic("Error")

			}
			result = X - Y
		case "*":
			if X > 10 || X < 1 || Y > 10 || Y < 1 {
				panic("Error")

			}
			result = X * Y
		case "/":
			if Y == 0 {
				panic("Div 0")
			}
			if X > 10 || X < 1 || Y > 10 || Y < 0 {
				panic("Error")

			}
			result = X / Y
		default:
			panic("Error")

		}
		if isRoman(a) && isRoman(b) {
			var e = arabicToRoman(result)
			if e == "" {
				panic("Error")

			}
			fmt.Println("=", arabicToRoman(result))
		} else {
			fmt.Println("=", result)
		}
	}
}
