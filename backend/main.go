package main

import (
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/tklara86/eshop-vue3-go/internal/category"
	"github.com/tklara86/eshop-vue3-go/internal/product"
)

func main() {
	r := gin.Default()

	products := []product.Product{
		{
			ID:               "1",
			Name:             "test",
			Description:      "this is product",
			PriceVATExcluded: money.New(40, money.GBP),
			VAT:              money.New(200, money.GBP),
		},
		{
			ID:               "2",
			Name:             "test2",
			Description:      "this is product",
			PriceVATExcluded: money.New(60, money.GBP),
			VAT:              money.New(200, money.GBP),
		},
		{
			ID:               "3",
			Name:             "test3",
			Description:      "this is product",
			PriceVATExcluded: money.New(105, money.GBP),
			VAT:              money.New(200, money.GBP),
		},
		{
			ID:               "3",
			Name:             "test3",
			Description:      "this is product",
			PriceVATExcluded: money.New(105, money.GBP),
			VAT:              money.New(200, money.GBP),
		},
	}

	categories := []category.Category{
		{
			ID:          "23",
			Name:        "categiory name 1",
			Description: "this is desc",
		},
		{
			ID:          "12",
			Name:        "categiory name 2",
			Description: "this is desc 2",
		},
	}

	r.GET("/products", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.JSON(http.StatusOK, products)
	})

	r.GET("/category", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.JSON(http.StatusOK, categories)
	})

	r.Run(":9090")
}
