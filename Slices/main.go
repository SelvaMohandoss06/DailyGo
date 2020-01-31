package main

import (
	"fmt"
	"strings"
)

type Modules [][]string

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	for i, n := range names {
		fmt.Printf("%d = %s\n", i, n)
	}

	fruits := []string{"Apple", "Orange", "Guava", "PineApple"}
	fruits = append(fruits, "Pear")
	moreFruits := []string{"Lemon", "Avocado"}
	fruits = append(fruits, moreFruits...)
	fmt.Println(fruits)
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	a := []string{}
	b := make([]string, 0)
	fmt.Println(a)
	fmt.Println(b)

	a1 := make([]int, 1, 3)
	fmt.Println(a1)

	a2 := make([]string, 2)
	a2 = append(a2, "foo", "bar")
	fmt.Println("%q", a2)

	course := Modules{
		[]string{"Chapter One : Syntax"},
		[]string{"Chapter Two : Arrays and Slices"},
	}

	fmt.Println(course)
	fmt.Println(fruits[1:3])
	fmt.Println(fruits[1:])
	fmt.Println(fruits[:2])
	students := []string{"John", "Paul", "George", "Rongo"}
	fmt.Println(students)
	firstGrades := students[:3]
	fmt.Println(firstGrades)
	for i, g := range firstGrades {
		firstGrades[i] = strings.ToUpper(g)
	}

	original := []string {"carrot","potato","cucumber","onion"}
	ref := original
	dup := make([]string, len(original))
	copy(dup,original)

	ref[0] ="zuchinni"
	original[1] ="tomato"

	fmt.Println("original : ",original)
	fmt.Println("ref : ",ref)
	fmt.Println("dup :",dup)

}
