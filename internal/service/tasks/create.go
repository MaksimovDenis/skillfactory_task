package tasks

import (
	"context"

	"skillfactory_task_30.8.1/internal/models"
)

func (s *serv) CreateTask(ctx context.Context, task *models.CreateTask) (int, error) {
	taskId, err := s.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return 0, err
	}
	return taskId, nil
}
