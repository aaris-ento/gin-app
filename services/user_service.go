package services

import (
	"gin-app/dtos"
	"gin-app/errs"
	"gin-app/models"
	"gin-app/repository"
	"gin-app/utils"
	"time"
)

type UserService struct {
	userRepo    *repository.UserRepo
	taskService *TaskService
}

func NewUserService(userRepo *repository.UserRepo, taskService *TaskService) *UserService {
	return &UserService{userRepo: userRepo, taskService: taskService}
}

func (s *UserService) Register(input *dtos.RegisterUserInput) *models.User {
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		panic(errs.InternalError("Could not hash password", err))
	}
	user := models.User{
		Email:    input.Email,
		Password: hashPassword,
		Role:     input.Role,
	}

	err = s.userRepo.Create(&user)
	if err != nil {
		panic(errs.InternalError("Could not register user", err))
	}

	emailVerification := models.UserToken{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(30 * time.Minute),
		Type:      models.UserTokenType(models.VerifyEmail),
	}
	s.userRepo.CreateVerificationTokenRecord(&emailVerification)

	s.taskService.SendEmailVerification(user.Email, emailVerification.Token)
	return &user
}

func (s *UserService) GetUserById(id uint) *models.User {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		panic(errs.BadRequestError("Could not find user", err))
	}

	return user
}

func (s *UserService) Login(input *dtos.LoginUserInput) *string {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		panic(errs.BadRequestError("User does not exists", err))
	}
	if b := utils.CheckPasswordHash(input.Password, user.Password); !b {
		panic(errs.BadRequestErrorMsg("Password mismatch"))
	}

	token := utils.GenerateJWT(user)
	return &token
}

func (s *UserService) VerifyEmail(token string) {
	emailVerification, err := s.userRepo.GetUserToken(token)
	if err != nil {
		panic(errs.BadRequestError("Invalid or expired token", err))
	}

	if emailVerification.ExpiresAt.Before(time.Now()) {
		panic(errs.BadRequestErrorMsg("Token has expired"))
	}

	user, err := s.userRepo.FindByID(emailVerification.UserID)
	if err != nil {
		panic(errs.BadRequestErrorMsg("User not found"))
	}

	user.EmailVerified = true
	s.userRepo.Update(user)
	s.userRepo.DeleteUserToken(token)
}

func (s *UserService) ForgotPassword(email string) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		panic(errs.BadRequestErrorMsg("User not found"))
	}

	resetPasswordToken := models.UserToken{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(30 * time.Minute),
		Type:      models.UserTokenType(models.PasswordReset),
	}

	s.userRepo.CreateVerificationTokenRecord(&resetPasswordToken)
	s.taskService.SendPasswordResetEmail(user.Email, resetPasswordToken.Token)
}

func (s *UserService) ResetPassword(token string, password string) {
	userToken, err := s.userRepo.GetUserToken(token)
	if err != nil {
		panic(errs.BadRequestErrorMsg("Token is invalid or expired"))
	}

	user, err := s.userRepo.FindByID(userToken.UserID)
	if err != nil {
		panic(errs.BadRequestErrorMsg("User does not exists"))
	}

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		panic(errs.InternalError("Could not hash password", err))
	}
	user.Password = hashPassword
	s.userRepo.Update(user)
}
