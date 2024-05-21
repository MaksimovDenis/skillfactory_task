package service

import (
	"context"

	"Skillfactory_task_30.8.1/internal/models"
)

type TasksService interface {
	CreateTask(ctx context.Context, task *models.CreateTasks) (int, error)
	GetTasks(ctx context.Context) (*[]models.Tasks, error)
	GetTaskByAuthorName(ctx context.Context, authorName string) (*models.Tasks, error)
	GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error)
}
