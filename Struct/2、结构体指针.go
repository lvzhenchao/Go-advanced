package main

import "fmt"

func main() {

	//1、值类型结构体
	p1 := Person{
		"小刚",
		24,
		"男",
		"成都",
	}
	fmt.Println(p1)
	fmt.Printf("%p, %T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p, %T\n", &p2, p2)

	p2.name = "哈哈"
	fmt.Printf("%p, %T\n", &p2, p2)

	//查看地址：得出是值传递
	//{小刚 24 男 成都}
	//0xc00001e080, main.Person
	//{小刚 24 男 成都}
	//0xc00001e140, main.Person
	//0xc00001e140, main.Person

	//2、结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p, %T\n", pp1,pp1)
	fmt.Println(*pp1)
	//0xc00001e080, *main.Person
	//{小刚 24 男 成都}

	pp1.name = "三哥"
	fmt.Println(pp1)
	fmt.Println(p1)
	//&{三哥 24 男 成都}
	//{三哥 24 男 成都}


	//3、使用函数new，专门用于创建某种类型的指针的函数【返回指针】【不是nil,空指针】【是0值】
	pp2 := new(Person)
	pp2.name = "jack"
	pp2.sex = "男"
	pp2.age = 21
	pp2.address = "成都"
	fmt.Println(pp2)

	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
	//0xc0000ac090
	//0



	
}


//定义结构体
type Person struct {
	name string
	age int
	sex string
	address string
}

//总结

//函数new()
//创建的不是nil，空指针
//它指向了新分配的类型的内存空间，里面存储的是0