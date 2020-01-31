package main

import "fmt"

func main() {

	parents := []string{"Carol","Mike"}
	kids := []string {"Marcia","Jan","Cindy","Greg","Peter","Bobby"}
	family := make([]string,len(parents)+len(kids))
	family = append(parents,kids...)
	fmt.Println(family)
	fmt.Println("len :",len(family))
	fmt.Println("cap :",cap(family))
	youngest := family[3:len(family)]  
	fmt.Println(youngest)
	extras := make([]string,1,2)
	extras[0] ="Alice"
	fmt.Println(extras)

}