package main

import (
	"fmt"
)

type s struct {
	name string
}

func (sa s) f() {
	fmt.Println(sa.name)
}
func main() {

	c := s{name: "jdfjc"}

	c.f()
}
