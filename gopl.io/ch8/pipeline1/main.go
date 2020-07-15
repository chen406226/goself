// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import (
	"fmt"
	"log"
)

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			log.Println("index",x)
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			log.Println(x)
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	//if <-squares>100 {
	//	return
	//}
	//for {
	//	if <-squares>100 {
	//		close(squares)
	//		close(naturals)
	//		return
	//	}
	//	fmt.Println(<-squares)
	//	// 1,9,25,49,81,121 输出不对
	//}
	for x:= range squares{
		if x>100 {
			//close(squares)
			//close(naturals)
			return
		}
		fmt.Println(x)
		//直接对can遍历
	}
}

//!-
