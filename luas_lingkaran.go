package main

import "fmt"
import "math"

func Calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)
	circumference := math.Pi * d

	return area, circumference
}

func main() {
	a, c := Calculate(14)
	fmt.Printf("The area is : %v", a)
	fmt.Println()
	fmt.Printf("The circumference is : %v", c)
}