package main

import "fmt"

func main() {
	fmt.Println("Enter a number (F/feet): ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9
	fmt.Println(output, " Celsius")
	output = input * 0.3048
	fmt.Println(output, " meteres")
}
