package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("+123"))
	fmt.Println(myAtoi("0032"))
	fmt.Println(myAtoi("4193 with words"))
}

const MAXINT int = 1<<31 - 1
const MININT int = -1 << 31

func myAtoi(s string) int {
	if s == "" || strings.Trim(s, " ") == "" {
		return 0
	}
	str := stepOne(s)
	// fmt.Println(str)

	isPositive, str := stepTwo(str)
	// fmt.Println(isPositive, str)

	str = stepThree(str)

	resInt := stepFour(isPositive, str)

	return resInt
}

func stepFour(isPos bool, s string) int {
	if s == "" {
		return 0
	}
	res, _ := strconv.Atoi(s)
	if !isPos {
		res *= -1
	}
	if res < MININT {
		return MININT
	}
	if res > MAXINT {
		return MAXINT
	}
	return res
}

func stepThree(s string) string {

	for index, value := range s {
		if value < '0' || '9' < value {
			return s[:index]
		}
	}
	return s
}

func stepTwo(s string) (bool, string) {
	if s[0] == '-' {
		return false, s[1:]
	}
	if s[0] == '+' {
		return true, s[1:]
	}
	return true, s
}

func stepOne(s string) string {
	return strings.TrimLeft(s, " ")
}
