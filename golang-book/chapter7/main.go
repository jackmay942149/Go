package main

import "fmt"

type Shape interface {
	perimeter() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	len float64
	wid float64
}

const PI float64 = 3.1419

func (c *Circle) perimeter() float64 {
	return 2 * PI * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.len + r.wid)
}

func TotalPerimeter(shapes ...Shape) (total float64) {
	for _, v := range shapes {
		total += v.perimeter()
	}
	return
}

func main() {
	c := &Circle{0, 0, 1}
	r := &Rectangle{2, 2}
	fmt.Println(TotalPerimeter(c, r))
}
