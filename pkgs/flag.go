package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//获取命令行工具
	fmt.Println(os.Args)
	// style2()
}

// flag 1
func style() {
	//可选参数 -method,返回值是一个字符串的指针
	//使用flag.exe -method a -value 1
	methodPtr := flag.String("method", "default", "method of sample")
	valuePtr := flag.Int("value", -1, "value of sample")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr)
}

// flag 2
func style2() {
	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")

	flag.Parse()
	fmt.Println(method, value)
}

//获取xml节点 ...
func ReadFile() {
	content, err := ioutil.ReadFile("flag_test.xml")
	if err == nil {
		return
	}

	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Println(name)
		case xml.EndElement:
		}
	}
}
