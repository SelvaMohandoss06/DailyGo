package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	users := make(map[int]User)
	u := User{ID: 1, Name: "Selva Mohandosss"}
	users[u.ID] = u
	u = User {ID :2,Name :"Janesh Senthilnathan"}
	users[u.ID] = u
	fmt.Println(users)

}
