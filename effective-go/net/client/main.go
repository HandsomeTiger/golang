package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client 开始连接server 8888端口！")
	c, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client conn server failed :%v", err.Error())
		return
	}
	defer c.Close()
	read := bufio.NewReader(os.Stdin)
	for {
		sr, err := read.ReadString('\n')
		if err != nil {
			fmt.Printf("read stdin faied %v", err.Error())
			return
		}
		sr = strings.Trim(sr, " \n\r")
		if sr == "exit" {
			break
		}
		n, err := c.Write([]byte(sr + "\n"))
		if err != nil {
			fmt.Printf("write failed %v", err.Error())
			return
		}
		fmt.Printf("wirte success %d \n", n)
	}

	fmt.Println("client 关闭连接！")
}
