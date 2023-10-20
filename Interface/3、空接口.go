package main

import "fmt"

func main() {
	//不包含任何方法的方法，所以任何类型都实现了空接口，因此空接口可以存储任意类型的数值

	//fmt包内的方法：func Println(a ...interface{}) (n int, err error) {}

	var a1 A = Cat{"花猫"}
	var a2 A = Person{"王二狗", 20}
	var a3 A = "花花"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test1(a1)
	test1(a2)
	test1(3.14)
	test1("ruby")

	test2(a1)

	//map，value可以是任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "小明"
	map1["age"] = 12
	map1["friend"] = Person{"小刚", 12}
	fmt.Println(map1)

	//切片，存储任意类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1,a1,a2,a3,a4,100, "abc")
	fmt.Println(slice1)
	
	test3(slice1)
}

func test3(slice2 []interface{})  {
	for i:=0; i<len(slice2); i++ {
		fmt.Printf("第%d个数据：%v\n", i+1, slice2[i])
	}
}

func test1(a A)  {
	fmt.Println(a)
}

func test2(a interface{})  {
	fmt.Println(a)
}

type A interface {

}

type Person struct {
	name string
	age int
}

type Cat struct {
	color string
}
