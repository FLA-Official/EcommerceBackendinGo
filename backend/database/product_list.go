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
