package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Enter the distance in meters")
	fmt.Scanln(&x)
	inches := x * 39.3701
	fmt.Println("The distance in inches is", inches, "inches")
}
