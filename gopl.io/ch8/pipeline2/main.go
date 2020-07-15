// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			fmt.Println("xxxxxxxxxxxxxx=",x)
			naturals <- x
		}
		close(naturals)
		//close(naturals)
		//如果注释掉了， 那么报错
	}()

	// Squarer
	go func() {
		for x := range naturals {
			fmt.Println("yyyyyyyyyy==",x)
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//!-
