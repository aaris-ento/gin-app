package errs

import "net/http"

type InternalErrorType struct {
	StatusCode int
	Code       string
	Message    string
	Error      error
}

func InternalError(msg string, err error) *BadRequest {
	return &BadRequest{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_ERROR", Message: msg, Error: err}
}

func InternalErrorMsg(msg string) *BadRequest {
	return &BadRequest{StatusCode: http.StatusInternalServerError, Code: "INTERNAL_ERROR", Message: msg}
}
