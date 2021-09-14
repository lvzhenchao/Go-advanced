package main

import "fmt"

func main() {
	//类型断言是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型
	//格式：
	//value, ok := x.(T)

	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, ",", ok)

	x = "hello"
	value1, ok1 := x.(int)
	fmt.Println(value1, ",", ok1)

	x = 2.2
	getType(x)
}


//类型断言配合switch
func getType(a interface{})  {
	switch a.(type) {
		case int:
			fmt.Println("整形")
		case string:
			fmt.Println("字符串")
		case float64:
			fmt.Println("浮点类型")
		default:
			fmt.Println("未知类型")

	}
}
