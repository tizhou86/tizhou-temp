package main

import "fmt"

func main(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("runtime error caught")
		}
	}()

	panic("runtime error")
}
