// we have to pass pointer to the interface
package main

import (
	"fmt"
)

type Cat struct {
	name string
}

func (c *Cat) updateName() {
	c.name = "cat updated"
}

type Dog struct {
	name string
}

func (d *Dog) updateName() {
	d.name = "dog updated"
}

type Animal interface {
	updateName()
}

func main() {
	c := Cat{name: "cat"}

	d := Dog{name: "dog"}

	var animals = []Animal{&c, &d}

	for _, animal := range animals {
		animal.updateName()
	}

	fmt.Println(c.name)
	fmt.Println(d.name)

}
