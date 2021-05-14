package main

import "fmt"

func main() {
	t1a()
}

/**
	用户结构体
 */
type user struct {
	Name string
}

/**
	
 */
func t1a(){
	fmt.Println(t2b())
}

func t2b() user {
	return user {
		Name: "kkk",
	}
}
