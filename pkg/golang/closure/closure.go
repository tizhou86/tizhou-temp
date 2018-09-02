package main

import "fmt"

func main() {

	//anonymous function
	func(a int, b int) bool {
		return a > b
	}(1, 2)

	f := func(a int, b int) bool {
		return a > b
	}

	f(1,2)

	//closure based on anonymous function
	a := 1

	f2 := func()(func()){
		b := 1
		return func() {
			fmt.Println(a, b)
		}
	}

	f2()

	a = a * 2

	f2()
}