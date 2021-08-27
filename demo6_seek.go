package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "E:/GoPath/src/a/aa.txt"
	//打开文件,注意是否有此文件

	//file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd)
	file.WriteString("ABC")




}
