package handlers

import (
	"gin-app/dtos"
	"gin-app/errs"
	"gin-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input dtos.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	user := h.userService.Register(&input)

	c.JSON(http.StatusCreated, gin.H{
		"user": dtos.NewUserResponse(user),
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input dtos.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	token := h.userService.Login(&input)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userIdVal, ok := c.Get("user_id")
	if !ok {
		panic(errs.UnAuthorizedError())
	}

	userId, ok := userIdVal.(uint)
	if !ok {
		panic(errs.UnAuthorizedError())
	}

	user := h.userService.GetUserById(userId)

	c.JSON(http.StatusOK, gin.H{
		"user": dtos.NewUserResponse(user),
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	requesterRoleVal, ok := c.Get("role")
	if !ok {
		panic(errs.UnAuthorizedError())
	}

	requesterRole, ok := requesterRoleVal.(string)
	if !ok {
		panic(errs.BadRequestErrorMsg("Invalid role"))
	}

	targetUserIdVal := c.Param("id")
	targetUserId, err := strconv.Atoi(targetUserIdVal)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	targetUser := h.userService.GetUserById(uint(targetUserId))
	if targetUser == nil {
		panic(errs.NotFoundRequest())
	}

	if requesterRole == "admin" {
		// Admin can get anyone
		c.JSON(http.StatusOK, gin.H{"user": targetUser})
		return
	}

	if requesterRole == "customer" && targetUser.Role == "customer" {
		// Customers can only get customers
		c.JSON(http.StatusOK, gin.H{"user": dtos.NewUserResponse(targetUser)})
		return
	}

	panic(errs.UnAuthorizedError())
}

func (h *UserHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		panic(errs.BadRequestErrorMsg("Token is empty"))
	}

	h.userService.VerifyEmail(token)
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func (h *UserHandler) ForgotPassword(c *gin.Context) {
	var input dtos.ForgotPasswordInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	h.userService.ForgotPassword(input.Email)
	c.JSON(http.StatusOK, "Password reset link sent successfully")
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		panic(errs.BadRequestErrorMsg("Token is empty"))
	}
	var input dtos.ResetPasswordInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(errs.InvalidInputRequest(err))
	}

	h.userService.ResetPassword(token, input.Password)
	c.JSON(http.StatusOK, "Password changed successfully")
}
