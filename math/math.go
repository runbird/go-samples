package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func eular1() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}
func eular2() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(1i*math.Pi + 1)
	//取小数点前三位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	//只有强转，没有隐式转换，所以 C = int（）
	// var a,b = 3,4   c=(math.Sqrt((a*a + b*b)) 不用转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	/*	cannot convert "<=========c:" (type untyped string) to type int
		invalid operation: "<=========c:" + c (mismatched types string and int)
		fmt.Println("<=========c:"+ c)
	*/
	fmt.Println(c)
}

func main() {
	eular1()
	eular2()
	triangle()
}
