package main

import (
	"fmt"
)

type bk interface {
	store()
}

type admin struct {
	name, role, dep string
}
type user struct {
	name, role string
}

func (ad admin) store() {
	fmt.Println("this is the admin  ", ad.dep, ad.name, ad.role)
}
func (ad user) store() {
	fmt.Println("this is the admin  ", ad.name, ad.role)
}

func gh(b bk) {
	b.store()
}

func main() {
	j := admin{name: "qhbdkw", role: "admin", dep: "hr"}
	k := admin{name: "qhbdkw", role: "admin"}
	gh(j)
	gh(k)

}
