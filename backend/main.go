package main

import (
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/tklara86/eshop-vue3-go/internal/product"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

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
	}

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, products)

	})

	r.Run(":9090")
}
