package main

import "fmt"
import "strings"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3
	fmt.Println(lengthOfLongestSubstring("c")) // 1
	fmt.Println(lengthOfLongestSubstring("au")) // 2
	fmt.Println(lengthOfLongestSubstring("dvdf")) // 3
	fmt.Println(lengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ ")) // 95
}

func lengthOfLongestSubstring(s string) int {

	// split by spaces
	strArr := strings.Split(s, " ")

	// fmt.Println(strArr)

	if (len(s) > 0 && strArr[0] == "") || len(s)==1{
		return 1
	}

	tempLongStr := ""

	for i := 0; i < len(strArr); i++ {

		temp := findLongest(strArr[i])

		if len(temp) > len(tempLongStr) {
			tempLongStr = findLongest(strArr[i])
		}

	}
	result := len(tempLongStr)
	if len(strArr) > 1 {
		result++
	}
	fmt.Println(tempLongStr)
	return result

}

func findLongest(s string) string {
	tempOutput := ""
	longOutput := ""
	for i, v := range s {
		
		if !strings.Contains(tempOutput, string(v)) {
			tempOutput += string(v)
			if len(tempOutput) > len(longOutput) && i == len(s) - 1 {
				longOutput = tempOutput
			}
		} else {
			if len(tempOutput) > len(longOutput) {
				longOutput = tempOutput
			}
			if v == ' ' {
				tempOutput = ""
			} else {
				lastIndex := strings.LastIndex(s[:i], string(v))
				// fmt.Println(s[lastIndex:i], lastIndex)
				// tempOutput = string(v)			
				tempOutput=s[lastIndex:i]
			}

		}
	}
	// fmt.Println(longOutput)
	return longOutput
}

// Given a string s, find the length of the longest substring without repeating characters.
