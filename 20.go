package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	list := make([]byte,len(s))
	pos := 0
	for i := 0 ; i < len(s) ; i++ {
		switch s[i]{
		case '(','[','{':
			list[pos] = s[i]
			pos++
		case ')':
			pos--
			if pos < 0 || list[pos] != '('{return false}
			list[pos] = 0
		case ']':
			pos--
			if pos < 0 || list[pos] != '[' {return false}
			list[pos] = 0
		case '}':
			pos--
			if pos < 0 || list[pos] != '{' {return false}
			list[pos] = 0
		}
	}
	if list[0] != 0{return false}
	return true
}
