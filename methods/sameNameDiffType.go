package main

import (  
    "fmt"
    "math"
)

type Rectangle struct {
    length int
    width int
}

type Circle struct {
    radius float64
}

func (r Rectangle) Area() int {
    return r.length * r.width
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle {
        length: 33,
        width: 20,
    }

    fmt.Printf("Area of rectangle %d \n", r.Area())

    c := Circle {
        radius: 10,
    }

    fmt.Printf("Radius of circle %f", c.Area())
}