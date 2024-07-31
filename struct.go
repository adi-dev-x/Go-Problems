package main

import (
	"fmt"
)

type h struct {
	name string
	age  int
}

func (gh h) fg() {

	fmt.Println(gh.name, "    ", gh.age)
}
func main() {
	j := h{name: "ahjcbhjascbhjacsbhjacs", age: 23}
	j.fg()

}
