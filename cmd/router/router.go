package router

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/cmd/handler"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/customers"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/invoices"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/sales"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r, r.Group("/api/v1"), db}
}

func (r *router) MapRoutes() {
	r.buildCustomersRoutes()
	r.buildInvoicesRoutes()
	r.buildProductsRoutes()
	r.buildSalesRoutes()
}

func (r *router) buildCustomersRoutes() {

	repo := customers.NewRepository(r.db)
	service := customers.NewService(repo)
	handler := handler.NewHandlerCustomers(service)

	c := r.rg.Group("/customers")
	{
		c.GET("", handler.GetAll())
		c.POST("", handler.Post())
		c.POST("/insert", handler.InsertMany())
	}
}

func (r *router) buildInvoicesRoutes() {
	repo := invoices.NewRepository(r.db)
	service := invoices.NewService(repo)
	handler := handler.NewHandlerInvoices(service)

	i := r.rg.Group("/invoices")
	{
		i.GET("", handler.GetAll())
		i.POST("", handler.Post())
	}
}

func (r *router) buildProductsRoutes() {
	repo := products.NewRepository(r.db)
	service := products.NewService(repo)
	handler := handler.NewHandlerProducts(service)

	p := r.rg.Group("/products")
	{
		p.GET("", handler.GetAll())
		p.POST("", handler.Post())
		p.POST("/insert", handler.InsertMany())
	}
}

func (r *router) buildSalesRoutes() {
	repo := sales.NewRepository(r.db)
	service := sales.NewService(repo)
	handler := handler.NewHandlerSales(service)

	s := r.rg.Group("/sales")
	{
		s.GET("", handler.GetAll())
		s.POST("", handler.Post())
	}
}
