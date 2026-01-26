package database

// if the variable name starts with small letter then this struct will be private and can be accessed only by main package
// otherwise it will be public
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgurl"`
}
