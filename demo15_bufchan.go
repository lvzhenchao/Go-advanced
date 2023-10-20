package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch3 := make(chan string, 4)
	go sendData(ch3)

	for{
		v, ok := <-ch3
		if !ok{
			fmt.Println("读完了。。。", ok)
			break
		}
		fmt.Println("读取的数据是：", v)

	}

	fmt.Println("main...over...")
}

func sendData(ch chan string)  {
	for i:= 0; i<10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中发送的第 %d 个数据\n", i)
	}
	close(ch)
}
