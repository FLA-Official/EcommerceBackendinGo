package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

// GetProducts handles GET /products and returns a list of products.
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	utils.SendData(w, database.List(), http.StatusOK)
}
