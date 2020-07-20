// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	//conn, err := net.Dial("tcp", "192.168.127.103:7001")
	conn, err := net.Dial("tcp", "127.0.0.1:7009")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	fmt.Println("开始")
	go func() {
		fmt.Println("中间")
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("sfdsdfs","done")
		done <- struct{}{} // signal the main goroutine
	}()
	fmt.Println("jieshu")
	mustCopy(conn, os.Stdin)
	conn.Close()
	//写了会等待goroutine运行直到输出done 并且chan被赋值，不写主进程会直接中断，goroutine也中断不会输出done
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
