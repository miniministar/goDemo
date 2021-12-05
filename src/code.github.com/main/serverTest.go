package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
	}
	fmt.Println("listen ", "127.0.0.1:2000")
	for true {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		go process(conn) //启动一个goroutine处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from client failed, err:%v\n", err)
			break
		}
		recStr := string(buf[:n])
		fmt.Printf("收到客户端消息：%v\n", recStr)
		conn.Write([]byte(recStr)) //发送数据
	}
}
