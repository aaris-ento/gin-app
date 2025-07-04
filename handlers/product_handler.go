package handlers

import (
	"gin-app/dtos"
	"gin-app/errs"
	"gin-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input dtos.CreateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	product := h.productService.CreateProduct(&input)

	c.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var input *dtos.UpadteProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(errs.BadRequestErrorMsg("Invalid ID"))
	}
	product := h.productService.UpdateProduct(id, input)

	c.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	product := h.productService.GetProduct(id)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products := h.productService.GetProducts()
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
