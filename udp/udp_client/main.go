package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("UDP拨号出错，err：", err)
		return
	}
	defer socket.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n') // 读键盘输入一行字符串
		if err != nil {
			fmt.Println("读键盘输入错误。err：", err)
			continue
		}
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 输入 Q 就退出
			return
		}
		sendData := []byte(inputInfo)
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Println("发送数据失败。err：", err)
			continue
		}
		data := make([]byte, 4096)
		n, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败。err：", err)
			continue
		}
		fmt.Printf("收回来数据内容：%v \n 数据来源地址：%v 。收回来数据数量：%v \n ", string(data[:n]), addr, n)
	}
}
