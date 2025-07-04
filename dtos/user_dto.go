package dtos

import "gin-app/models"

type RegisterUserInput struct {
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required,min=6"`
	Role     models.UserRole `json:"role" binding:"oneof=admin customer"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordInput struct {
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    uint            `json:"id"`
	Email string          `json:"email"`
	Role  models.UserRole `json:"role"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}
