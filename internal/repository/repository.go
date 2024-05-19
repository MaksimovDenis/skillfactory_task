package repository

import (
	"context"

	"Skillfactory_task_30.8.1/internal/models"
)

type TasksRepository interface {
	Get(ctx context.Context, userID int) (*models.Tasks, error)
}
