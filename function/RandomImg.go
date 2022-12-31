package function

import (
	"fmt"
	"math/rand"
	"net/http"
	"serverapi/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RandomImg(c *gin.Context, apiname string) {
	var randomImg model.Image
	var imgs []model.Image
	// database
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("./database/%s.db", apiname)), &gorm.Config{})
	db.AutoMigrate(&model.Image{})
	if err != nil {
		//log.Printf("Database err:%s", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// api
	size := c.Query("size")
	direction := c.Query("direction")
	apitype := c.Query("type")

	// database
	tdb := db
	switch size {
	case "small":
		tdb = tdb.Where("size < 1000000")
	case "large":
		tdb = tdb.Where("size >= 1000000")
	}
	switch direction {
	case "vertical":
		tdb = tdb.Where("direction = ?", "vertical")
	case "horizontal":
		tdb = tdb.Where("direction = ?", "horizontal")
	}
	tdb.Find(&imgs)

	// rand
	l := len(imgs)
	rand.Seed(time.Now().Unix())
	randomImg = imgs[rand.Intn(l)]

	// return
	switch apitype {
	case "json":
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"imgsrc": randomImg.Md5,
			"width":  fmt.Sprintf("%d", randomImg.Width),
			"height": fmt.Sprintf("%d", randomImg.Height),
		})
	default:

	}
}
