package main

import (
	"fmt"
	r "reflect" // using an alias
)

func main() {
	rune := 'd'
	fmt.Println(rune)
	fmt.Printf("%T\n", rune)
	fmt.Println(r.TypeOf(rune))
}
