package main

import "fmt"

func main() {
	//嵌套：
	//	模拟继承：匿名字段  【提升字段，可以直接访问】
	//  聚合关系：非匿名字段【不能提升字段，不能直接访问】

	

	//1、父类
	p1 := Person{name:"张三", age: 12}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	//2、子类
	s1 := Student{Person{"小明", 14}, "美迪康大学"}
	fmt.Println(s1)

	//提升字段
	s1.name = "ruby"
	s1.age = 19
	fmt.Println(s1)
}

//1、定义父类
type Person struct {
	name string
	age int
}

//2、定义子类
type Student struct {
	Person //匿名字段，继承
	school string //子类的新增属性
}


//提升字段：匿名结构体的字段为提升字段，可以被直接访问【简写】