package main

import (
	"fmt"
	"bytes"
	"./collection"
)

func main() {
	//sliceBase()
	//sliceBase_2()
	//copyAndAppend()
	exercise()
}

/**
用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，
第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。
 */
func exercise() {
	data := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	result := Filter_2(data, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(result)
	result.And(1)
	fmt.Println(result)
}
func Filter(arr []int, fun func(i int) bool) []int {
	result := make([]int, 0)
	for _, v := range arr {
		if fun(v) {
			result = append(result, v)
		}
	}
	return result
}

func Filter_2(arr []int, fun func(i int) bool) collection.List{
	list := collection.New()
	for _, v := range arr {
		if fun(v) {
			list.And(v)
		}
	}
	return *list
}

func copyAndAppend() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	dst := make([]int, 10, 50)
	copy(dst, src)
	//将src copy到dst的0 到 min(len(src),len(dst))的 位置上 （拷贝个数是 src 和 dst 的长度最小值）
	fmt.Println(dst)
	copy(dst[5:10], src)
	//将src copy到dst的 5 到 10的角标上
	fmt.Println(dst)
	dst = append(dst, src...)
	// append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片
	//如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储
	fmt.Println(dst)
	fmt.Println(len(dst))
	fmt.Println(cap(dst))
}

func sliceBase_2() {
	var arr = []int{0, 1, 2, 4, 5, 6 }
	fmt.Println(arr)
	//[0 1 2 4 5 6]
	changeArrItem_2(arr)
	//将切片传递给函数,可以在函数内修改数组的元素
	fmt.Println(arr)
	//[0 1 2 4 5 6]
	arr1 := make([]int, 5)
	//使用 make() 函数来创建一个切片 同时创建好相关数组
	arr1[0] = 10
	arr1[1] = 11
	fmt.Println(arr1)
	var buf = bytes.NewBufferString(" adf dsafds ")
	//var str ,_= buf.ReadString(0)
	fmt.Println(buf.String())
	b := bytes.TrimSpace(buf.Bytes())
	fmt.Println(string(b))
	fmt.Println("========================>")
	arr = []int{0, 1, 2, 4, 5, 6 }
	fmt.Println(arr)
	for i, v := range arr {
		arr[i] = v * 2
	}
	fmt.Println(arr)
	fmt.Println("========================>")
	arrF := []float32{12.3, 9.3, 14.1, 15.2, 3.5}
	sum, avg := sumAndAvg(arrF)
	min := min(arrF)
	max := max(arrF)
	fmt.Printf("sum=%f,avg=%f,min=%f,max=%f", sum, avg, min, max)
}

func min(arr []float32) (min float32) {
	min = arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return
}

func max(arr []float32) (max float32) {
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return
}

func sumAndAvg(arr []float32) (sum float32, avg float32) {
	if len(arr)<1{
		return
	}
	for _, v := range arr {
		sum += v
	}
	avg = sum / float32(len(arr))
	return
}

func changeArrItem_2(ints []int) {
	ints[0]=100
}
func sliceBase() {
	/**
	切片：一个切片是一个隐藏数组的引用，并且对于该切片的切片也引用同一个数组
		切片容量固定是因为数组的长度是固定的，切片的容量即隐藏数组的长度。长度可变指的是在数组长度的范围内可变
    s :=[] int {1,2,3 }
    直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
    s := arr[:]
    初始化切片s,是数组arr的引用
    s := arr[startIndex:endIndex]
    将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
    s := arr[startIndex:]
    缺省endIndex时将表示一直到arr的最后一个元素
    s := arr[:endIndex]
    缺省startIndex时将表示从arr的第一个元素开始
    s1 := s[startIndex:endIndex]
    通过切片s初始化切片s1
    s :=make([]int,len,cap) 其中 cap 是可选参数
    通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
    当相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片 同时创建好相关数组

	和数组语法上的区别：创建切片跟创建数组唯一的区别在于 Type 前的“ [] ”中是否有数字，为空，则代表切片，否则则代表数组。
						因为切片是长度可变的
	 */

	var arr = [6]int{0, 1, 2, 4, 5, 6 }
	var slice = arr[1:2]
	// arr[:]=(arr[0:len(arr)) arr[1:]=(arr[1:len(arr)) arr[:5]=(arr[0:5) arr[1:2]
	for i, v := range slice {
		fmt.Printf("i=%d,v=%d\n", i, v)
	}
	fmt.Printf("slice len=%d\n", len(slice))
	fmt.Printf("slice cap=%d\n", cap(slice))
	fmt.Println()
	slice = slice[0:len(arr)-1]
	for i, v := range slice {
		fmt.Printf("i=%d,v=%d\n", i, v)
	}
	fmt.Printf("slice len=%d\n", len(slice))
	fmt.Printf("slice cap=%d\n", cap(slice))
}
