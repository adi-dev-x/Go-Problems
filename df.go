package main

import (
	"fmt"
)

func h(s []int, g chan<- int) {
	for _, val := range s {
		g <- val
	}
	close(g)
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)

	var slice []int
	for i := 0; i < 101; i++ {
		if i%3 == 0 {
			slice = append(slice, i)
		}
	}
	go h(slice, ch)
	func() {
		for val := range ch {
			fmt.Println("vall is   ---", val)
		}

	}()
	// for val := range ch {
	// 	fmt.Println("vall is   ---", val)
	// }

}
