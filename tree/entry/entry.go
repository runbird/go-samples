//一个目录下 tree 下只能有一个包 （tree || main）
package main

//type 大写为public 小写 private
import "scy.com/runbird-go-samples/tree"
import "fmt"


//扩展他人实现，如 treeNode
type myTreeNode struct {
	node *tree.Node
}

//扩展，新增后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	//通过myTreeNode{}包装
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()

	fmt.Println()
	mynode := myTreeNode{&root}
	mynode.postOrder()

}
