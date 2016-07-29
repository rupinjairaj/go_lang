package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3
	mySlice[3] = 4
	mySlice[4] = 5
	fmt.Println(mySlice)
}
