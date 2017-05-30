package main

import "fmt"

func main() {
	nums := []int{3,2,2,3}
	fmt.Println(removeElement(nums,3))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	length := 0
	for i:= 0 ; i < len(nums) ; i++ {
		if nums[i] != val{
			nums[length] = nums[i]
			length++
		}
	}
	return length;
}