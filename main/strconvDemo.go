package main

import (
	"strconv"
	"fmt"
)

func main() {
	var intSize = strconv.IntSize
	fmt.Println(intSize)

	var intStr string = strconv.Itoa(100) //整形转字符串
	fmt.Println(intStr)
	var intV, err = strconv.Atoi(intStr) //字符串转整形
	if err==nil {
		fmt.Println(intV)
	}
	intV, _ = strconv.Atoi(intStr) //字符串转整形 忽略错误（一个返回值）

	var floatStr string = strconv.FormatFloat(20.13131,'g',3,64)// 小数转字符串
	fmt.Println(floatStr)
	var floatV, _ = strconv.ParseFloat(floatStr, 64) //字符串转小数
	fmt.Println(floatV)

	var boolStr = strconv.FormatBool(true)
	var boolV, _ = strconv.ParseBool(boolStr)
	fmt.Println(boolV)

}
