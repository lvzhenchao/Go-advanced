package main

import "fmt"

func main() {

	//定义

	//方法1
	var p1 Person
	fmt.Println(p1)
	p1.name = "小明"
	p1.age = 21
	p1.sex = "男"
	p1.address = "北京"
	fmt.Printf("姓名：%s, 年龄：%s, 性别：%d, 地址：%s\n", p1.name,p1.age,p1.sex,p1.address)

	//方法2
	p2 := Person{}
	p2.name = "ruby"
	p2.age  = 18
	p2.sex  = "女"
	p2.address = "南京"
	fmt.Printf("姓名：%s, 年龄：%s, 性别：%d, 地址：%s\n", p2.name,p2.age,p2.sex,p2.address)

	//方法3  【常用】
	p3 := Person{
		name: "晓华",
		age:20,
		sex:"女",
		address:"天津",
	}
	fmt.Println(p3)

	//方法4
	p4 := Person{
		"小刚",
		24,
		"男",
		"成都",
	}
	fmt.Println(p4)


}

//定义结构体
type Person struct {
	name string
	age int
	sex string
	address string
}

//总结

//数据类型

//值类型：
//	int, float, bool, string, array
//
//引用类型：
//	slice, map, function, pointer