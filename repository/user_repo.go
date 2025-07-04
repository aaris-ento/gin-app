package repository

import (
	"gin-app/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	dB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{dB: DB}
}

func (r *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.dB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.dB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindAll() ([]*models.User, error) {
	var users []*models.User
	if err := r.dB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Create(user *models.User) error {
	if err := r.dB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Update(user *models.User) error {
	if err := r.dB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) CreateVerificationTokenRecord(tokenRecord *models.UserToken) error {
	if err := r.dB.Create(tokenRecord).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetUserToken(token string) (*models.UserToken, error) {
	var tokenRecord *models.UserToken
	if err := r.dB.Where("token = ?", token).First(&tokenRecord).Error; err != nil {
		return nil, err
	}
	return tokenRecord, nil
}

func (r *UserRepo) DeleteUserToken(token string) (*models.UserToken, error) {
	tokenRecord := &models.UserToken{}
	if err := r.dB.Where("token = ?", token).First(tokenRecord).Error; err != nil {
		return nil, err
	}

	if err := r.dB.Delete(tokenRecord).Error; err != nil {
		return nil, err
	}

	return tokenRecord, nil
}
