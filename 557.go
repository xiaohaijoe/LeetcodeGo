package main

import "fmt"
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	//start := 0
	//end := 0
	//str := ""
	str := make([]byte,len(s))
	temp := ""
	index := 0
	for i:= 0 ; i < len(s) ;i++{
		space := ""
		divide := false
		if(s[i] == ' '){
			space = " "
			divide = true
		}
		if(i == len(s) - 1){
			space = ""
			temp = temp + string(s[i]);
			divide = true
		}
		if divide {
			for j := len(temp)-1 ; j >= 0; j-- {
				//fmt.Println(string(str)," ",index)
				str[index] = temp[j]
				index++
			}
			if(space == " "){
				str[index] = ' '
				index++
			}
			temp = ""
		}else{
			temp = temp + string(s[i]);
		}
	}
	return string(str)
}
