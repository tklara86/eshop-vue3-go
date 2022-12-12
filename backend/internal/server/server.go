package server

import (
	"fmt"
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/tklara86/eshop-vue3-go/internal/category"
	"github.com/tklara86/eshop-vue3-go/internal/product"
)

type Server struct {
	engine *gin.Engine
	port   uint
}

type Config struct {
	Port uint
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	s := &Server{
		engine: engine,
		port:   config.Port,
	}
	engine.GET("/products", s.Products)
	engine.GET("/category", s.Categories)

	return s, nil

}

func (s *Server) Run() error {
	return s.engine.Run(fmt.Sprintf(":%d", s.port))
}

func (s *Server) Categories(c *gin.Context) {
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, categories)
}

func (s *Server) Products(c *gin.Context) {
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

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, products)
}
