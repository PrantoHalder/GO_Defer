package main 

import "fmt"

func defer1 (){
	defer fmt.Println("This is defer")
	fmt.Println("This will work after this")
}