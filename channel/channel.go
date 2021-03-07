package main

import "fmt"

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2
}

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func main() {
	chanDemo()
}