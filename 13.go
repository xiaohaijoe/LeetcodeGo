package main

import "fmt"
func main() {
	fmt.Println(romanToInt("XXIV"))
}

func romanToInt(s string) int {
	nums := map[byte]int{
		'I' : 1,
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
		'D' : 500,
		'M' : 1000,
	}
	sum := 0
	for i := len(s)-1 ; i >= 0 ; i-- {
		if(i == len(s)-1 || nums[s[i]] >= nums[s[i+1]]){
			sum = sum + nums[s[i]]
		}else{
			sum = sum - nums[s[i]]
		}
	}
	return sum
}
