package main

import (
	"fmt"
)
func main() {
	num1 := "9"
	num2 := "9"
	fmt.Println(multiply(num1,num2))
}

func multiply(num1 string, num2 string) string {
	len1 := len([]rune(num1))
	len2 := len([]rune(num2))
	product := make([]int,len1+len2)
	for i :=  0 ; i < len1 ; i++ {
		carry := 0
		for j := 0 ; j < len2 ; j++{
			n1 := num1[len1-i-1] - '0'
			n2 := num2[len2-j-1] - '0'
			sum := int(n1*n2) + product[i+j] + carry
			carry = sum/10
			product[i+j]=sum%10;
		}
		if(carry > 0){
			product[i+len2] = carry;
		}
	}
	start := 0
	for i:= len(product)-1 ; i >=0 ; i-- {
		if(product[i] != 0){
			start = i
			break;
		}
	}
	if start==-1 {
		return "0"
	}
	str := ""
	for i := start ; i>=0 ; i--{
		str = str + string(product[i]+'0');
	}
	return str
}
