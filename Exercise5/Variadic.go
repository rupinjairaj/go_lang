package main

import "fmt"

func averageAge(age ...int) (average int) {
	total := 0
	for _, value := range age {
		total += value
	}
	average = total / len(age)
	return
}

func main() {
	a := []int{10, 10, 10}
	fmt.Println(averageAge(a...))
}
