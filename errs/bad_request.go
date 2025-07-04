package errs

type BadRequest struct {
	StatusCode int
	Code       string
	Message    string
	Error      error
}

func BadRequestError(msg string, err error) *BadRequest {
	return &BadRequest{StatusCode: 400, Code: "BAD_REQUEST", Message: msg, Error: err}
}

func BadRequestErrorMsg(msg string) *BadRequest {
	return &BadRequest{StatusCode: 400, Code: "BAD_REQUEST", Message: msg}
}
