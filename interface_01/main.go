package main

import (
	"fmt"
)

type cat struct {
	name string
}

func (c *cat) updateName() {
	c.name = "cat updated"
}

type dog struct {
	name string
}

func (d *dog) updateName() {
	d.name = "dog updated"
}

type animal interface {
	updateName()
}

func main() {
	c := cat{name: "cat"}

	d := dog{name: "dog"}

	var animals = []animal{&c, &d}

	for _, animal := range animals {
		animal.updateName()
	}

	fmt.Println(c.name)
	fmt.Println(d.name)

}
