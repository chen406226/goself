// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
	"time"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
//最多执行20个信号量。

var m = 0

func dey()  {
	fmt.Println(m,time.Now())
	tokens <- struct{}{} // acquire a token
	m++
	time.Sleep(time.Second*3)
	<-tokens // release the token
	//有缓存的管道队列会等待 释放通道  然后继续执行
	fmt.Println(m,time.Now())
}


var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	//for i := 0; i<30;i++  {
	//	go dey()
	//}
	//time.Sleep(time.Second*10)
	//return


	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		fmt.Println("ni")
		list := <-worklist
		fmt.Println("外面")
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				fmt.Println("中间")
				go func(link string) {
					fmt.Println("i里面")
					worklist <- crawl(link)
				}(link)
			}
		}
	}
	fmt.Println("nsdffsdfdsi")

}

//!-
