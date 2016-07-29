package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println("Enter a float number")
	fmt.Scan(&x)
	fmt.Println(math.Ceil(x))
}
