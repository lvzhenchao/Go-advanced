package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := "E:/GoPath/src/a/ab.txt"
	//打开文件,注意是否有此文件

	//file, err := os.Open(fileName)
	//file, err := os.Create()
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)


	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	//file, err := os.Open("writeAt.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(os.Stdout)
	//writer.ReadFrom(file)
	//writer.Flush()

	//ioutil.ReadFile()

	//关闭文件
	defer file.Close()

	//写入数据
	//bs := []byte{65,66,67,68,69,70}
	bs := []byte{97,98,99,100}
	//n,err := file.Write(bs)
	n,err := file.Write(bs[:2])
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	//直接写出字符串
	n,err = file.WriteString("Hello world")
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	//切片也可以直接写入字符串
	n,err = file.Write([]byte("totay"))
	fmt.Println(n)
	HandleErr(err)
}

func HandleErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}
