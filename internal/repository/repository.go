package repository

import (
	"context"

	"skillfactory_task/internal/models"
)

type TasksRepository interface {
	CreateTask(ctx context.Context, task *models.CreateTask) (int, error)
	GetTasks(ctx context.Context) (*[]models.Tasks, error)
	GetTaskByAuthorName(ctx context.Context, authorName string) (*models.Tasks, error)
	GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetTaskById(ctx context.Context, id int) (*models.Tasks, error)
	UpdateTaskById(ctx context.Context, task *models.UpdateTask) error
	DeleteTaskById(ctx context.Context, id int) error
	DeleteTaskLabelById(ctx context.Context, id int) error
}
