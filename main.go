package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
)

func main() {
	fmt.Println("RESTful API")

	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("RESTful API"))
	})

	server := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}