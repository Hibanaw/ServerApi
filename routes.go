package main

import (
	"serverapi/admin"
	"serverapi/function"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/kemono", function.RandomKemono)
	r.GET("/admin/upload", admin.Upload)
	r.Static("/assets", "./assets")
}
