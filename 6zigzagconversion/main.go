package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))

}

func convert(s string, numRows int) string {
	var temp string
	strArr := make([]string, numRows)
	divider := numRows + numRows - 1
	fmt.Println(divider)
	for i, v := range s {
		var arrIndex int
		if i%divider < numRows+1 {

			arrIndex = i % numRows
			strArr[arrIndex] += string(v)
			fmt.Println(arrIndex, string(v))
		} else {

			// arrIndex = numRows - (i+1)%numRows
		}
		// strArr[arrIndex] += string(v)

	}
	for _, v := range strArr {
		temp += v
	}
	fmt.Println(strArr)
	return temp
}
