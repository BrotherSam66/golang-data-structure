package main

import (
	"bufio"
	"fmt"
	"gotest/stringutils"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		inputStr, _ := inputReader.ReadString('\n') // 键盘读入一行
		inputStr = strings.Trim(inputStr, "\r\n")
		if strings.ToUpper(inputStr) == "Q" {
			return
		}
		fmt.Println("发送字符串长度为：", len(inputStr))
		fmt.Println("发送字符串内容为：", inputStr)
		acceptStr := stringutils.Reverse(inputStr)
		fmt.Println("接收字符串长度为：", len(acceptStr))
		fmt.Println("接收字符串内容为：", acceptStr)
	}
}
