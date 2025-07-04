package repository

import (
	"gin-app/models"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(DB *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: DB}
}

func (r *ProductRepo) FindByID(id int) (*models.Product, error) {
	var product *models.Product
	if err := r.DB.First(product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepo) FindAll() ([]*models.Product, error) {
	var products []*models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepo) Create(product *models.Product) error {
	if err := r.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) Update(product *models.Product) error {
	if err := r.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}