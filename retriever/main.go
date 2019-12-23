package main

import (
	"fmt"
	"scy.com/runbird-go-samples/retriever/mock"
	"scy.com/runbird-go-samples/retriever/real"
)

//实现者
type Retriever interface {
	Get(url string) string
	//	Post(url string,body string) string
}

//使用者
func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"hello golang"}
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(download(r))
}
