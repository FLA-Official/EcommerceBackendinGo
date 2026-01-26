package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Using Post http Method to add product from user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Please provide valid json", 400)
		http.Error(w, "Request Data Invalid", http.StatusBadRequest)
		return
	}
	createdUser := newUser.Store()

	// creating encoder object
	utils.SendData(w, createdUser, http.StatusCreated)

}
