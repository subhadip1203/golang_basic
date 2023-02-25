package model

import "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/config"

func ModelMigration() {
	db = config.GetDB()
	db.AutoMigrate((&Book{}))
}
