package main

import "fmt"

const (
	year     = 365
	leapYear = int32(366)
)
const (
	Apple int = iota
	Banana
	Orange
)
const (
	read = 1 << iota
	write
	remove
	admin = read | write | remove
)

func main() {
	fmt.Println("Constant Demo")
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * leapYear)
	fmt.Printf("The Value of Apple is %v\n", Apple)
	fmt.Printf("The Value of Banana is %v\n", Banana)
	fmt.Printf("The Value of Orange is %v\n", Orange)
	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)
	fmt.Printf("admin =  %v\n", admin)

}
