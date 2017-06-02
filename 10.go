package main

import "fmt"

func main() {
	//isMatch("aa","a") → false
	//isMatch("aa","aa") → true
	//isMatch("aaa","aa") → false
	//isMatch("aa", "a*") → true
	//isMatch("aa", ".*") → true
	//isMatch("ab", ".*") → true
	//isMatch("aab", "c*a*b") → true
	s := "aaaaaaaaaaaaab"
	p := "a*a*a*a*a*a*a*a*a*a*c"
	fmt.Println(isMatch(s,p))
}

func isMatch(s string, p string) bool {
	if len(p) == 0 && len(s) == 0{
		return true
	}
	if len(p) == 0 {
		return false
	}
	if len(s) == 0 {
		//fmt.Println("4",p)
		if len(p)%2 != 0{
			return false
		}
		for i := 0 ; i < len(p)-1 ; i = i+2 {
			//fmt.Println("5",string(p[i]),string(p[i+1]))
			if p[i+1] != '*'{
				return false
			}
		}
		return true
	}
	if (s[0] == p[0] || p[0] == '.') && (len(p) == 1 || (len(p) > 1 && p[1] != '*')) {
		//fmt.Println("1:",s,p)
		return isMatch(s[1:], p[1:])
	}else if s[0] != p[0] && p[0] != '.' && len(p) > 1 && p[1] == '*'{
		//fmt.Println("2:",s,p)
		return isMatch(s,p[2:])
	}else if (s[0] == p[0] || p[0] == '.') && p[1] == '*'{
		//fmt.Println("3:",s,p)
		return isMatch(s[1:],p) || isMatch(s,p[2:])
	}
	return false
}
