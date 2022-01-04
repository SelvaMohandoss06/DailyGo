package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func main() {
	u := User{
		Name:     "Selva",
		Email:    "Selvamohandoss@gmail.com",
		Password: "sss",
	}
	fmt.Println(u)
	userData := User{Name: "Janesh", Email: "selva@gmail.com", Password: "eee"}
	fmt.Printf("user:%v\n", userData)
	err := json.NewEncoder(os.Stdout).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}

}
