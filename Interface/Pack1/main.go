package main

import (
	"fmt"
)

// Human ...
type Human struct {
	name  string
	age   int
	phone string
}

// Student ...
type Student struct {
	Human
	school string
	loan   float64
}

// Employee ...
type Employee struct {
	Human
	company string
	money   float64
}

// SayHi implemented by Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing implemented by Human
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la, ....", lyrics)
}

// Guzzle implemented by Human
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// SayHi overloaded by Employee
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work in %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// BorrowMoney implemented by Student
func (s *Student) BorrowMoney(amount float64) {
	s.loan += amount
}

// SpendSalary implemented by Employee
func (e *Employee) SpendSalary(amount float64) {
	e.money -= amount
}

// Men of type interface
type Men interface {
	SayHi()
	Sing(lyric string)
	Guzzle(beerStein string)
}

// YoungChap of type interface
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float64)
}

// ElderlyGent of type interface
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float64)
}
