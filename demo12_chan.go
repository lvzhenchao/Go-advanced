package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i:=0; i<10; i++ {
			fmt.Println("子goroutine中，i:", i)
		}
		//循环结束，向通道中写数据，表示要结束了。。。
		ch1 <- true
		fmt.Println("结束")
	}()


	//如果主进程先抢占资源，会去chan中读数据，如果没有数据，会阻塞下去..
	data := <- ch1
	fmt.Println("main...data-->:", data)
	fmt.Println("main主进程结束...")
}
