package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := "E:/GoPath/src/a/1.png"
	desFile := "E:/GoPath/src/a/222.png"
	//total,err := CopyFile1(srcFile, desFile)
	//total,err := CopyFile2(srcFile, desFile)
	total,err := CopyFile3(srcFile, desFile)
	fmt.Println(total, err)
}

func CopyFile3(srcFile, destFile string)(int, error)  {
	bs,err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0,err
	}
	err = ioutil.WriteFile(destFile,bs,0777)
	if err != nil {
		return 0,err
	}
	return len(bs), nil
}

func CopyFile2(srcFile, destFile string)(int64, error)  {
	file1,err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}

func CopyFile1(srcFile, destFile string)(int, error)  {
	file1,err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for{
		n,err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("报错了:", err)
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}

	return total,nil
}
