// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import "fmt"

func main() {
	//var intChan chan int
	//intChan = make(chan  int, 3)
	//num3 := <- intChan
	//fmt.Println("Hello, 世界",num3)

	// deadlock
	//
	//遍历管道
	//intChan2 := make(chan int, 100)
	//for i := 0; i < 100; i++ {
	//	intChan2 <- i *2  //放入100个数据到管道
	//}
	intChan2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			intChan2 <- i *2  //放入100个数据到管道
		}
	close(intChan2)
	}()
	//遍历管道不能使用普通的for循环
	//for i := 0; i < len(intChan2); i++ {
	//
	//}
	//1)在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//2)在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	for v := range intChan2 {
		fmt.Println("v = ", v)
	}


	//
}

//!-
