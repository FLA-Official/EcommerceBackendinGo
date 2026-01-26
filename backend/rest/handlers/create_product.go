package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Using Post http Method to add product from user
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Please provide valid json", 400)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	createdProduct := database.Store(newProduct)

	// creating encoder object
	utils.SendData(w, createdProduct, http.StatusCreated)

}
