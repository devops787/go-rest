package controllers

import (
	"net/http"
)

type Users struct {

}

// Create user
func (users *Users) CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("create user"))
}

// Update user
func (users *Users) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("update user"))
}

// Delete user
func (users *Users) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("delete user"))
}

// Get user
func (users *Users) GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("get user"))
}

// Get users
func (users *Users) GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("get users"))
}