// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 146.

// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"fmt"
	"log"
	"time"
)

//!+main
func bigSlowOperation() {
	fmt.Println("开始")
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	fmt.Println("结束")
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
	//defer  会在本函数结束执行时候执行， 所以休眠10s
}

func trace(msg string) func() {
	fmt.Println("中间")
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/
