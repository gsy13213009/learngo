package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 可以返回多个返回值
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}
func div(a, b int) (q, r int) {
	return a / b, a % b
}
func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 4)
	fmt.Println(q, r)
	// calling function main.pow with args (3, 4)
	fmt.Println(apply(pow, 3, 4))
	// 匿名函数 calling function main.main.func1 with args (3, 4)
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	a, b := 5, 6
	//swap(&a,&b)
	a, b = swapBetter(a, b)
	fmt.Println(a, b)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 更好的实现方式
func swapBetter(a, b int) (int, int) {
	return b, a
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 函数式编程重写eval方法
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
