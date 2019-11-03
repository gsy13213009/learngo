package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外定义变量不可以使用:=，必须要var或者func
var aa = "a var" // 包级变量
var (
	hh     = "hh"
	bb     = true
	heihei = 11
)
// 方法内的变量必须使用，否则会报错
func variable() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s) // 0, ""
}

func variableType() {
	// 类型推断，可以写在一行
	var a, b, c, d = 1, "kll", "d", true
	// 最好的定义变量的形式
	e, f, g := "a", true, 22
	fmt.Println(a, b, c, d) // 1 kll d 1
	fmt.Println(e, f, g)    // a true 22
}

// 欧拉公式
func euler() {
	result := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Printf("%.3f\n", result)
}
func triangle() {
	var a, b = 3, 4
	var c int
	// 需要强制转换，没有隐式转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	fmt.Println(filename)
}
func enums() {
	// 普通枚举类型
	const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	// 自增枚举类型
	const (
		cpp1 = iota // 自增值
		_
		python1
		golang1
		javascript
	)
	// 使用自增值的表达式
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
}
func main() {
	fmt.Println("Hello word")
	variable()
	variableType()
	euler()
}
