package main

import "fmt"
func main() {
	fmt.Println(intToRoman2(9))
}

func intToRoman(num int) string {
	res := ""
	list := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	chs := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	for i := 0 ; i < len(list) ; i++{
		for true {
			if(num - list[i] >= 0){
				num = num - list[i]
				res = res + chs[i]
			}else{
				break;
			}
		}
	}
	return res
}

func intToRoman2(num int) string{
	list := []string{
		"","I","II","III","IV","V","VI","VII","VIII","IX",
		"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC",
		"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM",
		"","M","MM","MMM","MMMM",
	}
	return list[num/1000+30] + list[(num/100)%10+20] + list[(num/10)%10+10] + list[num%10]
}
