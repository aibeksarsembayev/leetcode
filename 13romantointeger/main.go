package main

import "fmt"

type Roman struct {
	I int
	V int
	X int
	L int
	C int
	D int
	M int
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	// declare roman nums as reference
	romanNum := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	roman := s
	intResult := 0

	for len(roman) > 0 { //iterate until no roman letter is left
		if len(roman) > 1 { // compare first and seond letters
			first := romanNum[string(roman[0])]
			second := romanNum[string(roman[1])]
			if first < second {
				intResult = intResult + second - first
				roman = roman[2:]
			} else if first == second {
				intResult = intResult + second + first
				roman = roman[2:]
			} else {
				intResult = intResult + first
				roman = roman[1:]
			}
		}
		if len(roman) == 1 {
			intResult = intResult + romanNum[string(roman[0])]
			roman = ""
		}
	}
	return intResult
}

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9.
    X can be placed before L (50) and C (100) to make 40 and 90.
    C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.
*/
