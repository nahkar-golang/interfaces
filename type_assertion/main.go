package typeAssertion

import "fmt"

type Runner interface {
	Run() string
}

type Human struct {
	Name string
}

func (h *Human) Run() string {
	return "Human is running"
}

type Duck struct {
	Name, Surname string
}

func (d *Duck) Run() string {
	return "Duck is running"
}

func run(r Runner) {
	if _, ok := r.(*Human); ok {
		fmt.Println("I am a human, and I am running")
	} else if _, ok := r.(*Duck); ok {
		fmt.Println("I am a duck, and I am running")
	}
	fmt.Println(r.Run())
}
func run2(r Runner) {
	switch v := r.(type) {
	case *Human:
		fmt.Println("I am a human, and I am running", v.Run())
	case *Duck:
		fmt.Println("I am a duck, and I am running", v.Run())
	}
	fmt.Println(r.Run())
}

func ExampleTypeAssertion() {
	var runner Runner
	fmt.Println(runner == nil)                            // true
	fmt.Printf("type %T and value %v \n", runner, runner) // type <nil> and value <nil>

	var unnamedRunner *Human
	fmt.Printf("type %T and value %v \n", unnamedRunner, unnamedRunner) // type *main.Human and value <nil>

	runner = unnamedRunner
	fmt.Println(runner == nil)                            // false
	fmt.Printf("type %T and value %v \n", runner, runner) // type *main.Human and value <nil>

	namedRunner := &Duck{Name: "Duck", Surname: "Duck"}

	runner = namedRunner
	fmt.Printf("type %T and value %v \n", runner, runner) // type *main.Duck and value <nil>

	human := &Human{Name: "Human"}
	duck := &Duck{Name: "Duck", Surname: "Duck"}

	run(human)
	run(duck)

	run2(human)
	run2(duck)
}
