package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))

}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var temp string

	strArr := make([]string, numRows)

	joining := numRows - 2

	fmt.Println(joining)

	for i, v := range s {

		var arrIndex int

		if i%(numRows+joining) < numRows {

			arrIndex = i % (numRows+joining)

		} else {

			arrIndex = numRows - ((i % (numRows+joining)) - numRows + 2)

		}

		strArr[arrIndex] += string(v)

	}

	for _, v := range strArr {
		temp += v
	}

	fmt.Println(strArr)

	return temp
}
