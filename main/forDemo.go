package main

import "fmt"

func main() {
	/**
		形式一：基于计数器的迭代
		for 初始化语句; 条件语句; 修饰语句 {}
	 */
	for i := 0; i < 25; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	var i = 0
	for ; i<10;i++  {
		//fmt.Println("")
	}

	/**
	形式二：基于条件判断的迭代
	for 条件语句 {}
	 */
	var j int = 5
	for j > 0 {
		j = j - 1
		if j<1 {
			break
		}
		fmt.Printf("The variable i is now: %d\n", j)
	}

	/**
	形式三：无限循环
	条件语句是可以被省略的，如 i:=0; ; i++ 或 for { } 或 for ;; { }（;; 会在使用 gofmt 时被移除）：这些循环的本质就是无限循环。
	最后一个形式也可以被改写为 for true { }，但一般情况下都会直接写 for { }。
	 */
	//for true {
	//	fmt.Println("--->")
	//}
	//for {
	//	fmt.Println("--->")
	//}
	//for ; ; {
	//	fmt.Println("--->")
	//}

	/**
	形式四：for-range 结构
	它可以迭代任何一个集合（包括数组和 map，详见第 7 和 8 章）。
	语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。一般形式为：for ix, val := range coll { }。
	 */
	var str = "abc def ghi"
	for i, c := range str {
		fmt.Printf("i=%d, c=%c\n", i, c)
	}

	/**
	break:
	一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。
	但在 switch 或 select 语句中（详见第 13 章），break 语句的作用结果是跳过整个代码块，执行后续的代码
	 */
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}

	/**
	continue:关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件
	 */
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
}
