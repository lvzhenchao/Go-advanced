package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取goroot的目录
	fmt.Println("GOROOT:", runtime.GOROOT())
	//获取操作系统
	fmt.Println("os:",runtime.GOOS)
	//获取cpu的数量
	fmt.Println("逻辑cpu的数量：",runtime.NumCPU())

	//设置go程序执行的最大cpu的数量：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU());
	fmt.Println(n)

	//gosched, 让出cpu
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go1...")
		}
	}()

	for i:= 0; i<4; i++ {
		//让出时间片，先让别的goroutine执行
		runtime.Gosched()
		fmt.Println("main...")
	}

	runtime.Goexit()
	//终止当前的goroutine
}


