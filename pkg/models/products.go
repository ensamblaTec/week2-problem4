package models

type Product struct {
	ID    int
	Name  string
	Price float64
	Image string
}

func CreateProduct(id int, name, image string, price float64) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
		Image: image,
	}
}
