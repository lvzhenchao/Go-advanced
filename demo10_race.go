package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10

var mutex sync.Mutex

var wg sync.WaitGroup
func main() {


	wg.Add(4)
	go sale("售票口1")
	go sale("售票口2")
	go sale("售票口3")
	go sale("售票口4")

	wg.Wait()//main要等待
	fmt.Println("程序结束了")
	//time.Sleep(5*time.Second)
}

func sale(name string){
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for  {

		mutex.Lock()
		if ticket > 0  {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Microsecond)
			fmt.Println(name,"售出：", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println("售完票了")
			break
		}

		mutex.Unlock()
	}

}
