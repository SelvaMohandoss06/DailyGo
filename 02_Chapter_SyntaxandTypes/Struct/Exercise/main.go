package main

import "fmt"

type Movie struct {
	Title    string
	Released bool
	Length   int
}

func main() {
	mov := Movie{
		Title:    "SpiderMan",
		Released: true,
		Length:   125,
	}
	fmt.Printf("movie %v", mov)

}
