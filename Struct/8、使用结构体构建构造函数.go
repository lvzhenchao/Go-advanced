package main

//Go语言的类型或结构体没有构造函数的功能，但是可以使用结构体初始化的过程类模拟实现构造函数

//模拟构造函数重载
type Cat struct {
	Color string
	Name string
}

//根据结构体的特性【字段】，构建不同种类的猫的实例

func NewCatByName(name string) *Cat  {
	return &Cat{
		Name:name,
	}
}

func NewCatByColor(color string) *Cat  {
	return &Cat{
		Color:color,
	}
}

