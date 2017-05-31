package main

import "fmt"

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	length := 1
	for i:=len(nums) ; i > 0 ; i--{
		length = length*i
	}
	result := make([][]int,length)
	permuteRecursive(nums,0,result)
	return result;
}

func permuteRecursive(nums []int,begin int,result [][]int){
	if begin >= len(nums){
		i := 0
		for true {
			if result[i] == nil{
				t := make([]int,len(nums))
				for j := 0 ; j < len(nums) ; j++{
					t[j] = nums[j]
				}
				result[i] = t
				break;
			}
			i++
		}
		return
	}
	for i:= begin ; i < len(nums) ; i++ {
		fmt.Println("i=",i,",begin=",begin,",nums=",nums)
		temp := nums[begin]
		nums[begin] = nums[i]
		nums[i] = temp
		permuteRecursive(nums,begin+1,result)

		temp = nums[begin]
		nums[begin] = nums[i]
		nums[i] = temp
	}
}
