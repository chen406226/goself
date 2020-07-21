// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// The "+build ignore" tag (see p.295) excludes this file from the
// thumbnail package, but it can be compiled as a command and run like
// this:
//
// Run with:
//   $ go run $GOPATH/src/gopl.io/ch8/thumbnail/main.go
//   foo.jpeg
//   ^D
//
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopl.io/ch8/thumbnail"
)
//  不可在idea中运行 /go run main.go
//input m.jpg相对于当前运行目录而言
// D:\image.jpg 绝对路径目录
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}