package main

import "fmt"

func main() {
	var s []int //zero value for slice is nil
	for i := 0; i < 100; i++ {
		//nil
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	printSlice(s)

	s1 := []int{2, 4, 6}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	//len16 cap20
	s3 := make([]int, 16, 20)
	printSlice(s3)

	fmt.Println("Copying Slice")
	//target source
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	//s2[:3] + s2[4:]

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from back")
	tail := s2[len(s2)-2]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s))
}
