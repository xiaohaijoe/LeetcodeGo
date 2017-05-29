package main

import "fmt"
func main() {
	str := []string{
		"aa",
		"a",
	}
	fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if(len(strs) == 0){
		return prefix
	}
	for i := 0 ; i < len(strs[0]) ; i++{
		c := strs[0][i];
		for _,v :=range strs {
			if i > len(v)-1 || c != v[i]{
				return prefix
			}
		}
		prefix = prefix + string(c);
	}
	return prefix
}
