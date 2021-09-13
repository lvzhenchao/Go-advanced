package main

import "fmt"

func main() {
	//工厂方法名字已new 或New开头

	//调用：
	f := NewFile(10, "./test.txt")
	fmt.Println(f)

	//禁止在外部使用new函数实例化
	wrong := new(matrix) //编译失败（matrix是私有的）
	fmt.Println(wrong)

	//强制使用工厂方法
	right := NewMatrix(params) //实例化的唯一方式
	fmt.Println(right)
}

type matrix struct {

}
func NewMatrix(params) *matrix {
	m := new(matrix) // 初始化 m
	return m

}

type File struct {
	fd int   //文件描述符
	name string //文件名
}

//创建结构体类型对应的工厂方法
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}
