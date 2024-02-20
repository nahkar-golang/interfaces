package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiply() int
	Divide() float64
	Subtract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Divide() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n Numbers) Subtract() int {
	return n.num1 - n.num2
}

func main() {
	var numbersMethods NumbersInterface
	numbers := Numbers{5, 2}

	numbersMethods = numbers
	fmt.Println(numbersMethods.Sum())
	fmt.Println(numbersMethods.Multiply())
	fmt.Println(numbersMethods.Divide())
	fmt.Println(numbersMethods.Subtract())
}
