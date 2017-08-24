package main

import "fmt"

func main() {

	/**
	形式1
	switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
	}
	 */
	var i = 1
	switch i {
	case 1:
		fallthrough //如果不加fallthrough ，当i=1时，是不会进入case:2的
		//Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		// 但是可以使用fallthrough强制执行后面的case代码，
	case 2, 3:
		fmt.Println("1")
		fallthrough // fallthrough不会判断下一条case的expr结果是否为true
	case 4:
		fmt.Println("2")
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	/**
	形式二：
	switch {
	case condition1:
		...
	case condition2:
		...
	default:
		...
	}
	 */

	var a = 10
	switch {
	case a > 1:
		fmt.Println("a>1")
		fallthrough
	case a > 22 && a < 4:
		fmt.Println("a>2&&a<4")

	}

	/**
	形式三：
	switch initialization {
	case val1:
		...
	case val2:
		...
	default:
		...
	}
	 */
	switch a := getA(); {
	case a >= 1:
		fmt.Println("a>1")
	}

}
func getA() int {
	return 1
}
