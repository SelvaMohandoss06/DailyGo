package main
import (
	"fmt"
	
)



type Beatle struct {
	Name string
}

type foo struct{
	id int
	name string
}
 

func main() {
	b := &Beatle{Name: "Bingo"}
	fmt.Println(b.Name)
	changeBeatleName(b)
	fmt.Println(b.Name)
	changeBeatleNameWithoutPointer(*b)
	f := new(foo)
	f1 := &foo {}
	 i := new(int)
	 s := new(string)
	fmt.Printf("i: %v , *v :%d \n",i, *i)
	fmt.Printf("s : %v, *s :%q\n",s, *s)
	fmt.Printf("f: %v,*f: %+v\n",f,*f)
	fmt.Printf("f1 :%v,*f1:%+v\n",f1, *f1)

	ps := new(string)
	*ps = "hello"
	str := *ps
	fmt.Println(str)

	// var strOutput *string
    // *strOutput = "boom"



}

func changeBeatleName(b *Beatle) {
	b.Name = "George"
	fmt.Println(b.Name)
}

func changeBeatleNameWithoutPointer(b Beatle) {
	b.Name = "Test"
	fmt.Println(b.Name)
}
