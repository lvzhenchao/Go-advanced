package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	srcFile := "E:/GoPath/src/a/1.png"
	//获取文件名
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1: ]
	fmt.Println(destFile)
	temFile := destFile + "temp.txt"
	fmt.Println(temFile)

	file1,err := os.Open(srcFile)
	HandleErr(err)
	file2,err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleErr(err)
	file3,err := os.OpenFile(temFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	//1: 先读取临时文件的数据，再seek
	file3.Seek(0, io.SeekStart)//从0开始读
    bs := make([]byte, 100,100)//一个文件的长度
    n1,err := file3.Read(bs)
    //HandleErr(err)
    countStr := string(bs[:n1])
    count,err := strconv.ParseInt(countStr, 10, 64)
    //HandleErr(err)
    fmt.Println(count)

    //2：设置偏移量；设置读，写的位置
    file1.Seek(count, io.SeekStart)
    file2.Seek(count, io.SeekStart)
    data := make([]byte, 1024, 1024)//用来复制文件
    n2 := -1 //读取的数据量
    n3 := -1 //写出的数据量
    total := int(count) //读取的总量

	//3: 复制文件
	for{
		n2,err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("复制完毕", total)
			file3.Close()
			os.Remove(temFile)
			break
		}

		n3,err = file2.Write(data[:n2])
		total += n3

		//将复制的总量，存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		fmt.Print("total:%d\n", total)


		//假装断电
		if total > 8000{
			//panic("假装断电看看")
		}
	}




}

func HandleErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}
