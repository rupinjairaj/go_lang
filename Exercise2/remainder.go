package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter 2 numbers x and y to find remainder when y divides x")
	fmt.Scan(&x, &y)
	remainder := x % y
	fmt.Println(remainder)
}
