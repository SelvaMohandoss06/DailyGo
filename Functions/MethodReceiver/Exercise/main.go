package main

import "fmt"

type User struct {
	First string
	Last  string
	Email string
}

func main() {
	u := User{First: "Selva", Last: "Mohandoss", Email: "selvamohandoss#gmail.com"}
	fmt.Println(u.GetName())
	fmt.Println(u.FormattedEmail())
	u.Reset()
	fmt.Println(u)

}

func (u User) GetName() string {
	return u.First + u.Last + u.Email
}

func (u *User) FormattedEmail() string {
	u.Email = fmt.Sprintf("%s would return %s", u.First, u.Last)
	return u.Email
}

func (u *User) Reset() {
	u.First = ""
	u.Last = ""
	u.Email = ""
}
