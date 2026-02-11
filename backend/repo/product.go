// Package repo provides the repository layer for managing product data.
// It implements the repository pattern with in-memory storage for product operations.
// In the future, this will be integrated with a persistent database.
package repo

// Product represents a product entity in the ecommerce system.
// It contains all essential information about a product including its identification,
// description, pricing, and image reference.
type Product struct {
	// ID is the unique identifier for the product
	ID int `json:"id"`
	// Title is the name of the product
	Title string `json:"title"`
	// Description provides detailed information about the product
	Description string `json:"description"`
	// Price is the cost of the product
	Price float64 `json:"price"`
	// ImgURL is the URL to the product's image
	ImgURL string `json:"imgurl"`
}

// ProductRepo defines the interface for product repository operations.
// It provides methods for CRUD operations on products.
type ProductRepo interface {
	// Create adds a new product to the repository and returns the created product
	Create(p Product) (*Product, error)
	// Get retrieves a single product by its ID
	Get(productID int) (*Product, error)
	// List returns all products in the repository
	List() []*Product
	// Delete removes a product by its ID
	Delete(productID int) error
	// Update modifies an existing product's information
	Update(product Product) (*Product, error)
}

// productRepo is the concrete implementation of the ProductRepo interface.
// It stores products in memory for testing purposes.
type productRepo struct {
	// productList holds all products in memory
	productList []*Product
}

// NewProductRepo creates and initializes a new product repository instance.
// It returns a ProductRepo interface implementation pre-loaded with sample product data.
// This is used for testing purposes and will be replaced with database integration in the future.
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProduct(repo)
	return repo
}

// Create adds a new product to the repository.
// It automatically assigns the next sequential ID based on the current list length.
// Returns a pointer to the newly created product or an error if the operation fails.
func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

// Get retrieves a product by its ID.
// Returns a pointer to the found product, or nil if no product with the given ID exists.
func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}

// List returns a slice of pointers to all products in the repository.
func (r *productRepo) List() []*Product {
	return r.productList
}

// Delete removes a product from the repository by its ID.
// It creates a filtered list excluding the product with the matching ID.
// Returns nil as error for successful deletion.
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList

	return nil
}

// Update modifies an existing product's information in the repository.
// It searches for a product with the matching ID and replaces it with the updated version.
// Returns a pointer to the updated product or an error if the operation fails.
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil
}

// generateInitialProduct populates the repository with sample product data.
// This helper function creates hardcoded initial products for testing purposes.
// In the future, this will be replaced with data loaded from a persistent database.
func generateInitialProduct(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Mango",
		Description: "The king of fruit, fresh and sweet",
		Price:       2.89,
		ImgURL:      "https://www.mango.org/wp-content/uploads/2024/06/Mango_Varieties_hero_img.png",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Banana",
		Description: "A fruit available in all season",
		Price:       0.5,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQbgdJ5RnQe9bPm8jhKcmlV3FzGEtk_OhheqQ&s",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Jackfruit",
		Description: "The summer special national fruit of Bangladesh available, big, fresh and juicy",
		Price:       5.5,
		ImgURL:      "https://static.wikia.nocookie.net/fruit/images/a/ae/Jackfruit.jpg/revision/latest/thumbnail/width/360/height/450?cb=20241124050526",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Grapes",
		Description: "Sweet, sweet grapes",
		Price:       1.5,
		ImgURL:      "https://hips.hearstapps.com/hmg-prod/images/766/grapes-main-1506688521.jpg?resize=640:*",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Apple",
		Description: "Sweet Apple, available at reasonable price",
		Price:       2.00,
		ImgURL:      "https://assets.clevelandclinic.org/transform/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)

}
