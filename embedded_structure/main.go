package embedded_structure

import "fmt"

type Builder interface {
	Build()
}

type Person struct {
	Name string
	Age  int
}

type Building struct {
	Builder
	Name string
}

// ! wood builder
type WoodBuilder struct {
	Person
}

func (wb WoodBuilder) Build() {
	fmt.Println("build wood house")
}

// ! brick builder
type BrickBuilder struct {
	Person
}

func (bb BrickBuilder) Build() {
	fmt.Println("build brick house")
}

func EmbeddedStructureExample() {
	/*
		first of all we run Building struct
		that takes Builder and Name as arguments
		the Builder is an interface that has Build() method
		so we have to pass any struct that has Build() method

		we have 2 other structs that have Build() method
		they are WoodBuilder and BrickBuilder
		so we can pass them because they implement Builder interface
	*/
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "John",
			Age:  40,
		}},
		Name: "Wood house",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				Name: "Mark",
				Age:  30,
			},
		},
		Name: "brick house",
	}
	brickBuilding.Build()
}
