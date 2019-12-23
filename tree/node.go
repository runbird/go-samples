package tree

import "fmt"

// 要改变内容必须使用指针接收者
// 结构过大也需要使用**
// 一致性，如有指针接收者都用**

//type TreeNode struct {
type Node struct {
	Value       int
	Left, Right *Node
}

//工厂函数 指定参数构造
func CreateNode(value int) *Node {
	//返回了局部变量的地址
	return &Node{Value: value}
}

// func ()接受者 method()
//接收者
//func Print(node TreeNode){fmt.Print(node.Value)}
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//指针 接收者
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting Value to nil node. Ignored")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	//中序遍历
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func main() {
	var root Node
	fmt.Println(root)
	fmt.Println(Node{Value: 3})

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	//报错
	//root.Left.Right = new(TreeNode)
	fmt.Println(root)

	root.Left.Right = CreateNode(2)
	fmt.Println(root)

	//func (node TreeNode) Print(){}
	root.Print()
	//func Print(node TreeNode){}
	//Print(root)
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	nodes := []Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	var nilRoot *Node
	//nil也可以调用 指针接收者
	nilRoot.SetValue(200)
	nilRoot = &root
	nilRoot.SetValue(300)
	nilRoot.Print()
}
