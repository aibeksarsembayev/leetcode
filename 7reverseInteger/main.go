package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(-10))
}

const MAXINT int = 1<<31 - 1
const MININT int = -1 << 31

func reverse(x int) int {

	// fmt.Println(MININT, MAXINT)
	var result int
	if x < 0 {
		result = reverseNegative(x)
	} else {
		result = reversePositive(x)
	}

	return result
}

func reverseNegative(x int) int {
	tempInt := x
	if x == MININT {
		tempInt += 1
	}
	tempInt *= -1
	result := 0
	check := 0

	for tempInt > 0 {
		if result > MAXINT/10 {
			return 0
		}
		result *= 10
		v := tempInt % 10
		tempInt /= 10

		if result > MAXINT-v {
			return 0
		}

		result += v
		if (x == MININT) && (check == 0) {
			result += 1
		}
		check++
	}

	if result == MININT {
		return 0
	}
	return 0 - result
}

func reversePositive(x int) int {
	tempInt := x
	result := 0
	for tempInt > 0 {

		if result > MAXINT/10 {
			return 0
		}
		result *= 10

		v := tempInt % 10
		tempInt /= 10

		if result > (MAXINT - v) {
			return 0
		}
		result += v
	}
	return result
}

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
