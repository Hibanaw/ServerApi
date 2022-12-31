package function

import (
	"github.com/gin-gonic/gin"
)

func RandomKemono(c *gin.Context) {
	RandomImg(c, "kemono")
}
