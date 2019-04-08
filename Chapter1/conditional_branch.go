package main

import (
	"fmt"
	"time"
)

func main() {

	//セミコロンで区切ることで条件判定の前処理を書くことが出来ます。
	//変数のスコープを限定したい時などに便利です。
	if hour := time.Now().Hour(); hour >= 6 && hour < 12 {
		fmt.Println("朝です。")
	} else if hour < 19 {
		fmt.Println("昼です。")
	} else {
		fmt.Println("夜です。")
	}

	//Go言語ではswitchの各caseの最後に何も書かれていない場合、
	// 自動的にbreakが補完されます。
	dayOfWeek := "月"
	switch dayOfWeek {
	case "土":
		//意図的に次のケースへ処理を続けたい場合は明示的にフォールスルーをします。
		fallthrough
	case "日":
		//Go言語ではswitchの各caseの最後に何も書かれていない場合、
		//自動的にbreakが補完されます。
		fmt.Println("休みです。")
	default:
		fmt.Println("仕事です・・・。")
	}

	//カンマを使って複数の条件をまとめることもできます。
	//上記の例ならば、こちらの方がすっきりとコードを書くことができます。
	dayOfWeek2 := "月"
	switch dayOfWeek2 {
	case "土", "日":
		fmt.Println("休みです。")
	default:
		fmt.Println("仕事です・・・。")
	}

}