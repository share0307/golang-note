package main

import (
	"fmt"
	"sync"
)

/**
	模拟
 */
func mockSendToServer(url string)  {
	fmt.Println("Server Url：", url)
}

func main() {
	urls := [...]string{"192.168.1.1:7000", "192.168.1.1:7001", "192.168.1.1:7002"}

	wg := sync.WaitGroup{}

	for _,url := range urls{
		wg.Add(1)
		go func() {
			mockSendToServer(url)

			defer wg.Done()
		}()
	}

	wg.Wait()
}