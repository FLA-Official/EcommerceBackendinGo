package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

// creating GET http Method to get data from database.
func GetProductsByID(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	product := database.Get(id)

	if product == nil {
		utils.SendData(w, "Product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, product, http.StatusNotFound)
}
