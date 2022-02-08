package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000") // 拨号呼叫服务端
	if err != nil {
		fmt.Println("拨号呼叫错误。err：", err)
		return
	}
	defer conn.Close() // 最终关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n') // 读键盘输入一行字符串
		if err != nil {
			fmt.Println("读键盘输入错误。err：", err)
			return
		}
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 输入 Q 就退出
			return
		}
		n, err := conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			fmt.Println("发送数据错误。err：", err)
			return
		}
		fmt.Println("发送数据数量为：", n)
		//reader := bufio.NewReader(conn)
		//var buff [128]byte
		//_, err = reader.Read(buff[:]) // 读取数据，读到的字节数为n
		//if err != nil {
		//	fmt.Println("读取错误。err：", err)
		//	break
		//}
		//recStr := string(buff[:n])
		//fmt.Println("收到服务端返回来的数据为：", recStr)
	}
}
