package main

import "fmt"

type Score int

func main() {
	var myScore Score = 100
	//型変換の例
	showInt(int(myScore))
}

func showInt(i int) {
	fmt.Printf("value: %d\n", i)
}