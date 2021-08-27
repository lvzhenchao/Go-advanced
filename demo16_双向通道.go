package main

import (
	"fmt"
)

func main() {
	/*
	双向通道：
		chan
			chan <- data,发送数据，写出
			data <- chan,获取数据，读取

	单向通道：定向通道
		chan <- T, 只支持写
		<- chan T, 只支持读

	*/

	ch1 := make(chan string)

	done := make(chan bool)
	go sendData(ch1, done)

	data := <- ch1 //读取
	fmt.Println("子goroutine传来的数据是：", data)

	ch1 <- "我是main的数据" //发送数据

	<- done//因为阻塞，所以会等待子goroutine的 通道done发送数据
	fmt.Println("main...over...")

	//time.Sleep(1)

}

func sendData(ch1 chan string, done chan bool)  {
	ch1 <- "我是lzc" //发送

	data := <- ch1 //读取数据
	fmt.Println("main goroutine传来的数据：", data)

	done <- true
}
