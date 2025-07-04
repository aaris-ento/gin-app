package bootstrap

import (
	"gin-app/tasks"
	"log"

	"github.com/hibiken/asynq"
)

func InitAsynqClient() *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	go func() {
		srv := asynq.NewServer(asynq.RedisClientOpt{Addr: "localhost:6379"}, asynq.Config{Concurrency: 5})
		mux := asynq.NewServeMux()
		mux.HandleFunc(tasks.TypeSendEmailVerification, tasks.VerifyEmailTaskHandler())
		mux.HandleFunc(tasks.TypeSendPasswordResetEmail, tasks.ResetEmailTaskHandler())

		log.Println("Asynq worker server starting")
		if err := srv.Run(mux); err != nil {
			log.Fatalf("Asynq worker crashed: %v", err)
		}
	}()

	return client
}
