package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"go-rest/controllers"
	"go-rest/models"
)

func main() {
	fmt.Println("RESTful API")

	// database
	db, err := models.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// server
	router := createRouter()
	server := &http.Server{
		Handler: router,
		Addr: "0.0.0.0:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Create router
func createRouter() (*mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("RESTful API"))
	})

	// users controller
	users := &controllers.Users{}
	router.HandleFunc("/users", users.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", users.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", users.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", users.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", users.GetUsers).Methods(http.MethodGet)

	return router
}