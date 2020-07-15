package main

import (
	"fmt"
	"time"
)

func main() {
	//使用select可以解决从管道取数据的阻塞问题
	//1. 定义一个管道10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 4; i++ {
		intChan <- i
	}
	//2. 定义一个管道5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	//问题：在实际开发中，可能不好确定什么时间关闭该管道
	//可以使用select方式解决
	//label:
	for {
		select {
		//注意：这里intChan一直没有关闭，不会一直阻塞而deadlock
		//会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了，程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			return
			//break label
		}
	}
}