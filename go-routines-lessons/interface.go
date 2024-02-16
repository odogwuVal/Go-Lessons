package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length * r.Width)
}

func main() {
	rectangle := &Rectangle{
		length: 12.0,
		Width:  4.0,
	}

	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())
}
