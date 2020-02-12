package main

import "fmt"

func echo(input string) {
	fmt.Println(input)
}

func sayHello(name string) string {
	return "Hello" + name
}

func main() {
	echo(sayHello("Selva"))
}
