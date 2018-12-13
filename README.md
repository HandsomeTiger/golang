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
   string 反引号“`”原生表示法，双引号“"”解释型表示法。
  
   ##### 参考资料
   [【golang】浅析rune数据类型](https://www.jianshu.com/p/4fbf529926ca)
   
   
   