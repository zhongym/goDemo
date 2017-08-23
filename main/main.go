package main

import (
	"fmt"
	"strings"
	"runtime"
	"os"
	"math"
)

const a int = 10
const a1,b1  = 1,2
const (
	a2,b2=1,2
)
const (
	b int = iota  //iota默认=0，每次调用+1，在新const重置为0
	c //如果常量列表第二个没写语句，会将会和第一行一样
	d
)

var intA int = 10
var intB =10 //自动类型推断
var strA string = "st,r";
func init()  {
	fmt.Println("inited")
	var arr []string = strings.Split(strA,",")
	fmt.Println(arr)
}

func main() {
	fmt.Printf("dsfafdsaf %d aa  %f",3,0.4)
	fmt.Println()
	fmt.Println(&intA)
	fmt.Println(&intB)

	fmt.Println("dsfsdf")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a1)
	fmt.Println(a2)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	method1()
}
func method1() {
	a := false
	b := true
	//c:=10
	d := a == b
	fmt.Println(d)
	var intA  = getInt()
	fmt.Println(intA)

	testInt()
}
func getInt() int {
	return 1
}

func testInt() {
	i,g := math.Modf(10.03)
	fmt.Println(i)
	fmt.Println(g)
}
