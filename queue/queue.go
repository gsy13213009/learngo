package queue

// 使用别名的方式扩展类
type Queue []int

func (q *Queue) Push(v int) {
	// 可以用指针的方式，改掉指针指向的对象
	// 每次运算后，q会指向新的对象
	*q = append(*q, v)
}
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
