// empty interface
package main

import (
	"fmt"
)

func main() {
	m := map[string]interface{}{}
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true

	fmt.Println(m)
}
