package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "E:/GoPath/src/a/aa.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	//可以用defer先关闭掉，以防漏写，造成资源泄露
	defer file.Close()

	//ioutil.ReadDir()
	ioutil.ReadFile()
	ioutil.WriteFile()
	ioutil.Discard()
	ioutil.TempDir("", "go-build")

	//读取数据
	bs := make([]byte, 4, 4)
	n := -1
	for  {
		n, err = file.Read(bs)
		if  n == 0 || err == io.EOF{
			fmt.Println("到了文件末尾，结束读取...")
			break
		}
		fmt.Println(string(bs[:n]))

	}
	////第一次读取
	//n, err := file.Read(bs)
	//fmt.Println(err)//<nil>
	//fmt.Println(n)//4
	//fmt.Println(bs)//[97 115 100 102]返回的是每个元素的字节编码数值
	//fmt.Println(string(bs))//asdf
	//
	////第二次读取
	//n, err = file.Read(bs)
	//fmt.Println(err)//<nil>
	//fmt.Println(n)//4
	//fmt.Println(bs)//[103 114 121 117]返回的是每个元素的字节编码数值
	//fmt.Println(string(bs))//gryu
	//
	////第三次读取
	//n, err = file.Read(bs)
	//fmt.Println(err)//<nil>
	//fmt.Println(n)//2
	//fmt.Println(bs)//[105 111 121 117]返回的是每个元素的字节编码数值
	//fmt.Println(string(bs))//ioyu
	//
	//
	////第四次读取
	//n, err = file.Read(bs)
	//fmt.Println(err)//EOF
	//fmt.Println(n)//0



}
