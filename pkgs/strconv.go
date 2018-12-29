package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

func main() {
	// strconv()
	serialize()

}

//字符串转换
func strconvTest() {
	//整型转字符串
	fmt.Println(strconv.Itoa(10))
	//字符串转整型
	fmt.Println(strconv.Atoi("711"))
	//解析成bool类型
	fmt.Println(strconv.ParseBool("false"))
	//解析浮点数类型
	fmt.Println(strconv.ParseFloat("3.14", 32))
	//格式化(与解析是互逆的)
	fmt.Println(strconv.FormatBool(false))
	//十进制转换成16进制
	fmt.Println(strconv.FormatInt(20, 16))

}

//结构体的序列化和反序列化
type Person struct {
	Name string `xml:"name,attr"`
	Age  int
}

func serialize() {

	p := Person{Name: "张三", Age: 18}

	var data []byte
	var err error
	//把结构体转换成xml
	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	p2 := new(Person)

	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p2)
}
