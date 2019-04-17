package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func main() {
	v := Vector{X: 2, Y: 5}
	fmt.Println(v)
}