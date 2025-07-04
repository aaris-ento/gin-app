package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

func VerifyEmailTaskHandler() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var payload EmailVerificationPayload
		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		go func() {
			verificationURL := fmt.Sprintf("http://localhost:8080/verify-email?token=%s", payload.Token)

			log.Printf("Verification link: %s", verificationURL)
			log.Printf("Email sent to %s", payload.Email)
		}()

		return nil
	}
}

func ResetEmailTaskHandler() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var payload EmailVerificationPayload
		if err := json.Unmarshal(t.Payload(), &payload); err != nil {
			return err
		}

		go func() {
			verificationURL := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", payload.Token)

			log.Printf("Verification link: %s", verificationURL)
			log.Printf("Email sent to %s", payload.Email)
		}()

		return nil
	}
}
