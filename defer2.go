package main

import "fmt"

func defer2 () {
	defer fmt.Println("defer is working with error")
	a := 10
	b := 0
	c:= a/b
	fmt.Println(c)
}