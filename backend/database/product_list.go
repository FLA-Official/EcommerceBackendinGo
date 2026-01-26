package database

// creating private database
var productList []Product

// to add more data
func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

// to show all data
func List() []Product {
	return productList
}

// to get data by id
// if its find the product then it will return the product otherwise it will return nil
func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

// creating hardcoded Data which will run before the main code. In future we will use database. This is just for test purpose
func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "The king of fruit, fresh and sweet",
		Price:       2.89,
		ImgURL:      "https://www.mango.org/wp-content/uploads/2024/06/Mango_Varieties_hero_img.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: "A fruit available in all season",
		Price:       0.5,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQbgdJ5RnQe9bPm8jhKcmlV3FzGEtk_OhheqQ&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Jackfruit",
		Description: "The summer special national fruit of Bangladesh available, big, fresh and juicy",
		Price:       5.5,
		ImgURL:      "https://static.wikia.nocookie.net/fruit/images/a/ae/Jackfruit.jpg/revision/latest/thumbnail/width/360/height/450?cb=20241124050526",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "Sweet, sweet grapes",
		Price:       1.5,
		ImgURL:      "https://hips.hearstapps.com/hmg-prod/images/766/grapes-main-1506688521.jpg?resize=640:*",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Apple",
		Description: "Sweet Apple, available at reasonable price",
		Price:       2.00,
		ImgURL:      "https://assets.clevelandclinic.org/transform/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)

}
