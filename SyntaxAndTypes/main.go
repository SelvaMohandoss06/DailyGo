package main

import "fmt"

const(
	year = 365
	leapYear  = int32(366)
)

func main() {

	var i int = 1
	fmt.Println(i)
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n",a,a)
	fmt.Printf("var b %T = %q\n",b,b)
	fmt.Printf("var c %T = %+v\n",c,c)
	fmt.Printf("var d %T = %+v\n\n",d,d)

	names := []string {"Mary","John","Bob","Anna"}

	for  i, n := range names {
		fmt.Printf("index: %d = %q\n",i,n)
	}

	j,k,l := "gopher",2.05,15
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	var pi float64 = 3.14
	var week int = 7
	fmt.Println(pi)
	fmt.Println(week)
	// var maxUint32 uint32 = 4294967295 + 1
	// fmt.Println(maxUint32)
	var ok bool
	fmt.Println(ok)
	str := `Hello "Go"`
	fmt.Println(str)
	interPretedString := "Say \"hello\" to Go!"
	fmt.Println(interPretedString)
	utfstring := "Hello, 世界"
	fmt.Println(utfstring)

	for i,c := range utfstring{
		fmt.Printf("%d: %s\n",i,string(c))
	}

	fmt.Println("length of 'Hello, 世界': ",len(utfstring))
	const gopher ="Genny"
	fmt.Println(gopher)
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)

	const (
		Apple int = iota
		Banana
		Orange
	)

	fmt.Printf("The value of Apple is %v\n",Apple)
	fmt.Printf("The Value of Orange is %v\n",Orange)
	fmt.Printf("The Value of Banana is %v\n",Banana)
	fmt.Println("Print a statement  followed by a line return")
	fmt.Printf("Hello %s\n",gopher)
  type User struct{
	  name string
  }

  u := User {name :"GoLang"}
  fmt.Printf("name : %+v",u)





}