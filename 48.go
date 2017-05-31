package main

import "fmt"

func main() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//fmt.Println(rotate(matrix))
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int)  {
	//1.对角线翻转,matrix[i][j] => matrix[j][i]
	for i:= 0 ; i < len(matrix) ; i++{
		for j := i ; j < len(matrix[0]) ; j++{
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	//2.中心翻转,matrix[i][j] => matrix[i][len(matrix.length-j-1)]
	for i := 0 ; i < len(matrix) ; i++{
		for j := 0 ; j < len(matrix[0])/2 ; j++{
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix[0])-j-1]
			matrix[i][len(matrix[0])-j-1] = temp
		}
	}
}
