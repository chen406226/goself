// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"time"
)

//!+
func f1(in chan int){
	fmt.Println(<-in)
}
func main() {
	out := make(chan int)
	go func() {
		fmt.Println("hou")
		out<-2
	}()
	fmt.Println("前")
	//所有必须有接受再有发送才可以不然报错
	<-out
	fmt.Println(out)
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick //阻塞
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
