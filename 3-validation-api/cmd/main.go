package main

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/internal/verify"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifHandler(router, verify.VerifHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
