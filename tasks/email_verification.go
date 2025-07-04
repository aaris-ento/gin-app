package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type EmailVerificationPayload struct {
	Email string
	Token string
}

func NewEmailVerificationTask(email string, token string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailVerificationPayload{
		Email: email,
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSendEmailVerification, payload), nil
}

type PasswordResetEmailPayload struct {
	Email string
	Token string
}

func PasswordResetEmailTask(email string, token string) (*asynq.Task, error) {
	payload, err := json.Marshal(PasswordResetEmailPayload{
		Email: email,
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSendPasswordResetEmail, payload), nil
}
