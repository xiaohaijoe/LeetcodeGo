package main

import "fmt"

func main() {
	nums := []int{1,3,5,6}
	fmt.Println(searchInsert(nums,5))
}

func searchInsert(nums []int, target int) int {
	for i:= 0 ; i < len(nums) ; i++ {
		if (i == len(nums)-1 && nums[i] < target) || (i < len(nums)-1 && nums[i+1] >= target && nums[i] < target) {return i+1}
	}
	return 0
}
