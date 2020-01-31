package main
import "fmt"
func main() {
	fruits := [4] string {"Banana","Orange","Pineapple","Strawberry"}
	for i:= 0; i < len(fruits); i++ {
     fmt.Println(fruits[i])
	}

	for i,n := range fruits{
		fmt.Printf("%d - %s\n",i,n)
	}
}
