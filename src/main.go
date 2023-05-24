package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
func main() {

	engine := gin.Default()

	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg":  "hello from cicd-demo!",
		})
	})
	fmt.Print()

	engine.Run(":9088")

}
