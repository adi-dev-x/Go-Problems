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
	//ch1 := make(chan int)
	var slice []int
	for i := 0; i < 101; i++ {
		if i%3 == 0 {
			slice = append(slice, i)
		}
	}
	go h(slice, ch)
	go func() {
		for val := range ch {
			fmt.Println("Received value from channel:", val)
			//ch1 <- val * val
		}
		//close(ch1)
	}()

	for valu := range ch1 {
		fmt.Println("Squared value received from channel:", valu)
	}
}
