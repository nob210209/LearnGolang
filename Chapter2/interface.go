package main

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {}
func (d *Dog) Bark() {
	fmt.Println("わんわん")
}

type Cat struct {}
func (c *Cat) Bark() {
	fmt.Println("にゃーにゃー");
}

func MakeAnimalBark(a Animal) {
	fmt.Println("鳴け！");
	a.Bark();
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	MakeAnimalBark(dog)
	MakeAnimalBark(cat)
}