package main

import "fmt"

func main() {
	beatles := map[string]string{}
	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"
	fmt.Println(beatles)
	emails := map[string]string{
		"selva@gmail.com":  "selva",
		"janesh@gmail.com": "janesh",
		"tim@gmail.com":    "tim",
	}
	fmt.Println(emails)
	fmt.Println(len(beatles))
	fmt.Println(beatles["Paul"])

	for key, value := range beatles {
		fmt.Printf("%s plays %s\n", key, value)
	}
}
