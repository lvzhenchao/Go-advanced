package main

import (
	"fmt"
)

func main() {
	//匿名结构体和匿名字段

	//1、匿名结构体：没有名字的结构体  在创建匿名结构体时，同事创建对象【意义不大，只能用一次】
	S1 := struct {
		username string
		age int
	}{
		username: "张三",
		age:12,
	}
	fmt.Println(S1)

	//2、匿名字段：一个结构体的字段没有名字【默认把数据类型当成名字】【不能重复，会冲突】
	w2 := Worker{"李四", 13}
	fmt.Println(w2)
	fmt.Println(w2.int)
	fmt.Println(w2.string)



}

type Worker struct {
	//name string
	//age int
	string
	int
}

type Student struct {
	name string
	age int
}
