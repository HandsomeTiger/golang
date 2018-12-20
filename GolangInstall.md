# Go learning

## Go安装

#### 安装
依据文档安装  

#### GOPATH
设置$GOPATH全局变量  
GOPATH 环境变量指定了你的工作空间位置。它或许是你在开发Go代码时， 唯一需要设置的环境变量。

首先创建一个工作空间目录，并设置相应的 GOPATH。你的工作空间可以放在任何地方， 在此文档中我们使用 $HOME/work。注意，它绝对不能和你的Go安装目录相同。 （另一种常见的设置是 GOPATH=$HOME。）
```go
$ mkdir $HOME/work
$ export GOPATH=$HOME/work
```
作为约定，请将此工作空间的 bin 子目录添加到你的 PATH 中：  
`$ export PATH=$PATH:$GOPATH/bin`

#### 包路径的设置和规范
标准库中的包有给定的短路径，比如 "fmt" 和 "net/http"。 对于你自己的包，你必须选择一个基本路径，来保证它不会与将来添加到标准库， 或其它扩展库中的包相冲突。

如果你将你的代码放到了某处的源码库，那就应当使用该源码库的根目录作为你的基本路径。 例如，若你在 GitHub 上有账户 github.com/user 那么它就应该是你的基本路径。

#### GOBIN
设置$GOBIN全局变量

#### 测试安装
```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
#### 参考文档
[Go语言中文社区](http://docscn.studygolang.com/doc/install)

