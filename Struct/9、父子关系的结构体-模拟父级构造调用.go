package main

import "fmt"

func main() {

	cat := NewCat("老猫")
	fmt.Println(cat)

	black := NewBlackCat("黑色")
	fmt.Println(black)

	
}

type Cat struct {
	Color string
	Name string
}

type BlackCat struct {
	Cat //嵌入Cat, 类似于派生
}

//构造基类
func NewCat(name string) *Cat  {
	return &Cat{
		Name:name,
	}
}

//构造子类
func NewBlackCat(color string) *BlackCat  {
	cat := &BlackCat{}
	cat.Color = color
	cat.Name = "猫是" + color
	return cat
}



