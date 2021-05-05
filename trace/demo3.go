package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

/**
	模拟
 */
func mockSendToServer3(url string)  {
	fmt.Println("Server Url：", url)
}

func main() {
	test()
}

func test()  {
	// 创建一个记录trace日志的文件
	f, err := os.Create("demo3-trace.log")

	if err != nil {
		panic(err)
	}
	// 关闭文件
	defer f.Close()

	// 启动 stace 记录
	err = trace.Start(f)

	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	urls := [...]string{"192.168.1.1:7000", "192.168.1.1:7001", "192.168.1.1:7002"}

	wg := sync.WaitGroup{}

	for _,url := range urls{
		wg.Add(1)
		go func() {
			mockSendToServer3(url)

			defer wg.Done()
		}()
	}

	wg.Wait()
}