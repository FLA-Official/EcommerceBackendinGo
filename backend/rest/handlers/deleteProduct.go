package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	database.Delete(id)

	utils.SendData(w, "Successfully deleted product", http.StatusCreated)
}
