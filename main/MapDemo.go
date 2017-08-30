package main

import (
	"fmt"
	"sort"
)

func main() {
	//baseMap()
	//mapVFunc()
	mapOper()
}

/**
测试键值对是否存在及删除元素
 */
func mapOper() {
	map1 := map[string]string{"a": "a", "b": "b", "c": "c"}
	if v, isPresent := map1["b"]; isPresent { //判断key是否存在
		fmt.Println(v)
	}

	delete(map1, "a") //删除key
	if _, isPresent := map1["a"]; !isPresent {
		fmt.Println("key:a不存在")
	}
	fmt.Println(map1)

	/**
		for-range 遍历map
	 */
	for key := range map1 { //只取key
		fmt.Printf("key is: %d\n", key)
	}
	for _, value := range map1 { //只取value
		fmt.Printf("value is: %d\n", value)
	}
	for key, value := range map1 {
		fmt.Printf("key:%s,value:%s\n", key, value)
	}

	/**
		map 类型的切片
	 */
	userMaps := make([]map[string]interface{}, 5)
	for i := 0; i < cap(userMaps); i++ {
		userMaps[i] = make(map[string]interface{})
		userMaps[i]["name"] = fmt.Sprintf("zym%d", i)
		userMaps[i]["age"] = 12 + i
	}
	for _, user := range userMaps {
		fmt.Println(user)
	}

	/**
	map 的排序:
	map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序，
	然后可以使用切片的 for-range 方法打印出所有的 key 和 value
	 */
	fmt.Println("-------------------->")
	map1 = map[string]string{"e": "e", "a": "a", "b": "b", "c": "c"}
	fmt.Println(map1)
	keys := make([]string, len(map1))
	index := 0
	for key := range map1 {
		keys[index] = key
		index++
	}
	sort.Strings(keys) //对key进行排序
	for _, key := range keys {
		fmt.Printf("key:%s,value:%s\n", key, map1[key])
	}
}

func baseMap() {
	/**
	map:
	key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。
		所以数组、切片和结构体不能作为 key，但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，
		这样可以通过结构体的域计算出唯一的数字或者字符串的 key。
	value 可以是任意类型的；通过使用空接口类型（详见第 11.9 节），我们可以存储任意值
		，但是使用这种类型作为值时需要先做一次类型断言

	map定义：
	var map1 map[keytype]valuetype
	var map1 map[string]int
	 */
	var map1 map[string]int
	map1 = make(map[string]int)
	//初始化一个空map  不能使用 new
	map1 = make(map[string]int, 15)
	//初始化一个空map,并设置初始化容量
	map1 = map[string]int{}
	//初始化一个空map
	map1 = map[string]int{"a": 1, "b": 2}
	//初始化一个指定元素的map
	map1["zym"] = 1
	//put
	v := map1["a"]
	//get
	fmt.Println(v)
	fmt.Println(map1)
}

/**
map值可以是任意类型的 ,可以存放函数，实现一个分支功能
 */
func mapVFunc() {
	map2 := map[string]func(i, j int) int{
		"+": func(i, j int) int {
			return i + j
		},
		"-": func(i, j int) int {
			return i - j
		},
		"*": func(i, j int) int {
			return i * j
		},
		"/": func(i, j int) int {
			return i / j
		},
	}
	fmt.Println(map2["+"](5, 2))
	fmt.Println(map2["-"](5, 2))
	fmt.Println(map2["*"](5, 2))
	fmt.Println(map2["/"](6, 2))
}
