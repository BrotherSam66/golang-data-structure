package main

import (
	"arrayqueue/arrayqueueutils"
	"arrayqueue/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 循环数组实现队列，排队叫号。
func main() {
	//arr := make([]int, 20)                   // 队列本身(是一个切片)
	//var front int = 0                        // 队列头位置
	//var rear int = 0                         // 队列尾巴位置
	//var nowNum int = 1                       // 将要发号的号码

	arr := models.NewArrayQueue(20, 0, 0) // 队列本身(是一个结构体的切片)，开了20，只能容纳19人

	inputReader := bufio.NewReader(os.Stdin) // 读取键盘或文件的读取器
	fmt.Println("请输入：")
	for {
		input, err := inputReader.ReadString('\n') // 读取键盘，直到回车
		if err != nil {
			fmt.Println("键盘输入错误，err：", err)
			return
		}
		inputStr := strings.Trim(input, "\r\n")
		switch strings.ToUpper(inputStr) {
		case "Q":
			fmt.Println("打【Q】退出")
			return
		case "A":
			errStr := arrayqueueutils.AddOne(arr)
			if errStr != "" { // 说明有错误
				fmt.Println("有错误，错误信息：", errStr)
				continue
			}
			fmt.Println("打【A】追加了一个排队人，编号：", (*arr).Arr[(*arr).Rear]) // 队尾的那个值就是新增加排队人的序号
		case "G":
			errStr := arrayqueueutils.GetOne(arr)
			if errStr != "" { // 说明有错误
				fmt.Println("有错误，错误信息：", errStr)
				continue
			}
			fmt.Println("打【G】叫号了一个排队人，编号：", (*arr).Arr[(*arr).Front]) // 队头的那个值就是
		case "L":
			errStr := arrayqueueutils.ListALl(arr)
			if errStr != "" { // 说明有错误
				fmt.Println("有错误，错误信息：", errStr)
				continue
			}
		default:
			fmt.Println("键盘输入错误，重新输入：")
			continue
		}

	}

}
