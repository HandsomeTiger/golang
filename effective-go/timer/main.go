package main

import (
	"fmt"
	"time"
)

func main() {
	//timerHandler()
	//tickerHandler()
}

// timerHandler 阻塞定时执行
func timerHandler() {
	// NewTimer 5s后过期，向<-chan发送一个过期时间
	tm := time.NewTimer(2 * time.Second)
	defer tm.Stop()
	// tm.C接收到过期时间之后
	<-tm.C
	fmt.Println("timer expired")

	tm2 := time.NewTimer(2 * time.Second)
	go func() {
		// tm2在先调用stop，所以2秒后不会往tm2发送过期时间
		<-tm2.C
		fmt.Println("timer expired 2")
	}()
	stop := tm2.Stop()
	if stop {
		fmt.Println("timer2 stop")
	}
	time.Sleep(5 * time.Second)
}

//tickerHandler 阻塞定时循环
func tickerHandler() {
	// NewTicker 每隔5s,想tk.C发送一个当前时间
	tk := time.NewTicker(5 * time.Second)
	defer tk.Stop()
	// for循环
	for {
		// 每次从tk.C取出一次执行echoDone
		select {
		case <-tk.C:
			fmt.Println("ticker expired")
		}
	}
}
