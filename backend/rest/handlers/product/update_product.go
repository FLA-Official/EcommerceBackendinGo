package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// creating GET http Method to get data from database.
func (h *Handler) UpdateProducts(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	var newProduct database.Product
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Please provide valid json", 400)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = id
	// creating encoder object
	database.Update(newProduct)

	utils.SendData(w, "Successfully updated product", http.StatusCreated)
}
