package main

import (
	"chapter8/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
	fmt.Println(math.Max(xs))
	fmt.Println(math.Min(xs))
}
