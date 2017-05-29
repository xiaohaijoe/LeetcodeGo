package main

import "strconv"
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(strconv.Atoi("    010"))
	fmt.Println(myAtoi("+"));

}

func myAtoi(str string) int {
	result := 0
	beishu := 10
	sign := 1
	start := 0
	if len(str) == 0 {
		return result
	}
	for str[start] == ' ' && start < len(str) {
		start++;
	}
	if str[start] == '+' || str[start] == '-' {
		if str[start] == '+'{sign=1}
		if str[start] == '-'{sign=-1}
		start++
	}
	for i := start ; i < len(str) ; i++{
		num := int(str[i] - '0')
		if num < 0 || num > 9 {
			break;
		}
		if result > math.MaxInt32/10 || math.MaxInt32/10 == result && math.MaxInt32%10 < num {
			if sign == 1{
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*beishu + num;
	}
	return result*sign;
}