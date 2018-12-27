package main

import (
	"fmt"

	"github.com/golearn/action/model"
)

//main is Enterance of go program.
func main() {
	customer := model.Resident{"test", 1222}
	customer.SetName("my name")
	name := customer.GetName()
	fmt.Println(name)
}

//Serve is a serve for http
func Serve() {

}
