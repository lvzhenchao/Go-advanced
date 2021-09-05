package main

import "fmt"

func main() {
	//接口和类型的实现关系是 非嵌入式

	//1、创建Mouse类型
	m1 := Mouse{"罗技小红"}
	fmt.Println(m1.name)

	//2、创建FlashDisk
	f1 := FlashDisk{"闪迪64G"}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)
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


//测试方法
func testInterface(usb USB)  {
	usb.start()
	usb.end()
}

