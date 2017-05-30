package main

import "fmt"

func main() {
	nums := []int{1,1,2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	length := 0
	step := 1;
	for i := 0 ; i < len(nums) ; i = i+step{
		j := i
		for j < len(nums)-1 && nums[j+1] == nums[j] {j++}
		step = j - i + 1
		nums[length] = nums[i]
		length++
	}
	return length
}
