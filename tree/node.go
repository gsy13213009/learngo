package tree

import "fmt"

// 定义结构
type Node struct {
	Value       int
	Left, Right *Node
}

// 值传递，会拷贝一份node
func (node Node) Print() {
	fmt.Print(node.Value)
}
// 只有使用*才能修改值，因为值传递的原因，只能传指针进来才行
func (node *Node) SetValue(value int) {
	node.Value = value
}
// node的中序遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 使用工厂函数代替构造函数
func CreateNode(value int) *Node {
	// 在go语言中，局部变量也可以返回给别人用，C++不行
	// 不需要知道到底是分配到堆上还是栈上
	return &Node{Value: value} // 返回了局部变量的地址
}