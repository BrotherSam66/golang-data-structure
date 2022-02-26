package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"stickypag_server/encodedecode"
)

func process(conn net.Conn) {
	defer conn.Close() // 最终关闭连接
	reader := bufio.NewReader(conn)
	for {
		msg, err := encodedecode.Decode(reader) // 调用解码方法，去掉前面 4 字节的包长度字段
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode解码错误。err：", err)
			return
		}

		fmt.Println("收到客户端发来的数据为：", msg)

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
