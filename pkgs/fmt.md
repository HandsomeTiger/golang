# pkgs 
## 1.fmt
> fmt是go语言中的I/O

fmt.Printf("hello world %v\n","go") //格式化输出到屏幕 （类似于 echo）
fmt.Println("hellow","world","go") //
f := fmt.Sprintf("float %f",3.14159) //格式化输出到字符串 （这个是赋值）
fmt.Fprintln(os.Stdout,"A\n") //指定输出到接口
type Stringer