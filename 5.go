package main

import (
	"fmt"
)
func main() {
	s := "dfdfadsjfaklsdjf;lasjfldsaffadsfasdfaghadsfbabad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {return s}
	var start = 0
	var maxLen = 1
	for i := 0 ; i < len(s) ; {
		k := i
		j := i
		for k < len(s)-1 && s[k+1] == s[k] {
			k++
		}
		i = k + 1
		for k < len(s)-1 && j > 0 && s[k+1] == s[j-1] {
			k++
			j--
		}
		newLen := k-j+1
		if newLen > maxLen {
			start = j
			maxLen = newLen
		}
	}
	return s[start:start+maxLen]
}
