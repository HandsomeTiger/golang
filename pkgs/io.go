package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

//字符串读取
func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("han"), 12)
	fmt.Println(data)
}

//读取输入
func sampleReadStdin() {
	fmt.Println("please input from stdin:")

	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

//文件读入
func samleReadFile() {
	file, _ := os.Open("io.go")
	defer file.Close()

	data, _ := ReadFrom(file, 9)
	fmt.Println(string(data))
}

//io缓冲区（读取器/写入器）
func sampleReadBufio() {
	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)
	//Peek只读不取
	data, _ := bufReader.Peek(6)
	fmt.Println(data, bufReader.Buffered())
	//读取到空格为止，读取
	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())
	//stdOut输出设备，默认是屏幕 类似于文件fopen
	w := bufio.NewWriter(os.Stdout)
	//往w写入，类似于fwrite
	fmt.Fprint(w, "Hello ")
	fmt.Fprint(w, "world!")
	w.Flush()
}

//计算文件行数
func fileLineCalcer() {
	//os.Args获取到命令行启动时从命令行传入的参数
	if len(os.Args) < 2 {
		return
	}
	//参数从下标0开始，这里获取第一个参数 也就是传入的文件名
	filename := os.Args[1]
	fmt.Println(filename)
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//读取器
	reader := bufio.NewReader(file)
	var line int

	for {
		//读取文件的每一行，isPrefix返回是否是超宽行
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {

			line++
		}
		fmt.Println(string(data))
	}
	fmt.Println(line)
}

//读取bmp文件头
func readBmpHeader() {
	file, err := os.Open("io_test.bmp")
	if err != nil {
		return
	}
	defer file.Close()
	//二进制读取binary包
	var headA, headB byte
	//参数1是文件，参数二是二进制文件格式，windows和linux是LitteleEndian,第三个参数是接收的变量
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	var size uint32
	binary.Read(file, binary.LittleEndian, &size)
	//....
	fmt.Printf("%c%c\n%v\n", headA, headB, size)
	infoHeader := new(BitmapInfoHeader)
	binary.Read(file, binary.LittleEndian, infoHeader)
	fmt.Println(infoHeader)
}

type BitmapInfoHeader struct {
	Size           uint32
	Width          int32
	Heit           int32
	Places         uint16
	Bitcount       uint16
	Compression    uint32
	SizeImage      uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed       uint32
	ClrImpoerant   uint32
}

func main() {
	// sampleReadFromString()
	// sampleReadStdin()
	// samleReadFile()
	// sampleReadBufio()
	// fileLineCalcer()
	readBmpHeader()

}
