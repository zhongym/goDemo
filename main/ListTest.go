package main

import (
	"./collection"
	"fmt"
)
func main() {
	list := collection.New()
	for i := 0; i < 100; i++ {
		list.Add(i)
	}
	fmt.Println(list)
	fmt.Println(list.Get(0))
	fmt.Println(list.Set(0,100))
	fmt.Println(list.Get(0))
	fmt.Println(list.Contains(2))
	fmt.Println(list.Contains(200))
	list.Clear()
	fmt.Println(list)

	var list1 collection.ArrayList
	list1.Add(100)
	list.AddAll(list1)
	fmt.Println(list)
}
