package repository

import (
	"context"

	"Skillfactory_task_30.8.1/internal/models"
)

type TasksRepository interface {
	GetTaskByAuthorName(ctx context.Context, authorName string) (*models.Tasks, error)
	GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error)
}
