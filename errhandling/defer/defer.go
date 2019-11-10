package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	// defer是栈结构，先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	// 3
	// 2
	// 1
}

// 打印结果是从30到0，因为defer是先进后出的
func testDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print too many")
		}
	}
}

func writeFile(filename string) {
	// If there is an error, it will be of type *PathError.
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// 错误处理，针对PathError或者其他error
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 100; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	tryDefer()
	testDefer()
	writeFile("fib.txt")
}
