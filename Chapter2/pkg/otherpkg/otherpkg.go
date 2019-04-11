package otherpkg

import "../somepkg"

func OtherFunc() {
	somepkg.SomeFunc()
	somepkg.SomeVar = 5
}