package main

import "fmt"

// 定义结构
type treeNode struct {
	value       int
	left, right *treeNode
}

// 值传递，会拷贝一份node
func (node treeNode) print() {
	fmt.Print(node.value)
}
// 只有使用*才能修改值，因为值传递的原因，只能传指针进来才行
func (node *treeNode) setValue(value int) {
	node.value = value
}
// node的中序遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 使用工厂函数代替构造函数
func createNode(value int) *treeNode {
	// 在go语言中，局部变量也可以返回给别人用，C++不行
	// 不需要知道到底是分配到堆上还是栈上
	return &treeNode{value: value} // 返回了局部变量的地址
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.setValue(9)
	root.print()
}
