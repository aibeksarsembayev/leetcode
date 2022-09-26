/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

    '.' Matches any single character.​​​​
    '*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).
*/

package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
}

func isMatch(s string, p string) bool {
	pIndex := 0
	checker := false
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]), string(p[pIndex]))
		if p[pIndex] == '.' || p[pIndex] == '*' {
			checker = true
		} else {
			if s[i] == p[pIndex] {
				checker = true
			} else {
				return false
			}

			if i == len(s)-1 && pIndex != len(p)-1 {
				return false
			}
			if i != len(s)-1 && pIndex != len(p)-1 {
				pIndex++
			}
		}

	}
	return checker
}
