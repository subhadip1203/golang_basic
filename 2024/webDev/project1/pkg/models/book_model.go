package model

import (
	"github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string ` json:"name" `
	Author      string ` json:"author" `
	Publication string ` json:"publication" `
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
