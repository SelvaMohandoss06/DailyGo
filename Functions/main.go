package main

import "fmt"

func main() {
	sayHello()
	sayGreeting("Hello")
	sayHello2("Hello", "Selva")
	one()
	l, c := slicer([]string{"John", "Paul", "George", "Ringo"})
	fmt.Println(l)
	fmt.Println(c)
	fmt.Println(MeaningOfLife())
}

func sayHello() {
	fmt.Println("Hello")
}

func sayGreeting(greeting string) {
	fmt.Println(greeting)
}

func sayHello2(greeting string, name string) {
	fmt.Printf("%s,%s", greeting, name)
}

func one() string {
	return "1"
}

func slicer(s []string) (int, int) {
	return len(s), cap(s)
}

func exists() (exist bool) {
	exist = true
	return
}

func someFunction(val int) (ok bool, err error) {
	if val == 0 {
		return false, nil
	}
	ok = true
	return
}

func MeaningOfLife() (rc int) {
	defer func() { rc++ }()
	return 41
}
