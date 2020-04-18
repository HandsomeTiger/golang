package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// ## defer推迟执行
	// 在打开资源后关闭的时候通常会用到defer，例如关闭打开的文件它能确保你不会忘记关闭文件。
	readFile()
	// 被推迟执行的函数，如果有实参传入该函数，那么该实参会在**推迟执行**的时候就会计算，而不是在**实际执行**的时候计算，例如：
	calcA()
	// defer的执行顺序是LIFO 先入后出
	calcB()
	//
	b()
	// defer 与return的执行顺序，输出结果为0 1，可以看出先执行return后面的方法，再执行的defer后面的方法
	// defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。
	deferReturn(0)
	// defer 推迟执行的方法，实参是指针，改变参数的值会影响defer后面函数的执行，因为defer的时候实参实际上是一个指针地址
	a := &ask{a: 1}
	deferReturn2(a)

	//E. defer的作用域
	//
	//1. defer只对当前协程有效（main可以看作是主协程）；
	//
	//2. 当panic发生时依然会执行当前（主）协程中已声明的defer，但如果所有defer都未调用recover()进行异常恢复，则会在执行完所有defer后引发整个进程崩溃；
	//
	//3. 主动调用os.Exit(int)退出进程时，已声明的defer将不再被执行。
}

// readFile defer常用来关闭打开的文件
func readFile() error {
	f, err := os.Open("README.md")
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	fmt.Printf("%s", content)

	return nil
}

// calcA 被推迟函数的实参（如果该函数为方法则还包括接收者）在推迟执行时就会求值， 而不是在调用执行时才求值。
func calcA() {
	a, b := 1, 2
	defer fmt.Println(sum(a, b))
	defer fmt.Println(a)

	a, b = 3, 4
	fmt.Println(sum(a, b))
}

// calcB 被推迟的函数按照后进先出（LIFO）的顺序执行。所以会输出结果是 7 3
func calcB() {
	a, b := 1, 2
	defer fmt.Println(sum(a, b))

	c, d := 3, 4
	defer fmt.Println(sum(c, d))
}

func sum(a, b int) int {
	return a + b
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

// b() 输出结果是：
//entering: b
//in b
//entering: a
//in a
//leaving: a
//leaving: b
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func printA(a int) int {
	fmt.Println(a)
	return a
}

// deferReturn defer和return的执行顺序，输出结果为0 1，可以看出先执行return后面的方法，再执行的defer后面的方法
func deferReturn(a int) int {
	defer printA(a + 1)
	return printA(a)
}

type ask struct {
	a int
}

func printAsk(a *ask) {
	fmt.Println(a.a)
}

// deferReturn2 输出结果是1 10，说明defer推迟执行的方法的参数如果是一个指针类型，在defer之后对参数值进行修改，会影响到defer后面的方法实际执行时参数的内容
func deferReturn2(a *ask) {
	printAsk(a)
	defer printAsk(a)
	a.a = 10
}
