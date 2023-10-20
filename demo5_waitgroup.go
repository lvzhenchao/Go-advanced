package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	/**
	WaitGroup: 同步等待组
		Add(), 设置等待组中要执行的子 goroutine的数量

		Wait(), 让主goroutine处于等待


		Done()，让等待组中的计数器减一，=在子goroutine中结束等待
	 */

	wg.Add(2)

	go fun1()
	go fun2()

	fmt.Println("main 进入阻塞状态，等待wg子goroutine结束")

	wg.Wait() //表示goroutine进入等待，意味着阻塞

	fmt.Println("main..结束阻塞")
}

func fun1()  {
	for i:= 1; i<10; i++ {
		fmt.Println("fun1()函数中打印...A", i)
	}

	wg.Done()//执行完会后，将计数器减1 wg.Add(-1)
}
func fun2()  {

	defer wg.Done()
	for j:= 1; j<10; j++ {
		fmt.Println("\tfun2()函数中打印...A", j)
	}
}
