package main

import (
	"fmt"
)

func main() {
	//funcBase()
	//function1()
	//defer_()

	//将函数当作函数的参数
	//var str = strings.Map(mapFuc,"abc")
	//fmt.Println(str)
	//
	//str = strings.Map(func(i rune) rune { //匿名函数
	//	return i + 10
	//}, "abc")
	//fmt.Println(str)

	var i =1
	var fun = func() (func(b int) int) {
		return func(b int) int {
			fmt.Println(b)
			return b
		}
	}
	i++
	fmt.Println(i) //2
	fun1 := fun()
	fun1(2)
}
func mapFuc(i rune) rune {
	return i+10
}


func funcBase() {
	i := getIntI(1)
	fmt.Println(i)
	a, b := getIntIAndJ(1, 2)
	fmt.Println(a)
	fmt.Println(b)
	a, _ = getIntIAndJ_2(1, 2)
	//第二个返回值赋给了空白符 _，然后自动丢弃掉 (空白符用来匹配一些不需要的值，然后丢弃掉)
	fmt.Println(a)
	changeValue(&a, 20)
	//通过传递指针给函数，在函数内部修改当前函数变量的值
	fmt.Printf("修改后的值->%d", a)
	str := toString("zym", "1", "2", "3")
	fmt.Println(str)
	strArr := []string{"1", "3", "3", "4"}
	str = toString("0", strArr...)
	//如果参数被存储在一个数组 arr 中，则可以通过 arr... 的形式来传递参数调用变参函数。
	fmt.Println(str)
}

func getIntI(i int) int { //单返回值 返回值可以不用() 把它们括起来
	return i
}

func getIntI_2(i int) (j int) { //单命名返回值  只有一个命名返回值，也需要使用 () 括起来
	j = i
	return
}

func getIntIAndJ(i,j int) (int,int) { //多返回值  当需要返回多个非命名返回值时，需要使用 () 把它们括起来，比如 (int, int)
	return i,j
}

func getIntIAndJ_2(i,j int) (x2 int, x3 int) { //命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的return语句。
	x2 = i
	x3 = j
	// return x2, x3
	//return 0,2 	//即使函数使用了命名返回值，你依旧可以无视它而返回明确的值。
	return //我们只需要一条简单的不带参数的return语句。
}

/**
改变外部变量
传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回
 */
func changeValue(i *int,v int)  {
	*i = v
}

/**
可变参数的函数 java也有
可变参数的参数，只能放在函数的最后一个参数，调用时长度可以为 0
这个函数接受一个类似某个类型的 slice 的参数 ,该参数可以通过第 5.4.4 节中提到的 for 循环结构迭代
变长参数可以作为对应类型的 slice 进行二次传递:
function F1(s … string) {
	F2(s …)
	F3(s)
}

func F2(s … string) { }
func F3(s []string) { }
 */
func toString(prefix string, values ...string) string  {
	for _, str := range values {
		fmt.Printf("传入的参数->%s",str)
	}
	fmt.Println()
	return fmt.Sprintf("%s%s",prefix,values)
}

func function1() {
	fmt.Printf("function1->1\n")
	defer function2()
	fmt.Printf("function1->2\n")
}

func function2() {
	fmt.Printf("defer function1->1\n")
}

func defer_() {
	i := 0
	defer fmt.Printf("defer->%d\n",i) //defer->0 虽然defer语句是最后执行，但i的值是在defer语句之前的值，
	i++
	fmt.Printf("->%d\n",i) //->1
	/**
		out:
		->1
		defer->0
	 */
	return
}