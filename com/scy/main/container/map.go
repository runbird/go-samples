package container

import "fmt"

func main() {
	//除了slice map function的内键类型都可以作为key
	m := map[string]string{
		"name":    "scy",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)

	var m3 map[string]string

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	//for k:= range m {
	//	fmt.Print(k)
	//}
	//
	//for _,v:= range m {
	//	fmt.Print(v)
	//}

	course := m["course"]
	fmt.Println(course)

	cause, ok := m["cause"]
	// shift + enter 直接跳到下一行
	fmt.Println(cause, ok) //  空+false
	if cause, ok := m["cause"]; ok {
		fmt.Println(cause)
	} else {
		fmt.Println("key doesn't exist")
	}

	delete(m, "site")
	name, ok := m["site"]
	fmt.Println(name, ok)
}
