package main

import (
	"fmt"
	"strings"
	"log"
	"runtime"
)

func main() {
	/**
	可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点
	 */

	var i = 10
	var Add = Add(i)
	var a = Add(10)
	fmt.Println(a) //20
	a = Add(20)
	fmt.Println(a) //30

	fmt.Println("==================>")
	var add2 = Add2()
	a = add2(1) //1
	fmt.Println(a)
	a = add2(10)
	fmt.Println(a) //11 再次都是拿一样的函数引用调用函数， 因为闭包的原因，第二次调用Add2()方法时，变量i=第一次调用之后的结果1

	add2 = Add2()
	a = add2(10)
	fmt.Println(a) //拿到新的函数引用，Add2函数中的i会初始化为0

	fmt.Println("==================>")
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file")) // returns: file.bmp
	fmt.Println(addJpeg("file")) // returns: file.jpeg

	/**
	打印调试信息：使用runtime.Caller(1) 或者 log 包中的特殊函数来实现这样的功能
	 */
	log.SetFlags(log.Llongfile)
	log.Print()

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func Add2() (func(b int) int) {
	var i = 0
	return func(b int) int {
		i += b
		return i
	}
}

func Add(a int) (func(b int) int) {
	return func(b int) int {
		return a + b
	}
}
