package main

import "fmt"

func k(l map[int]int) []int {
	j := make([]int, len(l))

	for i, value := range l {
		j[i] = value

	}

	return j
}

func main() {
	var x = [5]int{3, 4, 4, 55, 4}

	var j = make(map[int]int)

	for i := 0; i < len(x); i++ {
		j[i] = 2 * i
	}

	fmt.Println(j)

	fmt.Println("this is newwwwwwww   ", k(j))
}
