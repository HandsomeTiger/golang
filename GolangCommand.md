# Go learning

### Command
- go help <command> : 获取go命令的帮助信息
- go version : 显示go版本信息  
- go env : 打印go的环境信息 
- go build : 用于编译源码文件代码包和依赖，编译源码文件  
> go build要生成可执行文件，go源码文件中必须有main包 并且main包中必须有main函数
`go build hello.go`
- go run : 编译并运行go源码文件  
`go run hello.go`
- go get : （依赖管理工具）从指定源上面下载或者更新指定的代码和依赖，并对他们进行编译和安装，类似于composer    
> go get 可以借助代码管理工具通过远程拉取或更新代码包及其依赖包，并自动完成编译和安装。通过git clone命令或者pull拉取代码不会进行编译和安装。      
`go get github.com/handsomeTiger/home`
- go install : 编译并安装代码包和依赖，编译源码文件并在bin目录下生成可执行文件，会把依赖包放在工作目录的pkg文件夹下  
`go install`
- go clean : 移除可执行go程序，比如 hello.exe文件  
`go clean hello.exe` 
- go fmt : 运行gofmt进行格式化
- go doc : 查看包文档
- go test : 运行测试 
> - 测试源文件是名称以“_test.go”为后缀的
> - 测试源文件内含若干测试函数的源码文件
> - 测试函数一般是以“Test”为名称前缀, 并有一个类型为“testing.T”的参数。
 
#### 
[go语言学习-常用命令](https://www.cnblogs.com/itogo/p/8645441.html)