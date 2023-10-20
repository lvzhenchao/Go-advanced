package main

import "fmt"

func main() {
	//多继承

	//Cat类到底是哪个接口的实现，就看定义变量的时候声明的是【哪个接口类型】的

	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("--------")
	var a1 A = cat
	a1.test1()
	fmt.Println("--------")

	var b1 B = cat
	b1.test2()
	fmt.Println("--------")

	var c1 C = cat
	c1.test1()
	c1.test2()
	c1.test3()
	fmt.Println("--------")

	//var a2 C = a1 //不行，a1只实现了test1(),一个方法；其他两个方法没有实现
	var a2 A = a1
	a2.test1()

}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {//如果想实现接口C，那不止要实现接口C的方法，还要实现接口A、B的方法
	A
	B
	test3()
}

type Cat struct {

}
func (c Cat)test1()  {
	fmt.Println("test1()...")
}
func (c Cat)test2()  {
	fmt.Println("test2()...")
}
func (c Cat)test3()  {
	fmt.Println("test3()...")
}


