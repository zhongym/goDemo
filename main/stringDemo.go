package main

import (
	"fmt"
	"strings"
)

func main() {
	//base()
	stringsDemo()
}

func stringsDemo() {
	var str = "ABC1231"
	var isAStart = strings.HasPrefix(str, "A")
	fmt.Println(isAStart)

	var is3End = strings.HasSuffix(str, "3")
	fmt.Println(is3End)

	var isContains = strings.Contains(str, "1")
	fmt.Println(isContains)

	isContains = strings.ContainsAny(str, "1od") //只要有一个字符匹配就返回true
	fmt.Println(isContains)

	isContains = strings.ContainsRune(str, '1') // 判断字符串 s 中是否包含字符 r
	fmt.Println(isContains)

	var index = strings.Index(str, "1")
	fmt.Println(index)

	index = strings.LastIndex(str, "1")
	fmt.Println(index)

	index = strings.IndexRune(str, '1')
	fmt.Println(index)

	var newStr = strings.Replace(str, "1", "40", 1) //n=前几个old 如果str中出现三个old,n=2就会将前2个替换成new n-1时替换全部
	fmt.Println(newStr)

	var count = strings.Count(str, "1")
	fmt.Println(count)

	newStr = strings.Repeat(str, 2) //将str重复n次返回
	fmt.Println(newStr)             //ABC1231ABC1231

	var a = "A"
	a = strings.ToLower(a) //转小写
	fmt.Println(a)
	a = strings.ToUpper(a) //转大写
	fmt.Println(a)

	a = "  123  "
	a = strings.TrimSpace(a) //去掉前后空格
	fmt.Println(a)

	a = "！23！"
	a = strings.Trim(a, "！") //剪掉前后指定的字符
	fmt.Println(a)
	a = "[1231]"
	a = strings.TrimLeft(a, "[") //去掉前面指定字符
	fmt.Println(a)
	a = strings.TrimRight(a, "]") //去掉后面指定字符
	fmt.Println(a)
	a = "[1231]"
	a = strings.TrimPrefix(a, "[") //去掉前面指定字符
	fmt.Println(a)
	a = strings.TrimSuffix(a, "]") //去掉后面指定字符
	fmt.Println(a)

	a = "2 1\v3 1\n2 3 1\r3"             //
	var arr []string = strings.Fields(a) //将以一个或者多个空白符分割str
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i] + "-")
	}
	fmt.Println()

	a = "1,2,3,4,5"             //
	arr = strings.Split(a, ",") //将以一个或者多个空白符分割str
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i] + "-") //1-2-3-4-5-
	}
	fmt.Println()

	a = strings.Join(arr, "-") //拼接
	fmt.Println(a)

	var reader = strings.NewReader(a)
	var rune, size, err = reader.ReadRune()
	fmt.Println(rune)
	fmt.Println(size)
	fmt.Println(err)
}

func base() {
	var str string = "adfa\nadfa34"
	//解析型字符串
	fmt.Println(str)
	var str1 string = `this is string \n\t`
	//非解析型字符串原样输出
	fmt.Println(str1)
	fmt.Println(len(str1))
	var str2 string
	fmt.Println(len(str2))
	fmt.Println(str == str1)
	var str3 string = "123123" + "adfasf" +
		"-->sdafasf"
	fmt.Println(str3)
}
