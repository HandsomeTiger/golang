package main

import "fmt"

func main (){
	var key int
	for key!=4{
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("          1 收支明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出软件")
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("收支明细")
			break
		case 2:
		case 3:
			fmt.Println("登记支出")
		case 4:
			fmt.Println("退出")
		default:
			fmt.Println("输入错误")
		}
	}


}
