package main

import "fmt"

func main() {
	input := "hello"
	charMap := make(map[string]int)

	for _, char := range input {
		p := string(char)
		charMap[p]++

	}

	fmt.Println("String:", input)
	fmt.Println(charMap)

}
