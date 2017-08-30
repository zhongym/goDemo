package main

import "fmt"

func main() {
	arrBase()

	var in *[5]int = new([5]int)
	fmt.Println(*in)
}
func arrBase() {
	/**
	数组:
	var arr [5]int
	长度为5的数组,其元素值依次初始化为：0
    [5] int {1,2,3,4,5}
    长度为5的数组，其元素值依次为：1，2，3，4，5
    [5] int {1,2}
    长度为5的数组，其元素值依次为：1，2，0，0，0 。在初始化时没有指定初值的元素将会赋值为其元素类型int的默认值0,string的默认值是""
    [...] int {1,2,3,4,5}
    长度为5的数组，其长度是根据初始化时指定的元素个数决定的
    [5] int { 2:1,3:2,4:3}
    长度为5的数组，key:value,其元素值依次为：0，0，1，2，3。在初始化时指定了2，3，4索引中对应的值：1，2，3
    [...] int {2:1,4:3}
    长度为5的数组，起元素值依次为：0，0，1，0，3。由于指定了最大索引4对应的值3，根据初始化的元素个数确定其长度为5

	 */

	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 10
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, item := range arr {
		fmt.Printf("--->index=%d,value=%d\n", i, item)
	}

	var arrStr [5]string = [...]string{"123","121","aa","aad","dfas"} //声明数组时直接赋值，使用 [...]语法
	fmt.Println(arrStr)


	fmt.Printf("=================>")
	/**
	Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针）：
	1、将一个arr1数组变量赋值给arr2时，会将arr1的数组拷贝一份给arr2,修改arr2的数据不影响arr1
	2、在函数中数组作为参数传入时，如 func1(arr2)，会产生一次数组拷贝，func1 方法内不会修改原始的数组 arr2

	如果你想修改原数组，那么 arr2 必须通过&操作符以引用方式传过来，例如  func1(&arr2）
	 */
	var newArr  = [...]int{1, 2, 3, 4, 5}
	var newArr1 [5]int = [...]int{2, 3, 4, 5, 6}
	newArr = newArr1
	newArr[0]=1000
	fmt.Println(newArr)
	fmt.Println(newArr1)
	fmt.Printf("=================>\n")

	var intArr [5]int = [...]int{1,2,3,4,5}
	fmt.Printf("原数组=>%v\n",intArr)
	changeArrItem(intArr) //在函数中数组作为参数传入时，会产生一次数组拷贝， changeArrItem方法内不会修改原始的数组
	fmt.Printf("changeArrItem=>%v\n",intArr)
	changeArrItem_prt(&intArr)
	fmt.Printf("changeArrItem_prt=>%v\n",intArr)
}
func changeArrItem_prt(ints *[5]int) {
	arr := *ints
	arr[0] = 100
	fmt.Printf("changeArrItem_prt=>%v\n",arr)
}
func changeArrItem(ints [5]int) {
	ints[0]=100
}
