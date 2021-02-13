package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//readFile()
	//readFile2()
	fmt.Println(
		rank(59),
		rank(60),
		rank(70),
		rank(-1), //使用panic 中断程序的执行，报错不继续执行
		rank(101),
	)
}

func readFile() {
	const filename = "aaa.txt"
	contents, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println(error)
	} else {
		//contents为byte[] 使用%s打印出来
		fmt.Printf("%s\n", contents)
	}
}

func readFile2() {
	const filename = "aaa.txt"
	//if var1,var2 := function(); contion {}
	if contents, error := ioutil.ReadFile(filename); error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//不需要break,默认break,除非使用fallthrough
func rank(a int) string {
	rank := "?"
	switch { //switch之后可以没有表达式
	case a < 60:
		rank = "F"
	case a < 80:
		rank = "D"
	case a < 90:
		rank = "B"
	default:
		panic(fmt.Sprintf("wrong scores : %d", a))
	}
	return rank
}

//没有默认情况，不需要default
func rank2(a int) string {
	rank := "?"
	switch {
	case a < 0 || a > 100:
		panic(fmt.Sprintf("wrong scores : %d", a))
	case a < 60:
		rank = "F"
	case a < 80:
		rank = "D"
	case a < 90:
		rank = "B"
	}
	return rank
}

func eval(a, b int, op string) int {
	//result := 0
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("wrong function with option:" + op)
	}
	return result
}
