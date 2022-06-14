// Given a string s, return the longest palindromic substring in s.
package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("babad")) // "bab". "aba" is also a valid answer
	fmt.Println(longestPalindrome("cbbd"))  // "bb"
}

func longestPalindrome(s string) string {

	return s
}
