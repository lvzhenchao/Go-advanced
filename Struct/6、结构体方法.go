package main

import "fmt"

func main() {
	//方法只是一个函数，带有特殊的接收器类型；
	//	接收器可以是struct类型或非结构体类型

	//意义：某个类别的行为功能【有指向性】
	//		不同类别的行为名称可以相同

	w1 := Worker{name:"小明", age:12, sex:"男"}
	w1.work()

	w2 := &Worker{name:"小王", age:13, sex:"女"}
	w2.work()

    w2.rest()
	w1.rest()
}

type Worker struct {
	name string
	age int
	sex string
}



//定义方法
func (w Worker) work()  {
	fmt.Println(w.name,"在工作")
}

func (p *Worker) rest()  {
	fmt.Println(p.name,"在休息")
}