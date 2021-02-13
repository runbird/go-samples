package tree

import "fmt"

//func (node *Node) traverse() {
//	//中序遍历
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}
//函数式编程
func (node *Node) traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
