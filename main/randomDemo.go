package main

import (
	"math/rand"
	"fmt"
	"time"
)

type Date int

func (date Date) toString()  {
	fmt.Println(date)
}

func main() {
	var i int = 0
	for ; i < 10; i++ {
		var num = rand.Intn(100)
		fmt.Println(num)
	}

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	fmt.Println(time.Now().Month())

	var now Date = 1
	now.toString()



}
