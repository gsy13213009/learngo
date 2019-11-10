package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// return的函数对sum的引用，形成了一个闭包
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + .... + %d = %d\n", i, a(i))
		// 0 + 1 + .... + 0 = 0
		// 0 + 1 + .... + 1 = 1
		// 0 + 1 + .... + 2 = 3
		// 0 + 1 + .... + 3 = 6
		// 0 + 1 + .... + 4 = 10
		// 0 + 1 + .... + 5 = 15
		// 0 + 1 + .... + 6 = 21
		// 0 + 1 + .... + 7 = 28
		// 0 + 1 + .... + 8 = 36
		// 0 + 1 + .... + 9 = 45
	}
}
