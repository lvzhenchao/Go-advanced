package main

import (
	"fmt"
	"math"
)

func main() {
	//空接口
	//当一个函数的形参是interface{}
	//需要对形参进行断言，从而得到它的真实类型
	//语法：
	//<目标类型的值>,<布尔参数> := <表达式>.(目标类型)

	var t1 Triangle = Triangle{3,4,5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a,t1.b,t1.c)

	fmt.Println("--------")

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	fmt.Println("---------")

	var s1 Shape//定义接口类型的变量
	s1 = t1//将t1赋值给s1  只能访问自己的方法，不能访问实现类的字段
	s1.peri()
	s1.area()

	fmt.Println("----------")
	var s2 Shape
	s2 = c1
	s2.peri()
	s2.area()

	fmt.Println("----------")
	testShape(t1)
	testShape(c1)
	testShape(s1)

	fmt.Println("----------")
	//interface, ok := 接口对象.(type)
	getType(t1)
	getType(c1)
	getType(s1)

	//指针是单独的一种类型
	//断言是重新生成的数据，只不过数值是一样的

	var t2 *Triangle = &Triangle{3,4,2}
	fmt.Println("t2:%T, %p, %p\n", t2,&t2)
	getType(t2)
	getType2(t2)
	getType2(t1)

}

func getType2(s Shape)  {
	switch ins := s.(type) {
			case Triangle:
				fmt.Println("是三角形，三条边是：", ins.a, ins.b, ins.c)
			case Circle:
				fmt.Println("是圆形，半径是：", ins.radius)
			case *Triangle:
				fmt.Println("是三角形结构体指针：", ins.a, ins.b, ins.c)
	}
}

func getType(s Shape)  {
	//断言
	if ins, ok := s.(Triangle);ok {//判断s是不是三角行，如果是，返回三角形的实例【type Triangle struct】
		fmt.Println("是三角形，三条边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle);ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	} else if ins, ok := s.(*Triangle); ok{
		fmt.Printf("ins：%T, %p, %p\n", ins, &ins, ins)
		fmt.Printf("s：%T, %p, %p\n", s, &s, s)
		//&{3 4 2} 0xc000006030
		//ins：*main.Triangle, 0xc000006038
		//s：*main.Triangle, 0xc000040270
	} else {
		fmt.Println("未知属性")
	}
}

func testShape(s Shape)  {
	fmt.Printf("周长：%.2f, 面积：%.2f\n", s.peri(),s.area())
}

//1、定义一个接口
type Shape interface {
	peri() float64
	area() float64
}

//2、定义实现类：三角形
type Triangle struct {
	a,b,c float64
}

func (t Triangle) peri()float64  {
	return t.a + t.b + t.c
}
func (t Triangle) area()float64 {
	p := t.peri()
	s := math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return s
}

//3、定义实现类：圆形
type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}
func (c Circle) area() float64  {
	return math.Pow(c.radius, 2)*math.Pi
}


