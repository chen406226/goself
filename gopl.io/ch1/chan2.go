package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	//go fibonacci(c, quit)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	time.Sleep(time.Second*2)
}

//select的轮询机制
func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select { // select轮询机制
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}