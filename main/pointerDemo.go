package main

import "fmt"

func main() {
	//pointerBase()
	var i = 10
	fmt.Println(i)
	changVaule(&i, 20) //通过指针的传递，可以在别的函数中修改当前函数的局部变量的 值
	fmt.Println(i)
}
func changVaule(i *int, newInt int) {
	*i = newInt
}

func pointerBase() {
	var i int = 10
	var iPtr *int = &i
	//取地址符是 &  一个指针变量可以指向任何一个值的内存地址
	fmt.Printf("%d,%p", i, iPtr)
	//%p  指针的格式化标识符
	fmt.Println()
	fmt.Printf("%d", *iPtr)
	//符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；
	// 这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。
	fmt.Println()
	var b = i == *iPtr
	//对于任何一个变量 var， 如下表达式都是正确的：var == *(&var)
	fmt.Println(b)
	var str = "12313"
	fmt.Println(str)
	var strPrt *string = &str
	*strPrt = "adfasf"
	//通过对 *p 赋另一个值来更改“对象”，这样 原变量引用的值 也会随之更改
	fmt.Println(str)
	fmt.Println(*strPrt)


	var iPtr1 *int
	fmt.Println(*iPtr1) //对一个空指针的反向引用是不合法的，并且会使程序崩溃：
}
