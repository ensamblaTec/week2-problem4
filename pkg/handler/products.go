package handler

import (
	"net/http"
	"strconv"

	"github.com/ensamblaTec/learning/week2/problema4/pkg/models"
	"github.com/labstack/echo/v4"
)

var products = map[string][]*models.Product{
	"Products": {
		{
			Price: 20,
			Name:  "P1",
			Image: "2",
		},
	},
}

func GetProducts() map[string][]*models.Product {
	return products
}

func RegisterProduct(c echo.Context) error {
	productName := c.FormValue("productName")
	price := c.FormValue("productPrice")
	productPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return err
	}
	product := models.CreateProduct(productName, "image", productPrice)
	products["Products"] = append(products["Products"], product)
	return c.Render(http.StatusOK, "product-cards-list", product)
}
