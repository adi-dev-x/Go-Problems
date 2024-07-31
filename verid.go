package main

import (
	"fmt"
	"time"
)

func f(d ...int) {
	g := 0
	for _, value := range d {
		g = g + value

	}
	for i := 0; i < 4; i++ {
		fmt.Println(g)
		time.Sleep(10 * time.Millisecond)

	}

}

func main() {

	go f(1, 2, 3, 4)
	f(34, 5, 4)

}
