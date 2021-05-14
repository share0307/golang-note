package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//fmt.Println(t2a())
	fmt.Println(t2b())
}

func t2a() (num int32) {
	num = 0

	defer func() {
		atomic.AddInt32(&num, 1)

		recover()
	}()

	//panic(atomic.AddInt32(&num, 1))

	panic("aaa")

	return
}

func t2b() int32 {
	var num int32= 0

	defer func() {
		atomic.AddInt32(&num, 1)

		//recover()
	}()

	//panic(atomic.AddInt32(&num, 1))

	return num
}
