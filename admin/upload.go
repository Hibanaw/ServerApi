package admin

import (
	"crypto/md5"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"serverapi/model"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Upload(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./database/kemono.db"), &gorm.Config{})
	db.AutoMigrate(&model.Image{})
	if err != nil {
		log.Printf("Database err:%s", err.Error())
		return
	}

	filepath.Walk("/mnt/e/Pictures/kemono/", func(path string, info os.FileInfo, err error) error {
		// fmt.Printf("%s\n", path)
		mt, _ := mimetype.DetectFile(path)
		m := mt.String()
		if m[:5] != "image" {
			return nil
		}
		value := model.Image{}
		md5 := GetFileMD5(path)
		db.Limit(1).Find(&value, "md5 = ?", md5)
		if value.Md5 != "" {
			fmt.Printf("重复\n")
			return nil
		}
		file, _ := os.Open(path)
		defer file.Close()
		img, _, _ := image.DecodeConfig(file)
		h := img.Height
		w := img.Width
		var d string
		if h > w {
			d = "vertical"
		}
		if h < w {
			d = "horizontal"
		}
		if h == w {
			d = "square"
		}
		db.Create(&model.Image{
			Md5:       md5,
			Height:    h,
			Width:     w,
			Direction: d,
			Pixel:     h * w,
			Mime_type: m,
		})
		return nil
	})
}

func GetFileMD5(pathName string) string {
	f, err := os.Open(pathName)
	if err != nil {
		fmt.Println("Open", err)
		return ""
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return ""
	}
	has := md5hash.Sum(nil)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
