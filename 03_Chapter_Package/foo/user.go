package foo

type User struct {
	First    string
	Last     string
	password string
}

func NewUser(first, last, password string) User {
	return User{
		First:    first,
		Last:     last,
		password: password,
	}
}
