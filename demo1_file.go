package main

import (
	"fmt"
	"os"
)

func main() {
err := os.RemoveAll("E:/GoPath/src/a/cc/dd")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

fmt.Println("删除多层文件夹成功")



}
