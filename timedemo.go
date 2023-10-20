package main

import (
	"fmt"
	"time"
)

func main() {

	//1获取现在的时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	//2获取指定的时间
	t2 := time.Date(2008,7,16,16,30,28,0,time.Local)
	fmt.Println(t2)

	//3格式化时间 time-->string之间的转换(里面的模板时间必须为 2006.1.2 15:04:05)
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	s2 := t1.Format("2006-1-2")
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	//string-->time日期类型（模板时间同上，且格式需相同于数据源）
	s3 := "1999年10月10日"//sting
	t3 ,err:= time.Parse("2006年1月2日", s3)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(t3)
	fmt.Printf("%T\n", t3)

	fmt.Println(t1.String())

	//4.根据当前时间，获取指定的内容
	year,month,day := t1.Date()
	fmt.Println(year,month,day)

	hour,min,sec := t1.Clock()
	fmt.Println(hour,min,sec)


	//5.时间戳：指定的的日期，距离1970年1月1日0点0时0分0秒的时间差值
	t4 := time.Date(1970,1,1,1,0,0,0,time.UTC)
	timeStamp1 := t4.Unix()
	fmt.Println(timeStamp1)

	timeStamp2 := t1.Unix()
	fmt.Println(timeStamp2)

	timeStamp3 := t4.UnixNano()
	fmt.Println(timeStamp3)

	timeStamp4 := t1.UnixNano()
	fmt.Println(timeStamp4)


	//6.时间间隔
	t5 := t1.Add(time.Minute)
	fmt.Println(t1)
	fmt.Println(t5)
	fmt.Println(t1.Add(24*time.Hour))

	//差值
	t6 := t5.Sub(t1)
	fmt.Println(t6)

	//7.睡眠
	time.Sleep(3 * time.Second)
	fmt.Println("main...over...")

}
