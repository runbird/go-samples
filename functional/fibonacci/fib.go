package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	f := fibonacci()
	scan(f)
}

// 1,1,2,3,5,8,13 ...
//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

//函数实现接口 io.Reader
func (g intGen) Read(p []byte) (n int, e error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func scan(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
