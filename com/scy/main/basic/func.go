package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	//println(
	//	eval(10, 2, "+"),
	//)
	//q, r := div(15, 4)
	//print(q, r)
	//
	//fmt.Println(div2(13, 2))
	//fmt.Print(evl(10, 3, "x"))

	if result, error := eval2(3, 4, "+"); error != nil {
		fmt.Println("error: ", error)
	} else {
		fmt.Print(result)
	}

	fmt.Print(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return int(
				math.Pow(float64(a), float64(b)))
		}, 3, 4))

	var a, b = 3, 4
	a, b = swap2(a, b)
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}

//推荐使用div
func div2(i int, i2 int) (q, r int) {
	q = i / i2
	r = i % i2
	return q, r
}

func div(a int, b int) (q, r int) {
	return a / b, a % b
}

func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operator : %s", op)
	}
}

func apply(op func(a, b int) int, a, b int) int {
	// 获取函数名,通过反射获取指针再获取名
	pointer := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(pointer).Name()
	fmt.Printf("calling function %s with args (%d,%d)", name, a, b)
	return op(a, b)
}

//为func apply准备的方法
func pow(a, b int) int {
	//math.Pow 返回的是float64,重写为int
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列求和
func sum(numbers ...int) int {
	sum := 0
	// i := range
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

//指针
func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}
