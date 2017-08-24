package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int= getI(1)
	fmt.Println(i)
	var a = 10
	fmt.Println(a)
	if a := 0; a < i { // if 可以包含一个初始化语句（如：给一个变量赋值）  变量的作用域只能在if 和else 的｛｝中 ,
						//如果变量在 if 结构之前就已经存在，那么在 if 结构中，该变量原来的值会被隐藏
		fmt.Println(i)
	}else{
		fmt.Println(a)
	}

	if a := getI(1); a > 20 {// 可以在 if初始化语句中调用函数
		//.....
	}

	//最常用的用法
	var iA, err = strconv.Atoi("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iA)

	if iB, err := strconv.Atoi("1"); err != nil {
		fmt.Println(iB)
	}

}
func getI(i int) int {
	if i > 1 {
		return 100
	} else {
		return 200
	}
}
