package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
}

func printValue(s *string) {
	fmt.Println(&s)
	fmt.Println(*s)
	*s = "Selva"
	fmt.Println(*s)
}
