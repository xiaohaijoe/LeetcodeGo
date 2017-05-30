package main

import "fmt"

func main() {
	haystack := "jjlsldkfds"
	needle := "dk"
	fmt.Println(strStr(haystack,needle))
}

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {return 0}
	for i:=0 ; i < len(haystack) ; i++ {
		n := i
		m := 0
		for n < len(haystack) && m < len(needle) && haystack[n] == needle[m] {
			n++
			m++
		}
		if m == len(needle) {
			return i
		}
	}
	return -1
}
