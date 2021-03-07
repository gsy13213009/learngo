package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				// 此操作不会交出控制权
				a[ii]++
				// 主动交出控制权
				runtime.Gosched()
				// print是io操作，会有协程之间的切换
				//fmt.Println("Hello from "+"goroutine ", i)
			}
		}(i)
	}
	fmt.Println(a)
	time.Sleep(time.Millisecond)
}
