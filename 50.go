package main

import (
	"fmt"
	"time"
)

func main() {
	n1 := time.Now().UnixNano()/1000000
	fmt.Println(myPow(2,3))
	n2 := time.Now().UnixNano()/1000000
	fmt.Println(n2-n1)
}

func myPow(x float64, n int) float64 {
	if n == 0 {return 1}
	if(n<0){
		n = -n;
		x = 1/x;
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}else{
		return x*myPow(x*x,n/2)
	}
}
