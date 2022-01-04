package main

import "fmt"

func main() {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("var a  %T = %d\n", a, a)
	fmt.Printf("var b %T= %q\n", b, b)
	fmt.Printf("var c %T = %f\n", c, c)
	fmt.Printf("var d %T = %t\n\n", d, d)
	j, k, l := "gopher", 3.22, 4
	fmt.Printf("var j %q\n", j)
	fmt.Printf("var k %+v\n", k)
	fmt.Printf("var l %+v\n", l)
	var week int = 7
	fmt.Printf("var week %+v", week)
	Cup := 8
	Pint := Cup * 2
	Quart := Pint * 2
	Gallon := Quart * 2
	fmt.Println("1 Cup 2 Quart is ", Cup+Quart*2, "Ounces")
	fmt.Println("1 Cup and 2 Gallon is ", Cup+Gallon*2, "ounces")
	var maxUint8 uint8 = 255
	fmt.Println(maxUint8 + 1)
	// var maxUint8var uint8 = 255 + 1
	// fmt.Println(maxUint8var)
	var umaxUint8 uint8 = 11
	umaxUint8 = maxUint8 * 25
	fmt.Println("new value:", umaxUint8)
	rawString := `Say "hello" to Go`
	fmt.Println(rawString)
	interPretedString := "Say \"hello\" to Go"
	fmt.Println(interPretedString)
	unicodeStr := "Hello, 世界"
	fmt.Println(unicodeStr)
	character := 'A'
	fmt.Println(character)

}
