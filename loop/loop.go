package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
	)

	printFile("aaa.txt")

	printForever()
}
func printForever() {
	for {
		fmt.Println("hello world!")
	}
}

func printFile(filename string) {
	//读取文件内容
	content, err := os.Open(filename)
	if err != nil {
		panic("there have problem to print file: " + filename)
	}
	//bufio.scanner扫描
	scanner := bufio.NewScanner(content)
	//类似 while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//二进制转换
func convertToBin(i int) string {
	result := ""
	for ; i > 0; i /= 2 {
		lsb := i % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
