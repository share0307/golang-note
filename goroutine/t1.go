package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

/**
	为什么 goroutine panic recover，为什么跨 goroutine 失效？
 */
func main() {
	f,_ := os.Create("/usr/local/var/www/go/golang-note/goroutine/t1_trace.log")

	trace.Start(f)
	defer trace.Stop()

	go a()

	time.Sleep(3 * time.Second)
}

/**
	
 */
func a()  {
	defer func() {
		rec := recover()

		if rec != nil {
			fmt.Println("a defer:")
			fmt.Println(rec)
		}
	}()

	go b()
}

func b()  {
	defer func() {
		fmt.Println("pppppppppp defer:")

		rec := recover()

		if rec != nil {
			fmt.Println("b defer:")
			fmt.Println(rec)
		}
	}()

	panic("b panic:")
}
