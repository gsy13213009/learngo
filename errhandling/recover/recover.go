package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		// 判断r是不是error
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic("I don't know what to do")
		}
	}()
	b := 0
	a := 1 / b
	fmt.Println(a)
	//panic(errors.New("this is a error"))
}
func main() {
	tryRecover()
}
