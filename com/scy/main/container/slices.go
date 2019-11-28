package container

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 67, 8, 9, 10}
	//以下都为slice
	//result:[2 3 4 5]  左开右闭
	fmt.Println("arr[2:6] =", arr[2:6])
	//arr[:6] = [0 1 2 3 4 5]
	fmt.Println("arr[:6] =", arr[:6])
	//arr[2:] = [2 3 4 5 67 8 9 10]
	fmt.Println("arr[2:] =", arr[2:])
	//arr[:] = [0 1 2 3 4 5 67 8 9 10]
	fmt.Println("arr[:] =", arr[:])

	fmt.Println("arr[2:] =", updateSlice(arr[2:]))
	fmt.Println("arr[:] =", updateSlice(arr[:]))

	s2 := arr[2:]
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	extendingSlice()
	sliceAppend()
}

// slice
func updateSlice(arr []int) []int {
	arr[0] = 100
	return arr
}

// arr := [...] int{0, 1, 2, 3, 4, 5, 6,7, 8, 9, 10}
// s := arr[2:]
// s[0] = 10
// slice本身没有数据，是对底层array的一个view
// arr的值变为[0,1,10,3,4,5....]

//slice可以向后扩展，不能向前扩展
//s[i] 不能超越len(s),向后扩展不能超越底层数组cap(s)
func extendingSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Println("Extending Slice")
	fmt.Println("s1=", s1) //s1= [2 3 4 5]
	fmt.Println("s2=", s2) //s2= [5 6]
	//s1[2 3 4 5], len(s1)=4, cap(s1)=9
	fmt.Printf("s1%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s2[5 6], len(s2)=2, cap(s2)=6
	fmt.Printf("s2%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

}

func sliceAppend() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:6]
	s2 := s1[3:5]

	s3 := append(s2, 11)
	s4 := append(s3, 12)
	fmt.Println("s2,s3,s4 = ", s2, s3, s4)
	fmt.Println(arr)
}
