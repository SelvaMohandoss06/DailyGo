package main

import (
	"fmt"
	"sync"
)

//var wg sync.WaitGroup

//func sayHello() {
//	fmt.Println("Hello")
//	wg.Done()
//}

func main() {
	//wg.Add(1)
	//go sayHello()
	//wg.Wait()
	var wg sync.WaitGroup
	names := [] string {"Fred","Wilma","Barney","Betty"}
	wg.Add(len(names))
	for i, n := range  names {
       go func() {
         fmt.Printf("%q is %d at index position \n",n,i)
         wg.Done()
	   }()
	}
   wg.Wait()
}

