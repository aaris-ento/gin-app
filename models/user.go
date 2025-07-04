package models

import "time"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

type User struct {
	ID                 uint   `gorm:"primaryKey"`
	Email              string `gorm:"uniqueIndex"`
	Password           string
	Role               UserRole
	EmailVerified      bool        `gorm:"default:false"`
	VerificationTokens []UserToken `gorm:"foreignKey:UserID"`
}

type UserTokenType string

const (
	VerifyEmail   UserTokenType = "verify-password"
	PasswordReset UserTokenType = "password-reset"
)

type UserToken struct {
	ID        uint   `gorm:"primaryKey"`
	Token     string `gorm:"uniqueIndex;default:gen_random_uuid()"`
	ExpiresAt time.Time
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	Type      UserTokenType
}
