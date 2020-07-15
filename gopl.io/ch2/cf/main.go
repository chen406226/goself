// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(time.Second*2)
	//for _, arg := range os.Args[1:] {
	//	t, err := strconv.ParseFloat(arg, 64)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	//		os.Exit(1)
	//	}
	//	f := tempconv.Fahrenheit(t)
	//	c := tempconv.Celsius(t)
	//	fmt.Printf("%s = %s, %s = %s\n",
	//		f, tempconv.FToC(f), c, tempconv.CToF(c))
	//}
}
func test()  {
	//这里可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	//需要 make()
	myMap[0] = "golang" // error
}

//!-
