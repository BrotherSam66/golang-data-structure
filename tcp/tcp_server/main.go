package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 最终关闭连接
	reader := bufio.NewReader(conn)
	var buff [128]byte
	for {
		n, err := reader.Read(buff[:]) // 读取数据，读到的字节数为n
		if err != nil {
			fmt.Println("读取错误。err：", err)
			break
		}
		recStr := string(buff[:n])
		fmt.Println("收到客户端发来的数据为：", recStr)
		conn.Write([]byte("服务端收到的数据是：" + recStr)) // 收到的数据返回给客户端

	}

}
func main() {
	fmt.Println("TCP服务端开始了 ...... ")
	listen, err := net.Listen("tcp", "127.0.0.1:10000") // 以TCP协议监听 ip+端口号
	if err != nil {
		fmt.Println("监听错误。err：", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("建立连接错误。err：", err)
			continue
		}
		go process(conn) // 调用一个 goroutine 去处理数据

	}

}
