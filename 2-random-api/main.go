package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func randomNumHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa((rand.Intn(6) + 1))))
	return
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", randomNumHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
