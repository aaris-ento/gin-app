package errs

import "net/http"

type UnAuthorizedErrorType struct {
	StatusCode int
	Code       string
	Message    string
	Error      error
}

func UnAuthorizedError() *NotFoundInput {
	return &NotFoundInput{StatusCode: http.StatusUnauthorized, Code: "UNAUTHORIZED", Message: "You are not authorized"}
}
