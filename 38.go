package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	s := "1"
	for i := 1 ; i < n ; i++{
		temp := ""
		for j := 0 ; j < len(s) ; {
			k := j
			count := 1;
			for k+1 < len(s) && s[k] == s[k+1] {
				count++
				k++
			}
			temp = temp+strconv.Itoa(count)+string(s[k])
			j = j+count
		}
		s = temp
	}
	return s
}
