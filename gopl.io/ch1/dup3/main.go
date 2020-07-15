// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
//$$ go run main.go C:\Users\Administrator\Desktop\test.txt C:\Users\Administrator\Desktop\te.txt
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		fmt.Println(len(strings.Split(string(data), " ")))
		//\n 不能通过空格区分， 用字符串才可以
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n::::::::::::::::", n, line)
		}
	}
}

//!-
