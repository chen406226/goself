// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import "fmt"

//!+
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}
// 两个形参   chan<- int  定义赋值发送  <-chan int  存储 接收型不能发送
// 调用close 函数只能在发送者goroutine才能调用，不能在接受类型调用
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in chan<- int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go squarer(squares, naturals)
	go counter(naturals)
	printer(squares)
}

//!-
