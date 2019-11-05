package main

import (
	"../../tree"
	"fmt"
)
// 使用组合的方式，扩展tree.Node
type myNode struct {
	node *tree.Node
}

func (mynode *myNode) postTraverse() {
	if mynode == nil || mynode.node == nil {
		return
	}
	left := myNode{mynode.node.Left}
	left.postTraverse()
	// 必须提成变量，否则无法取地址
	right := myNode{mynode.node.Right}
	right.postTraverse()
	mynode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.SetValue(9)
	root.Print()
	fmt.Println()
	root.Traverse()
	fmt.Println()
	// 使用组合的方式
	mynode := myNode{&root}
	mynode.postTraverse()
}
