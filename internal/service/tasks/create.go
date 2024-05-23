package tasks

import (
	"context"

	"skillfactory_task/internal/models"
)

func (s *serv) CreateTask(ctx context.Context, task *models.CreateTask) (int, error) {
	taskId, err := s.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return 0, err
	}
	return taskId, nil
}
