package main

import "fmt"
func main() {
	fmt.Println(convert("PAYPALISHIRING",3))
}

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows <= 1 {return s}
	ret := ""
	var delta = 2*numRows - 2
	for i := 0 ; i < numRows ; i++{
		for j := i ; j < len(s) ; j = j+delta {
			ret = ret + string(s[j])
			//找出中间值
			mid := (j-i) + delta -i
			if i != 0 && i != numRows-1 && mid < len(s) {
				ret = ret + string(s[mid])
			}
		}
	}
	return ret
}

func convert2(s string, numRows int) string {
	if len(s) == 0 || numRows <= 1 {return s}
	rows := make([]string, numRows)
	row := 0
	delta := 1
	for i := 0 ; i < len(s) ; i++ {
		rows[row] = rows[row] + string(s[i])
		row = row + delta
		if row >= numRows {
			row = numRows - 2
			delta = -1
		}
		if row < 0 {
			row = 0
			delta = 1
		}
	}
	return func() string{
		ret := ""
		for i := 0 ; i < len(rows) ; i++ {
			ret = ret + rows[i]
		}
		return ret
	}()
}


