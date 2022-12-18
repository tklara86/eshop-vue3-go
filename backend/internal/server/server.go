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
	Engine *gin.Engine
	port   uint
}

type Config struct {
	Port uint
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	s := &Server{
		Engine: engine,
		port:   config.Port,
	}
	engine.GET("/products", s.Products)
	engine.GET("/categories", s.Categories)

	return s, nil

}

func (s *Server) Run() error {
	return s.Engine.Run(fmt.Sprintf(":%d", s.port))
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
	price := money.New(200, money.GBP)
	price2 := money.New(400, money.GBP)

	products := []product.Product{
		{
			ID:          "1",
			Name:        "test",
			Image:       "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fashitani.jp%2Fgolangtips%2Fgopher.png&f=1&nofb=1&ipt=62ebfcd94342f8cd408afb77421c379495ac628b0a84ae0b6d8e36ef8cbc2ecb&ipo=images",
			Description: "this is product",
			PriceVATExcluded: product.Amount{
				Money:   price,
				Display: price.Display(),
			},
			VAT: product.Amount{
				Money:   price,
				Display: price.Display(),
			},
			TotalPrice: product.Amount{
				Money:   price2,
				Display: price2.Display(),
			},
		},
		{
			ID:          "2",
			Name:        "test2",
			Description: "this is product 2",
			Image:       "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fashitani.jp%2Fgolangtips%2Fgopher.png&f=1&nofb=1&ipt=62ebfcd94342f8cd408afb77421c379495ac628b0a84ae0b6d8e36ef8cbc2ecb&ipo=images",
			PriceVATExcluded: product.Amount{
				Money:   price,
				Display: price.Display(),
			},
			VAT: product.Amount{
				Money:   price,
				Display: price.Display(),
			},
			TotalPrice: product.Amount{
				Money:   price2,
				Display: price2.Display(),
			},
		},
		{
			ID:          "3",
			Name:        "test3",
			Image:       "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fashitani.jp%2Fgolangtips%2Fgopher.png&f=1&nofb=1&ipt=62ebfcd94342f8cd408afb77421c379495ac628b0a84ae0b6d8e36ef8cbc2ecb&ipo=images",
			Description: "this is product 3",
			PriceVATExcluded: product.Amount{
				Money:   price,
				Display: price.Display(),
			},
			VAT: product.Amount{
				Money:   price,
				Display: price.Display(),
			},
			TotalPrice: product.Amount{
				Money:   price2,
				Display: price2.Display(),
			},
		},
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, products)
}
