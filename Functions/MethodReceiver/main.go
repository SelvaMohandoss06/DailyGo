package main

import (
	"fmt"
	"strings"
)

type Beatle struct {
	Name       string
	Instrument string
}
type Student struct {
	Name  string
	Grade string
}

func String(b Beatle) string {
	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)

}

func (b Beatle) String() string {
	return fmt.Sprintf("%s  plays %s", b.Name, b.Instrument)
}

func (s Student) String() string {
	return fmt.Sprintf("%s studies %s", s.Name, s.Grade)
}

type User struct {
	First string
	Last string
}

func(u *User) Titleize(){
	u.First = strings.Title(u.First)
	u.Last = strings.Title(u.Last)
}

func(u User) Reset()  {
	u.First =""
	u.Last =""
}

func main() {
	b := Beatle{Name: "Ringo", Instrument: "Guitar"}
	fmt.Println(String(b))
	fmt.Println(b.String())

	s := Student{Name: "Janesh", Grade: "Fifth"}
	fmt.Println(s.String())
	
	u := User {First :"Selva", Last:"Mohandoss"}
	u.Titleize()
	fmt.Println(u)

	u.Reset()
	fmt.Println(u)
}
