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
			sum := int((num1[len1-i-1] - '0')*(num2[len2-j-1] - '0')) + product[i+j] + carry
			carry = sum/10
			product[i+j]=sum%10;
		}
		//if(carry > 0){
			product[i+len2] = carry;
		//}
	}
	str := ""
	for i := len(product)-1 ; i>=0 ; i--{
		if !(len(str) == 0 && product[i] == 0) {
			str = str+string(product[i]+'0')
		}
	}
	if str == ""{
		return "0"
	}else{
		return str
	}
}
