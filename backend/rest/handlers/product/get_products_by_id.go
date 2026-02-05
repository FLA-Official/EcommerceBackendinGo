package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

// GetProductsByID handles GET /products/{id} and returns the requested product if found.
func (h *Handler) GetProductsByID(w http.ResponseWriter, r *http.Request) {
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
	// Return the found product with 200 OK
	utils.SendData(w, product, http.StatusOK)
}
