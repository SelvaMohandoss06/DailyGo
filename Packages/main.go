package main

import (
	"fmt"	
)

type User struct{
	Name string
}

func main() {
	u := User {Name :"Golang"}
	fmt.Println(u)
	
}