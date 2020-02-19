package main

import (
	"fmt"
	"math"
)

type Entertainer interface {
	Play()
}



type Beatle struct{
	Name string
	Instrument string
}

func(b Beatle) Play() {
	fmt.Printf("%s plays %s",b.Name,b.Instrument)
}
func entertain(e Entertainer)  {
	e.Play()
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius *c.Radius
}

func(c Circle) String() string {
	return fmt.Sprintf("Circle{Radius:%.2f}",c.Radius)
}

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) Area() float64{
	return  r.Height *r.Width
}

func (r Rectangle) String() string {
	return  fmt.Sprintf("Rectangle{Height: %.2f,width:%.2f}",r.Height,r.Width)
}

type Sizer interface {
	Area() float64
}

type Shaper interface{
	Area() float64
	fmt.Stringer
}
func main() {
  b := Beatle {Name : "Ringo", Instrument: "Drums"}
  entertain(b)
  c := Circle {Radius:15}
  fmt.Println(c)
  fmt.Println(c.Area())
  r := Rectangle {Height:20,Width:30}
  fmt.Println(r)
  fmt.Println(r.Area())

}



