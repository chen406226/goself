package pdf1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Main()  {
	//echo1()
	echo2()
	//dup1()
}
func echo1() {
	// go run main.go args1 args3 args3
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func echo2() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println("key:=",i)
		// 1
		s += sep + arg
		sep = " "
		//2
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(s)
}
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++ //初始化0  查询重复数
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Println(line,"==========",n)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
