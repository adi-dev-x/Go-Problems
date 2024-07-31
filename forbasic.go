package main

import "fmt"

func h(f []int) {
	i := 0
	fmt.Println("this is the len", len(f))
	for i < len(f)-1 {
		var h int
		fmt.Scanf("%d", &h)
		f[i] = h
		i++

	}
	k := j(f)
	fmt.Println("added arr", sum(k, f))

}
func sum(h []int, k []int) []int {
	fmt.Println("in sum")
	u := 0
	g := make([]int, 0, len(h))
	//  g:=make([] int,len(h))
	for u < len(h) {
		fmt.Println("in sum", h[u], "   ", k[u])
		g = append(g, h[u]+k[u])
		u++

	}
	return g
}
func j(h []int) []int {
	i := 0
	j := make([]int, 5)
	for i < len(h)-1 {
		j[i] = h[i] * h[i+1]
		i++
	}
	return j
}

func main() {
	p := make([]int, 6, 10)
	h(p)

}
