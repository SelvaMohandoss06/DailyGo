package main

import (
	"errors"
	"fmt"
)

type Invite struct{
	ID int
	Email
}

type Email struct {
	DisplayName string
	Username string
	Domain string
}


func main() {
	invite := Invite {
		ID :1,
		Email : Email {
			DisplayName:"Selva Mohandoss",
			Username: "smohandoss",
			Domain: "gmail.com",
		},
	}
	if err := SendEmail(invite); err != nil {
		fmt.Printf("error sending email :%s",err)
		return
	}
	fmt.Printf("sent email to %s",invite.Address())
}

func (e Email) Address() string{
	return fmt.Sprintf("%s <%s@%s>",e.DisplayName,e.Username,e.Domain)
}

func SendEmail(i Invite) error {
	address := i.Address()

	if address == "" {
		return  errors.New("no email address provided")
	}
	return  nil
}

