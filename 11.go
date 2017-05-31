package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{3,2,3,4,5,4,4}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	var res float64 = 0
	l := 0
	r := len(height)-1
	for l<r {
		var min float64 = math.Min(float64(height[l]),float64(height[r]))
		res = math.Max(res,float64(min*float64(r-l)))
		if(height[l] < height[r]){
			l++
		}else{
			r--
		}
	}
	return int(res)
}
