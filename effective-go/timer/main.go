package main

import (
	"fmt"
	"time"
)

func main() {
	//timerHandler()
	tickerHandler()
}

func echoDone() {
	fmt.Println("done")
}

// timerHandler 阻塞定时执行
func timerHandler() {
	// NewTimer 5s后过期，向<-chan发送一个过期时间
	tm := time.NewTimer(5 * time.Second)
	defer tm.Stop()
	// select 没有定义default，会一直阻塞直到从 <-tm.C中取到值
	select {
	// tm.C接收到过期时间之后
	case <-tm.C:
		// 执行echoDone
		echoDone()
	}
}

// tickerHandler 阻塞定时循环
func tickerHandler() {
	// NewTicker 每隔5s,想tk.C发送一个当前时间
	tk := time.NewTicker(5 * time.Second)
	defer tk.Stop()
	// for循环
	for {
		// 每次从tk.C取出一次执行echoDone
		select {
		case <-tk.C:
			echoDone()
		}
	}
}
