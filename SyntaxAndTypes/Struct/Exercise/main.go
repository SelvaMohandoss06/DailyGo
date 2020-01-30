package main

import "fmt"

type Movie struct{
	Title string
	Released bool
	Length int
}

func main() {

	movie := Movie{Title :"Wizard of Oz",Released:true, Length :125}
	fmt.Println(movie)


}