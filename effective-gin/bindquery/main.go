package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("", testB)
	s.Run(":8080")
}

func testB(ctx *gin.Context) {
	in := struct {
		Name *uint
		Age  *uint
	}{}
	fmt.Printf("%+v", in)
	ctx.ShouldBindQuery(&in)
	fmt.Printf("%+v", in)
}
