package services

import (
	"gin-app/dtos"
	"gin-app/errs"
	"gin-app/models"
	"gin-app/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepo
}

func NewProductService(productRepo *repository.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (h *ProductService) CreateProduct(input *dtos.CreateProductInput) *models.Product {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
	}
	if err := h.productRepo.Create(&product); err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	return &product
}

func (h *ProductService) UpdateProduct(id int, input *dtos.UpadteProductInput) *models.Product {
	product, err := h.productRepo.FindByID(id)
	if err != nil {
		panic(errs.NotFoundRequest())
	}

	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Image != "" {
		product.Image = input.Image
	}

	if err := h.productRepo.Update(product); err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	return product
}

func (h *ProductService) GetProduct(id int) *models.Product {
	product, err := h.productRepo.FindByID(id)
	if err != nil {
		panic(errs.NotFoundRequest())
	}
	return product
}

func (h *ProductService) GetProducts() []*models.Product {
	products, err := h.productRepo.FindAll()
	if err != nil {
		panic(errs.BadRequestError("Products not found", err))
	}
	return products
}
