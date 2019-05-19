package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hoge.txt")
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}
	defer file.Close()   // <- ここでCloseを遅延実行する

	//ファイル操作処理省略
	//deferが行われた以降の行のどこでreturnしても（またはpanicが発生しても）
	//必ずfile.Closeが実行されます。 このように、deferはよくリソースの解放に使用されます。

}
