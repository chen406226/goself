// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"fmt"
	"time"
)

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
)
func main() {
	ch := make(chan string)
	go func(ch chan string) {
		fmt.Println("kaishi")
		//for n:= range ch{
		//	fmt.Println(n)
		//}
		ch<- "sdfjksf"
		//一个无缓存的chan的发送操作讲导致发送者goroutine阻塞，知道被接受   所以无缓存在一个goroutine里（main）先发送是直接阻断的报错
		//当chan成功后，两个goroutine可以继续执行后面的语句，如果接受操作先发生，那么接收者goroutine也将阻塞，知道另一个goroutine在chan里执行发送操作
		//有缓存chan当发送超出范围了还没有被接受，则阻塞，在范围内不阻塞
		println("阻塞")
		//必须goroutine????
	}(ch)
	fmt.Println("jieshu")
	time.Sleep(time.Second*3)
	<-ch
	time.Sleep(time.Second)

	//n := <-ch
	//fmt.Println(n)
	//fmt.Println(m)
	//time.Sleep(time.Second)

	//counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	counts[input.Text()]++
	//}
	//// NOTE: ignoring potential errors from input.Err()
	//for line, n := range counts {
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}
}

//!-
