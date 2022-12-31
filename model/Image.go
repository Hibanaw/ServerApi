package model

type Image struct {
	Md5       string `gorm:"primaryKey"`
	Mime_type string
	Height    int
	Width     int
	Pixel     int
	Direction string
}
