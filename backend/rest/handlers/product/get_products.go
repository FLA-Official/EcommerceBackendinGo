package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

// creating GET http Method to get data from database.
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	utils.SendData(w, database.List(), http.StatusOK)
}
