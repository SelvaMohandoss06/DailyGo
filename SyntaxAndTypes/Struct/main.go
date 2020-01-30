package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	u := User{}
	fmt.Println(u)
	user := User{Name: "Selva", Email: "Selva@gmail.com"}
	fmt.Println(user)

}
