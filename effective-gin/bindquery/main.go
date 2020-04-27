package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("", bindQueryExample)
	s.Run(":8080")
}

type query struct {
	Name   string  `json:"name" form:"name"`
	Age    int     `json:"age" form:"age"`
	Email  *string `json:"email" form:"email"`
	Gender *int    `json:"gender" form:"gender"`
	Phone  *string
	Height *int
}

// curl
func bindQueryExample(ctx *gin.Context) {
	in := &query{}
	log.Printf("init in :%+v", in)
	if err := ctx.ShouldBindQuery(in); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	log.Printf("bind in :%+v", in)
	ctx.JSON(200, "ok")
}
