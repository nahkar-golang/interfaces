// Function can accept an interface as an argument

package main

import (
	"fmt"
)

type goodMorning struct{}

func (gm *goodMorning) greet(name string) string {
	return fmt.Sprintf("Good Morning %s", name)
}

type goodEvening struct{}

func (ge *goodEvening) greet(name string) string {
	return fmt.Sprintf("Good Evening %s", name)
}

type greeter interface {
	greet(string) string
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

func main() {
	var gm = goodMorning{}
	var ge = goodEvening{}

	fmt.Println(gm.greet("Carl"))
	fmt.Println(ge.greet("Carl"))

	var g = []greeter{&gm, &ge}
	for _, greeter := range g {
		fmt.Println(greeter.greet("John"))
	}

	sayHello(&gm, "Mike")
	sayHello(&ge, "Mike")

	sayHello(&goodMorning{}, "Alice")
	sayHello(&goodEvening{}, "Alice")
}
