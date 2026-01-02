package main

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/internal/product"
	"validation-api/internal/verify"
	"validation-api/pkg/db"
	"validation-api/pkg/middleware"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	ProductRepo := product.NewProductRepository(db)
	verify.NewVerifHandler(router, verify.VerifHandlerDeps{
		Config: conf,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: ProductRepo,
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.LoggingMiddleware(router),
	}
	fmt.Println("Listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
