package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ensamblaTec/learning/week2/problema4/pkg/models"
	"github.com/ensamblaTec/learning/week2/problema4/pkg/utils"
	"github.com/labstack/echo/v4"
)

var products = map[string][]*models.Product{
	"Products": {
		{
			ID:    1,
			Price: 20,
			Name:  "P1",
			Image: "https://i.giphy.com/media/AHj0lQstZ9I9W/giphy.webp",
		},
		{
			ID:    2,
			Price: 30,
			Name:  "P2",
			Image: "https://i.giphy.com/media/eSwGh3YK54JKU/giphy.webp",
		},
	},
}

func GetProducts() map[string][]*models.Product {
	return products
}

func RegisterProduct(c echo.Context) error {
	var image string
	productName := c.FormValue("productName")
	if len(productName) == 0 {
		log.Println("entro")
		return c.Render(http.StatusOK, "msgErrProductName", map[string]interface{}{
			"ErrorMensaje": "El campo Nombre no puede estar vacío.",
		})
	}
	price := c.FormValue("productPrice")
	// Source Image
	header := strings.Split(c.Request().Header.Get("Content-Type"), ";")
	if len(header) > 0 && header[0] == "multipart/form-data" {
		if file, err := c.FormFile("image"); err == nil {
			if image, err = utils.UploadFile(file); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		image = c.FormValue("image")
	}
	productPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return err
	}
	product := models.CreateProduct(
		len(products["Products"])+1,
		productName,
		image,
		productPrice,
	)
	products["Products"] = append(products["Products"], product)
	return c.Render(http.StatusOK, "product-cards-list", product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Request().Header.Get("Hx-Target")
	if len(id) == 0 {
		return echo.ErrBadRequest
	}

	str := id[1:]
	idConvert, err := strconv.Atoi(string(str))
	if err != nil {
		return echo.ErrInternalServerError
	}

	lw, sup := 0, len(products["Products"])-1
	prom := 0
	for lw <= sup {
		prom = 1 + (lw-sup)/2

		if products["Products"][prom].ID == idConvert {
			products["Products"] = append(products["Products"][0:prom], products["Products"][prom+1:]...)
			break
		}

		if prom < sup {
			sup = prom - 1
		} else {
			lw = prom + 1
		}
	}

	log.Println(products["Products"])
	log.Println(products["Products"][prom])
	log.Println(products["Products"][prom].Price)

	return c.Render(http.StatusOK, "product-cards-list", nil)
}
