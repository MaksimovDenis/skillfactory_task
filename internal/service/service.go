package service

import (
	"context"

	"Skillfactory_task_30.8.1/internal/models"
)

type TasksService interface {
	Get(ctx context.Context, userID int) (*models.Tasks, error)
}
