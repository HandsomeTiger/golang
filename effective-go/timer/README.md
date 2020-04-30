## time包实现定时器
golang的time包的Timer和Ticker都可以实现定时功能，两者的区别是Timer用于定时执行一次的任务，Ticker用于循环定时执行多次任务。  
Stop用于关闭定时器，在关闭后，将不会往定时器发送信息。**Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。**  
#### 基本用法
```golang
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

```

#### 源码学习
* [【源码学习】time.Timer 和 time.Ticker](https://my.oschina.net/u/2004526/blog/3042442)
* [golang进阶(八)——隐藏技能go:linkname](https://blog.csdn.net/lastsweetop/article/details/78830772)
