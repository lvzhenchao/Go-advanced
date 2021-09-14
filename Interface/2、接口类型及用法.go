package main

import "fmt"

func main() {
	//非侵入性

	//接口最大的意义就是解耦合
	//突破性设计，最好的设计
	//实现多态：
		//针对子类来讲：不同形态可访问的字段属性和方法不同【具体行为功能不同】【子类的形态或父类的形态】；
			//-子类是一个特殊的父类类型，所有可以看做父类的对象
			//-子类可以看做子类对象

	//多态
	//	一个事物的多种形态
	//  	-看成实现本身的类型，能搞访问实现类中的属性和方法  usb = m1  usb = f1
	//  	-看成对应的接口类型，只能够访问接口中的方法

	//**************
	//接口的用法
	//	-一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的【任意实现类型对象】作为参数
	//	-定义一个类型为接口类型的变量，实际上可以赋值实现类的对象
	//**************

	//鸭子类型：针对动态语言；是一种类型推断策略；关注的是这个对象如何被使用；实现了动态语言的便利和静态语言类型检查



	//1、创建Mouse类型
	m1 := Mouse{"罗技小红"}
	fmt.Println(m1.name)

	//2、创建FlashDisk
	f1 := FlashDisk{"闪迪64G"}
	fmt.Println(f1.name)

	//用法一
	testInterface(m1)
	testInterface(f1)

	//用法二
	var usb USB
	//usb = m1
	usb = f1
	usb.start()
	usb.end()

	f1.delete()

	var arr [2]USB
	arr[0] = m1
	arr[1] = f1
	fmt.Println(arr)
}

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start()  {
	fmt.Println(m.name,"鼠标开始工作")
}

func (f FlashDisk)start()  {
	fmt.Println(f.name,"U盘开始工作")
}

func (m Mouse) end()  {
	fmt.Println(m.name,"鼠标结束工作")
}

func (f FlashDisk)end()  {
	fmt.Println(f.name,"U盘结束工作")
}

func (f FlashDisk) delete()  {
	fmt.Println(f.name,"u盘删除数据")
}


//接口用法1
//接口类型作为参数
//传入任意实现类型参数
func testInterface(usb USB)  {//USB = m1  USB = f1
	usb.start()
	usb.end()
}

