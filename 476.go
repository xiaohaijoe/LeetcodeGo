package main

import "fmt"
func main() {
	fmt.Println(findComplement(5))
	//fmt.Println(5&1)
	//fmt.Println(0^0)
	//fmt.Println(1^1)
	//fmt.Println(5)
}

func findComplement(num int) int {
	var temp = num
	var mask = 0
	for temp>0 {
		mask = mask << 1 + 1;
		temp = temp >> 1;
	}
	return num^mask
}