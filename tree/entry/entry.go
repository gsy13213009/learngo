package main

import (
	"../../tree"
	"fmt"
)

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
}
