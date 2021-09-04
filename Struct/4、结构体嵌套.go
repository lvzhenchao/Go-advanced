package main

import "fmt"

func main() {
	//结构体嵌套：一个结构体中的字段，是另一个结构体的类型
	

	s3 := Student{
		name:"小明",
		age: 12,
		book:Book{
			bookName:"十万个为什么",
			price:12.5,
		},
	}
	fmt.Println(s3)

	b4 := Book{"呼啸山庄", 12.1}
	s4 := Student2{name:"ruby", age:12, book:&b4 }
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println(s4.book)

}

type Book struct {
	bookName string
	price float64
}

type Student struct {
	name string
	age int
	book Book//值拷贝
}

type Student2 struct {
	name string
	age int
	book *Book//引用拷贝
}
