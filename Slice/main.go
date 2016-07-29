package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4}
	myOtherSlice := []int{5, 6, 7, 8}
	fmt.Println(mySlice)
	fmt.Println(myOtherSlice)
	// NOTE THE SYNTAX FOR APPENDING; KEEP IN MIND THE '...'
	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
	// NOTE THE SYNTAX FOR APPENDING; KEEP IN MIND THE '...'
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
