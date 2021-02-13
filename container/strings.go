package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes一二三四五!"

	fmt.Println(len(s)) // 19获取字节数
	fmt.Printf("%s\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s)) //9 获取字符数

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//rune 相当于go的char
	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%s)", i, ch)
		fmt.Println()
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
