package errs

type NotFoundInput struct {
	StatusCode int
	Code       string
	Message    string
	Error      error
}

func NotFoundRequest() *NotFoundInput {
	return &NotFoundInput{StatusCode: 404, Code: "NOT_FOUND", Message: "Resource not found"}
}
