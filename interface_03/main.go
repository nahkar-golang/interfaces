// Composite interface implementation (SOLID interface segregation)
package main

import (
	"fmt"
)

type animal interface {
	walker
	runner
}

type walker interface {
	walk()
}
type runner interface {
	run()
}

type cat struct{}

func (c *cat) walk() {
	fmt.Println("Cat is walking")
}
func (c *cat) run() {
	fmt.Println("Cat is running")
}

// ! Bird
type bird interface {
	walker
	flier
}

type flier interface {
	fly()
}

type eagle struct{}

func (e *eagle) walk() {
	fmt.Println("Bird is walking")
}
func (e *eagle) fly() {
	fmt.Println("Bird is flying")
}

func walk(w walker) {
	w.walk()
}

func main() {
	var c animal = &cat{}
	var e bird = &eagle{}

	walk(c)
	walk(e)
}
