package main

import "fmt"

type Matrix [3][3]int

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}
	a2 = a1
	fmt.Println(a2)

	m := Matrix{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
	}

	fmt.Println(m)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}
