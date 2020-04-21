package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("server 监听8888 端口！")
	defer l.Close()

	for {
		fmt.Println("等待客户端连接！")
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("连接到客户端 %v \n", c.RemoteAddr().String())
		go process(c)
	}

	fmt.Println("server 退出！")
}

func process(c net.Conn) {
	defer c.Close()
	for {
		rb := make([]byte, 1024)
		n, err := c.Read(rb)
		if err != nil {
			fmt.Printf("server read failed: %v \n", err.Error())
			break
		}
		fmt.Printf("got client msg :%v \n", string(rb[:n]))
	}
	fmt.Println("server process exit!")
}
