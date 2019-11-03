package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:]
	fmt.Println("s1 = ", s1) // s1 =  [2 3 4 5 6 7]
	s2 := arr[:]
	fmt.Println("s2 = ", s2) // s2 =  [0 1 2 3 4 5 6 7]
	updateSlice(s1)
	fmt.Println("after update")
	fmt.Println("s1 = ", s1)   // s1 =  [100 3 4 5 6 7]
	fmt.Println("s2 = ", s2)   // s2 =  [0 1 100 3 4 5 6 7]
	fmt.Println("arr = ", arr) // arr =  [0 1 100 3 4 5 6 7]
	extending()
	appendFunc()
	sliceOps()
}
func updateSlice(arr []int) {
	arr[0] = 100
}
func extending() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1 = ", s1)                                            // s1 =  [2 3 4 5]
	fmt.Println("s1 = ", s1)                                            // s1 =  [2 3 4 5]
	fmt.Println("s2 = ", s2)                                            // s2 =  [5 6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) // s2=[5 6], len(s2)=2, cap(s2)=3
	fmt.Println("s1[3:] = ", s1[3:])                                    // s2 =  [5]
}
func appendFunc() {
	fmt.Println("appendFunc")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 14)
	fmt.Println("s1 = ", s1)                 // s1 =  [2 3 4 5]
	fmt.Println("s2 = ", s2)                 // s2 =  [5 6]
	fmt.Println("s3, s4, s5 = ", s3, s4, s5) // s3, s4, s5 =  [5 6 10] [5 6 10 11] [5 6 10 11 14]
	// s4 and s5 不再是对arr的view
	fmt.Println("arr = ", arr) // arr =  [0 1 2 3 4 5 6 10]
}
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d value=%v\n", len(s), cap(s), s)
}
func sliceOps() {
	// 创建slice
	var s []int // zero value， 默认为nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1) // len=4, cap=4 value=[2 4 6 8]
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2) // len=16, cap=16 value=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	printSlice(s3) // len=10, cap=32 value=[0 0 0 0 0 0 0 0 0 0]
	// 复制 slice
	copy(s2, s1)
	printSlice(s2) // len=16, cap=16 value=[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]
	// 删除元素
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2) // len=15, cap=16 value=[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]
	// 删除头
	front := s2[0]
	s2 = s2[1:]
	// 删除尾
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("front and tail", front, tail) // front and tail 2 0
	printSlice(s2)                             // len=13, cap=15 value=[4 6 0 0 0 0 0 0 0 0 0 0 0]
}
