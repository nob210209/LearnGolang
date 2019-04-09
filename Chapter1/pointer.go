package main

import (
	"fmt"
)

func main() {
	// int型のポインタ変数
	var pointer *int
	// int型変数
	var n int = 100

	// &（アドレス演算子）を使って、nのアドレスを代入
	pointer = &n

	fmt.Println("nのアドレス：", &n)
	fmt.Println("pointerの値：", pointer)

	fmt.Println("nの値：", n)
	// *(間接参照演算子）を利用して、ポインタの中身を取得
	fmt.Println("pointerの中身：", *pointer)

	a, b := 10, 10

	// aはそのまま、bはアドレス演算子をつけて呼び出す
	called(a, &b)

	fmt.Println("値渡し：", a)
	fmt.Println("ポインタ渡し：", b)
}

func called(a int, b *int) {
	// 変数をそのまま変更
	a = a + 1
	// 変数の中身を変更
	*b = *b + 1
}