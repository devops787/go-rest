package controllers

import (
	"net/http"
	"go-rest/models"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Users struct {

}

// Create user
func (users *Users) CreateUser(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	user := &models.User{
		Name: request.PostForm["name"][0],
		Email: request.PostForm["email"][0],
	}

	models.GetDatabase().Create(user)

	response, _ := json.Marshal(user)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// Update user
func (users *Users) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	request.ParseForm()

	user := &models.User{}
	models.GetDatabase().First(user, params["id"])
	if user != nil {
		user.Name = request.PostForm["name"][0]
		user.Email = request.PostForm["email"][0]
	}
	models.GetDatabase().Save(user)

	response, _ := json.Marshal(user)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// Delete user
func (users *Users) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	user := &models.User{}
	models.GetDatabase().First(user, params["id"])
	models.GetDatabase().Delete(user)

	response, _ := json.Marshal(user)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// Get user
func (users *Users) GetUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	user := &models.User{}
	models.GetDatabase().First(user, params["id"])
	response, _ := json.Marshal(user)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

// Get users
func (users *Users) GetUsers(writer http.ResponseWriter, request *http.Request) {
	results := make([]models.User, 0)
	models.GetDatabase().Find(&results)
	response, _ := json.Marshal(results)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}