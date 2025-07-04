package errs

type InvalidInput struct {
	StatusCode int
	Code       string
	Message    string
	Error      error
}

func InvalidInputRequest(err error) *InvalidInput {
	return &InvalidInput{StatusCode: 400, Code: "BAD_REQUEST", Message: "Invalid input recieved", Error: err}
}
