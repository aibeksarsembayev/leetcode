// Given a string s, return the longest palindromic substring in s.
package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("babad")) // "bab". "aba" is also a valid answer
	fmt.Println(longestPalindrome("cbbd"))  // "bb"
	fmt.Println(longestPalindrome("a"))  // "a"
	fmt.Println(longestPalindrome("ac"))  // "a"
	fmt.Println(longestPalindrome("bb"))  // "bb"
	fmt.Println(longestPalindrome("aacabdkaca"))  // "bb"
	fmt.Println(longestPalindrome("abcda"))  // "bb"



}

func longestPalindrome(s string) string {	

	tempStr := ""
	length := len(s)

	for i:=0; i< length-1; i++{ // iterate from 1st letter one by one 

		tempStr2 := ""

		if len(tempStr) > len(s[i:]) { // if current found palindrome lentgh longer than rest string part, break to decrease time
			break
		}

		for j:=length-1; j>0; j--{ // iterate from end for each part of 1st loop
			
			if s[i] == s[j] && i < j { // if 1st and last current substring letters are matched, then go inside to check each letter
				counter := 0
				check := 0
			
				for counter < len(s[i:j+1])/2 { // iterate half of length of substring and compare each mirrowed letters
					if s[i+counter] == s[j-counter] { //if match then count checker
						check++
					}
					counter++ //counter to stop in the middle of substring
				}

				if check == len(s[i:j+1])/2 { // if half letters length and checker counter are equal then it means letters are mirrowed as palindrom
					tempStr2 = s[i:j+1] //assign current substring to temp2string
				}
				
				if len(tempStr) < len(tempStr2) { // if temp2string are longer than current temp result, then overwrite
					tempStr = tempStr2			
				}
			}
		}
	}

	if tempStr == "" { // if no palindrom is found, then assign first letter
		tempStr = string(s[0])
	}
	
	return tempStr 
}

