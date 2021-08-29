package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4 //x可以看成一个空接口类型
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("----------------")

	//根据反射的值，来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())

	fmt.Println("----------------")
	var num float64 = 1.23
	//从接口值得到反射对象
	value := reflect.ValueOf(num)
	//反射可以从反射对象获得接口值
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)


	pointer := reflect.ValueOf(&num)
	convertPointer := pointer.Interface().(*float64)//注意类型统一性
	fmt.Println(convertPointer)

}
