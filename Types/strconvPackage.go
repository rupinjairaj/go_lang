package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("-42")
	if err != nil {
		return
	}
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
	s := strconv.Itoa(-42)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
