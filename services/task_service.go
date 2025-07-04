package services

import (
	"gin-app/tasks"
	"log"

	"github.com/hibiken/asynq"
)

type TaskService struct {
	asynqClient *asynq.Client
}

func NewTaskService(asynqClient *asynq.Client) *TaskService {
	return &TaskService{asynqClient: asynqClient}
}

func (h *TaskService) SendEmailVerification(email string, token string) {
	task, err := tasks.NewEmailVerificationTask(email, token)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
	}

	_, err = h.asynqClient.Enqueue(task)
	if err != nil {
		log.Printf("Failed to enqueue verification email: %v", err)
	}
}

func (h *TaskService) SendPasswordResetEmail(email string, token string) {
	task, err := tasks.NewEmailVerificationTask(email, token)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
	}

	_, err = h.asynqClient.Enqueue(task)
	if err != nil {
		log.Printf("Failed to enqueue verification email: %v", err)
	}
}
