package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	go sendData(ch1)

	//for循环的 for range形式可用于从通道接收值，直到它关闭为止，不需要单独判断
	for v := range ch1{
		fmt.Println("读取数据：", v)
	}
	fmt.Println("main..over.....")

}

func sendData(ch1 chan int)  {
	//发送方：10条数据
	for i:=0; i <10; i++ {
		time.Sleep(1*time.Second)
		ch1 <- i
	}
	close(ch1)
}
