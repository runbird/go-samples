package basic

import (
	"fmt"
	"math"
)

//没有全局变量的含义，只有包内变量
var aa = 3
var ss = "suocaiyuan"
var bb = true

var (
	a1 = 31
	s1 = "suocaiyuan1"
	b1 = true
)

// 在函数外面必须要有var关键字等 不能使用 :=

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s) //%q 打出""  %s 空串
}

func variableInitiaValue() {
	//指定类型
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	//自动指定类型
	var a, b, c, s = 3, 4, true, "default"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//在函数内 第一次使用变量 用 :=  赋值
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

/*
	go中的常量
*/
//也可以定义在包外
const filename = "scy.txt"

func consts() {
	const a, b = 3, 4
	// sonst name ="scy"
	const name string = "scy.txt"
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, name)
}
func consts2() {
	const (
		a, b     = 3, 4
		filename = "scy.txt"
	)
}

/*
	枚举定义
*/
func enums() {
	const (
		java   = 0
		python = 1
		golang = 2
		cpp    = 3
	)
	fmt.Println(java, python, golang) // 0 1 2
}

//自增枚举
func enums2() {
	const (
		//iota 自增值
		java = iota
		_
		golang
		cpp
		javaScript
	)
	//灵活使用 iota
	const (
		b = 1 << (10 * iota) // 1* 10左移 iota
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(java, javaScript, golang, cpp) //0 4 2 3
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb, a1, s1, b1)
	consts()
	enums()
	enums2()
}
