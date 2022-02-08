package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("UDP服务端开始了 ...... ")
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("监听端口出错，err：", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("接收数据出错，err：", err)
			continue
		}
		fmt.Printf("接收数据内容：%v \n 数据来源地址：%v 。接收数据数量：%v \n ", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 数据原样返回客户端
		if err != nil {
			fmt.Println("数据原样返回客户端出错，err：", err)
			continue
		}
	}
}
