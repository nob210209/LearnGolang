package somepkg

import "fmt"

var SomeVar int
var someVar2 int

func SomeFunc() {
	SomeVar = 10
	someVar2 = 5
	fmt.Printf("value: %d\n", SomeVar)
}