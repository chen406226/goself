// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//http.HandleFunc("/", handler) // each request calls handler
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))



	//ch := make(chan int) //这样就报错
	ch := make(chan int,1)

	log.Println(<-ch)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case ch <- i:
	//	case x := <-ch:
	//		fmt.Println(x) // "0" "2" "4" "6" "8"
	//	}
	//}
	select {
	case ch <- 2:
	case x := <-ch:
		fmt.Println(x) // "0" "2" "4" "6" "8"
	}
	time.Sleep(time.Second)
	return



	v := make(chan int)

	go func() {
		v<-1
	}()

	x := <-v
	log.Println(x)

	c :=make(chan int)
	go testDeadLock(c)
	go test(c)
	time.Sleep(time.Millisecond)

}
func test(c chan int){
	c<-'A'
}

func testDeadLock(c chan int){
	for{
		fmt.Println(<-c)
	}
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//!-
