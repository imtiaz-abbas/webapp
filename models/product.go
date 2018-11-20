package models

type Product struct {
	ID	string `json:"id"`
	Name string `json:"name"`
	Price int	`json:"price"`
	Category string `json:"category"`
}

func getAllProducts() []Product {
	allProducts := []Product {
		Product {
			ID: "1",
			Name: "GoPro",
			Price: 1200,
			Category: "Photography",
		},
		Product {
			ID: "2",
			Name: "Tent",
			Price: 400,
			Category: "Trekking",
		},
		Product {
			ID: "3",
			Name: "Royal Enfield",
			Price: 800,
			Category: "Bikes",
		},
		Product {
			ID: "4",
			Name: "Bycicle",
			Price: 400,
			Category: "Trekking",
		},
		Product {
			ID: "5",
			Name: "Tripod",
			Price: 50,
			Category: "Photography",
		},
		Product {
			ID: "6",
			Name: "Helmet",
			Price: 100,
			Category: "Bikes",
		},
	}

	return allProducts
}

func GetAllProducts() []Product {
	return getAllProducts()
}

func GetProduct(id string) Product {
	allProducts := getAllProducts();
	result := allProducts[0]
	for _, item := range getAllProducts() {
		if(item.ID == id) {
			return item
		}
	}
	return result
}