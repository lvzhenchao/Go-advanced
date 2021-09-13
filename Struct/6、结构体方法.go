package main

import "fmt"

func main() {
	//方法只是一个函数，带有特殊的接收器类型；
	//	接收器可以是struct类型或非结构体类型

	//意义：某个类别的行为功能【有指向性】
	//		不同类别的行为名称可以相同

	//子类可以访问父类的属性和方法
	//子类可以新增自己的属性和方法
	//子类可重写父类的方法

	w1 := Worker{name:"小明", age:12, sex:"男"}
	w1.work()

	w2 := &Worker{name:"小王", age:13, sex:"女"}
	w2.work()

	w1.rest()
	w2.rest()
}

type Worker struct {
	name string
	age int
	sex string
}



//定义方法

//非指针类型的接收器
//代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效
//小对象由于值复制时的速度较快，所以适合使用非指针接收器
func (w Worker) work()  {
	fmt.Println(w.name,"在工作")
}

//指针类型的接收器，更接近于面向对象中的this或者self
//由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
//大对象因为复制性能较低，适合使用指针接收器, 在接收器和参数间传递时不进行复制，只是传递指针
func (p *Worker) rest()  {
	fmt.Println(p.name,"在休息")
}