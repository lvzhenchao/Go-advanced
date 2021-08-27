package main

func main() {
	/*
		单向：单向
			chan <- T, 只支持写
			<- chan T, 只读
	*/

	//ch1 := make(chan int)//双向，可读可写
	//ch2 := make(chan <- int)//单向，只能写，不能读
	ch3 := make(<- chan int)//单向，只能读，不能写

	//ch2 <- 100
	//data := <- ch2

	ch3 <- 100
}
