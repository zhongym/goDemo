package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	var now = time.Now() //获得当前时间
	fmt.Println(now) //2017-08-23 22:36:19.5468327 +0800 CST
	fmt.Printf("%d年%d月%d日 %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println()

	var nowStr = now.Format("2006-01-02 15:04:05") //日期格式化 （2006-01-02 15:04:05 是固定的）只能更改格式，不能更改值
	fmt.Println(nowStr)
	fmt.Println(now.Format(time.RFC822)) //常用的日期格式

	var tim ,_= time.Parse("2006-01-02 15:04:05","2018-08-23 22:00:00") //字符串转日期

	var isAfter = now.After(tim) //比较两个日期的大小
	var isBefore = now.Before(tim)
	fmt.Println(isAfter)
	fmt.Println(isBefore)

	var newNow = now.AddDate(0,0,1) //时间加上年，月 日
	fmt.Println(newNow)

	//var day time.Duration = 60 * 60 * 24 * 1e9 //纳秒  1,000,000,000 纳秒 = 1秒
	var day time.Duration = 24 * time.Hour //
	newNow = now.Add(day) //时间加上一天 （单位纳秒）
	fmt.Println(newNow)

	newNow = now.Add(-day)//时间加上减天 （单位纳秒）
	fmt.Println(newNow)

	time.Sleep(time.Second * 5) //当前线程睡5s
	fmt.Println("结束")

}
