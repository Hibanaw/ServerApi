package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Routes(r)
	r.Run(":9090")
}
