package main

import "fmt"

func main() {
	var arr1 [6]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1, arr2, arr3)

	//2行3列
	var grid [2][3]int
	fmt.Println(grid)

	//遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//推荐变量
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//只想获取数值，通过 _ 省略变量，不仅是range,任何地方都可以
	//如果只想要i 可以使用 for i:= range numbers
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//求最大值
	numbers := [...]int{-8, 121, 43, 1, 9, 0. - 9, 43, 3}
	maxi := -1
	maxValue := -1
	for i, v := range numbers {
		// if numbers[i] > maxValue
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)

	//求最大值 只关心值
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	//未改变arr1 arr3
	printArray(arr1)
	printArray(arr3)

	chageArray(&arr1)
	chageArray2(arr2[:])
	chageArray2(arr3[:])
}

// go 只有值传递,调用arr只会拷贝数组 func不改变原有的arr
func printArray(arr [6]int) {
	for i := range arr {
		fmt.Print(i)
	}
	fmt.Println()
}

// 使用指针,改变了arr
func chageArray(arr *[6]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// arr []int 切片，不是数组
//使用slice,改变了 arr
func chageArray2(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
