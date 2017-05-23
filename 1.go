package main

import (
	"fmt"
	"errors"
)
func main() {
	//fmt.Println("123455")
	nums := []int{2,7,11,15}
	res := twoSum(nums,9)
	fmt.Println("%d",res);
}

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for i,d:=range nums{
		m[d] = i
	}
	for i:=0 ; i < len(nums) ; i++{
		if j := target-nums[i]; 0 != m[j]{
			return []int{i,m[j]}
		}
	}
	panic(errors.New("no two sum solutiojn"))
}
