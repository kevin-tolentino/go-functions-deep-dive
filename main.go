package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	answer, _ := divide(6, 3) //you should handle the error appropriately
	fmt.Printf("%f\n", answer)

	total := sum(2.1,3,4,5,6)
	fmt.Println("!!total", total)

	numbers := []float64{2.1,3,4,5,6}
	summedSlice := sum(numbers...) //spreading the slice values as arguments for the sum
	fmt.Printf("total of number %f", summedSlice)

}
//func to show variadic functions (spread syntax)
func sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}

	return p1 / p2, nil
}
