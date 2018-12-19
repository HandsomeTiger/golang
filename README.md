# golang learning

## Go语言第一课
#### Go语言基本数据类型
##### Go语言-程序实体与关键字  
　　go语言源码文件是由若干个程序实体组成。
　　在GO语言中，任何变量、函数、常量、结构体和接口都被统称为“程序实体”，他们的名字统称为“标示符”。
    
　　**注意：**  
    在Go语言中，我们对程序实体的访问权限控制只能通过它们的名字来实现。
    名字首字母为**大写（public）**的程序实体可以被任何代码包中的代码访问到。
    而名字首字母为**小写（private** 的程序实体则只能被同一个代码包中的代码所访问。 
    
　　**GO语言关键字：**    
    程序声明：import、package
    程序实体声明和定义：chan、const、func、interface、map、struct、type、var
    程序流程控制：go、select、break、case、continue、default、defer、else、fallthrough、for、goto、if、range、return、switch

##### Go语言-变量和常量  
　　**变量声明（var）**  
  > 注释：普通赋值，由关键字var、变量名称、变量类型、特殊标记=，以及相应的值组成。
   若只声明不赋值，则去除最后两个组成部分即可。  
   几种变量的声明方式：  
   ```GO
      var num1 int = 1  
      var num2, num3 int = 2, 3 // 注释：平行赋值  
      var ( // 注释：多行赋值
          num4 int = 4（注意多行赋值后面没有逗号）
          num5 int = 5
      )
   ````
   短变量声明：
   > : =  
    
   短变量只在当前作用域生效。  
   
   **常量的声明（const）**
   > 常量的声明跟变量的声明基本一致，只是常量的声明必须赋值。
   
##### Go语言-整数类型的命名和宽度  
   零值：0  
　　GO语言中的整数类型有10个，分别为：  
   > int  
   > uint  
   > int8  
   > uint8  
   > int16  
   > uint16  
   > int32  
   > uint32  
   > int64  
   > uint64  
   
   类型宽度指的是bit，1字节=8bit  
   
   数值范围：
   > 8bit -128-127 0-255  
   
   8进制和16进制的表示：  
   例如 一个值为12的int类型的表示为  
   > var num int = 12 (10进制)  
   var num int = 014 (8进制，前面的0表示8进制表示法)  
   var num int = 0xC (16进制，0x表示是16进制表示法)  
   
   ##### Go语言-浮点数类型
   零值：0.0  
   两种浮点数类型，分别是：  
   float32（4字节），  float64（8字节）。  
   **在Go语言里，浮点数的相关部分只能由10进制表示法表示，而不能由8进制表示法或16进制表示法表示。比如，03.7表示的一定是浮点数3.7。**
   
   ##### Go语言-复数类型  
   complex64（8字节）和complex128（16字节）。   
   > z = a(实部) + b(虚部)i(虚数单位)  
   
   ##### byte与rune
   byte是uint8（0~255）的别名类型，而rune（-2147483648~ 2147483647）则是int32的别名类型。  
   rune类型可以表示unicode编码要用单引号引起来  
   
   byte 等同于int8，常用来处理ascii字符
   rune 等同于int32,常用来处理unicode或utf-8字符
   
   ##### Go语言-字符串类型
   零值：\'\'  
   string 反引号“`”原生表示法，双引号“"”解释型表示法。
  
   ##### 参考资料
   [【golang】浅析rune数据类型](https://www.jianshu.com/p/4fbf529926ca)
  
   #### GO语言高级数据类型  
   ##### Go语言-数组类型（array）  
   数组类型的声明：  
   `type MyNumbers [3]int`  
   > 注：类型声明语句由关键字type、类型名称和类型字面量组成。   
        
   在这个数组声明中，type表示的是声明类型的关键字 MyNumber表示的是类型的名称,\[3\] 表示的是类型的长度，int表示字面量的类型.  
   类型长度可以用\[...\]来省略长度的数字。  
   声明一个数组并赋值给一个变量的语法是：  
   `var numbers = [3]int{1,2,3}` 或者 `var numbers = [...]int{1,2,3}`  
   获取数组中的元素，可以用下标索引的方式，如：  
   `numbers[0]` //得到数组的第一个元素。
   通过赋值可以给数组中的某个索引的值赋值：  
   `number[1] = 4`   
   获取数组的长度：用内置函数**len(numebrs)**。  
   
   ##### Go语言-切片类型（slice）  
   零值：nil  
   切片类型的声明：  
   `type MySlice []int`  
   与数组的声明类似，只是没有类型长度。  
   切片表达式一般由字符串、数组或切片的值以及由方括号包裹且由英文冒号 **“:”** 分隔的两个正整数组成。  
   `var slice1 = numbers3[1:4]`  \[1:4\]表示元素上界索引和下界索引，类似于php中的array_slice函数；  
   **注意，被“切下”的部分不包含元素上界索引指向的元素。**  
   为了获取数组、切片或通道类型的值的容量，我们可以使用内建函数**cap**  
   一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值  
   
   ##### Go语言-切片的更多操作方法
   在切片表达式中加入了第三个索引（即容量**上界索引**），如：  
   `var slice1 = numbers3[1:4:4] //[]{2,3,4}`   
   **append**会对切片值进行*扩展*并返回一个新的切片值  
   `slice1 = append(slice1, 6, 7) //[]{2,3,4,6,7}` 
   **copy**会把两个类型相同的值作为参数，第二个参数的元素复制到第一个参数中的响应的索引对应的位置。第二个参数的元素一定要小于等于第一个参数。  
   
   ##### Go语言-字典类型（map）
   零值：nil  
   字典类型相当于php里的关联数组类型（键值对）   
   字典类型的声明：  
   `map[K]T //K表示键的类型，T表示值的类型`  
   例如：  
   `map[int][string]` 表示键的类型为int，值的类型为string类型的字典类型。  
   `map[int]string{1: "a", 2: "b", 3: "c"}`  
   *对于字典值来说，如果其中不存在索引表达式欲取出的键值对，那么就以它的值类型的空值（或称默认值）作为该索引表达式的求值结果*  
   用**delete**删除字典中的键值对：（**有则删除，无则不做**）
   `delete(map,4)` //类似于php的unset  
    针对字典的索引表达式可以有两个求值结果。第二个求值结果是bool类型的。它用于表明字典值中是否存在指定的键值对。
    `e,ok :=map[5]`  
   
   ##### Go语言-通道类型（Channel）
   零值：nil  
   > 类似于**栈？**，先进先出？array_pop,array_shift(php)    
   
   特点：是Go语言中一种独特的数据结构，特点是**并发**且**安全**。  
   通道类型的声明表示：  
   `chan T` //关键词为chan，T表示通道允许传递的数据类型。  
   make函数可接受两个参数。第一个参数是代表了将被初始化的值的类型的字面量（比如chan int），而第二个参数则是值的长度。例如，若我们想要初始化一个长度为5且元素类型为int的通道值，则需要这样写：  
   `make(chan int, 5)`  
   向通道发送数据的用法：  
   `ch <- "value"`  
   从通道取出数据的用法：  
   `<- ch`  
   通道类型的取值也有两个结果值：
   `value,ok := <- ch` ok表示的是通道值的状态，是否关闭，返回的也是个bool值。  
   关闭通道：  
   `close(ch1)`   
   
   **注意：**
   - 通道值已满会阻塞
   - 重复关闭通道，引起运行时恐慌
   - 向已满的通道值中发送数据 ，引起运行时恐慌
   - 从已空的通道值中取值，阻塞
   
   ##### Go语言-通道的更多种类
   通道分为**缓冲通道**和**非缓冲通道**。  
   >      非缓冲通道：发送方在向通道值发送数据的时候会立即被阻塞，直到有某一个接收方已从该通道值中接收了这条数据
   非缓冲通道的声明：  
   `make(chan int, 0)`  
   单向管道，例如：  
   `type Receiver <-chan int`
   
   #### Go语言-高级数据类型2
   ##### Go语言-函数（func）  
   函数类型的字面量由关键字func、由圆括号包裹**参数声明**列表、空格以及可以由圆括号包裹的**结果声明**列表组成。  
   函数的声明：  
   `func(input1 string,input2 string) string`
   函数声明中的参数名称和结果名称都可以统一省略，并且在只有一个无名称的结果声明时还可以省略括号  
   函数类型的声明：  
   `type MyFunc func(input1 string ,input2 string) string`  
   函数的写法：  
   ```go
        func myFunc(part1 string, part2 string) string {
            return part1 + part2
        }  
   ```
   上述函数 myFunc是MyFunc函数类型的实现。  
   
   ##### Go语言-结构体（struct）和方法
   > 类似于php中的类class      
   
   结构体的声明：
    ```go
    type Person struct {
        Name   string
        Gender string
        Age    uint8
    }
    ```
   结构体的值：  
   `Person{Name: "Robert", Gender: "Male", Age: 33} `   
   或  
    `Person{"Robert", "Male", 33} `  
   匿名结构体和匿名函数  
   ......  
   
   结构体方法的声明：  
   ```
    func (person *Person) Grow() {
        person.Age++
    } 
   ``` 
   >其中 (person *Person)类似于class类中的$this，可以在方法func中调用Person里的属性方法，\*表示的是指针类型。
   
   **注意**：Go语言中不存在继承关系，但是可以模仿继承。
   
   ##### Go语言-接口（interface）
   > 跟php中的接口的定义相似    
   
   接口的声明：
   ```go
    type Animal interface {
        Grow()
        Move(string) string
    }
   ```
   type关键字 + 类型名称 + interface关键字  
   完全实现了接口中定义的方法的结构体就是该接口的实现    
   
   空接口类型:  
   不包含任何方法声明的接口类型，简称空接口。  
   Go语言中的包含预定义的任何数据类型都可以被看做是空接口的实现。比如：  
   ```go
   p := Person{"Robert", "Male", 33, "Beijing"}
   v := interface{}(&p)
   ```
   
   类型转换：  
   如上p本来是Person类型的一个实现，把他转换成接口类型，就可以进行断言来判断他是不是接口Animal的一个实现，继而可以判断结构体Person是不是Animal接口的实现。
   
   关于断言：  
   > tip：我们是不能在一个非接口类型的值上应用类型断言来判定它是否属于某一个接口类型的。
   `h,ok = v.(Animal)`
   
   ##### Go语言-指针（ \& 和 * ）
   
   > 一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。    
   
   #### Go语言-流程控制语句
   ##### if语句
   if语句跟php不太一样的地方是if里面可以有该if作用于下的一个临时变量    
   
   ##### switch case语句
   从上往下查找，当找到与表达式值相等的case就停止，不再查找下面的case  
   fallthrough：既是一个关键字，又可以代表一条语句。表示的跳过当前 case 执行下个case。  
   fallthrough有两点需要注意，fallthrough语句出现的时候一定是当前case的最后一个语句 也就是它后面不能有别的执行，第二点是 执行跳过的这个case之后一定还有别的case。  
   
   ##### for语句
   知识点：range 类似于foreach，continue，break
   
   ##### select语句
   它只能用于通道。
   ```go
   ch1 := make(chan int, 1)
   ch2 := make(chan int, 1)
   // 省略若干条语句
   select {
   case e1 := <-ch1:
       fmt.Printf("1th case is selected. e1=%v.\n", e1)
   case e2 := <-ch2:
       fmt.Printf("2th case is selected. e2=%v.\n", e2)
   default:
       fmt.Println("No data!")
   } 
   ```
   
   
   
   
   
   
   
    
   
   
   
   
   
   
   
   