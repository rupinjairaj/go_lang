package main

import "fmt"

func main() {
	i := 0
	total := 0
	for i < 1000 {
		if i%3 == 0 {
			total += i
			// fmt.Println(total)
		} else if i%5 == 0 {
			total += i
			// fmt.Println(total)
		}
		i++
	}
	fmt.Println("The sum is ", total)
}
