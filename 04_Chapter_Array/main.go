package main

import "fmt"

type Matrix [3][3]int

func main() {
	names := [4]string{}
	names[0] = "Selva"
	names[1] = "Mohandoss"
	names[2] = "John"
	names[3] = "Paul"
	fmt.Println(names)
	employees := [4]string{"selva", "mohandoss", "janesh"}
	fmt.Println(employees)
	ints := [2]int{}
	fmt.Println(ints)
	a1 := [2]string{"data1", "data2"}
	a2 := [2]string{}
	a1[0] = "test"
	a2 = a1

	fmt.Println(a2)
	a3 := [3]string{}
	// a3 = a2
	fmt.Println(a3)

	m := Matrix{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(m)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fruits := [4]string{"Apple", "Orange", "StrawBerry", "Grapes"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for _, f := range fruits {
		fmt.Println(f)
	}

	for i, f := range fruits {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, f)
	}

}
