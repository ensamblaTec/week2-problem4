package models

type Product struct {
	Name  string
	Price float64
	Image string
}

func CreateProduct(name, image string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
		Image: image,
	}
}
