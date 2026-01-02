package main

import (
	"validation-api/configs"
	"validation-api/internal/product"
	"validation-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	db.AutoMigrate(&product.Product{})
}
