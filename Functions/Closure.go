package main

import "fmt"

func main() {

	name := "Ringo"
	f := func() {
		fmt.Printf("Hello %s", name)
	}
	f()

	funcVariable := func() string {
		return "Test"
	}

	sayHello(funcVariable)
}

func sayHello(fn func() string) {
	fmt.Println(fn())
}


func sayHello2(func() string) {
	return "Hello"
})

