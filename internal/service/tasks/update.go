package tasks

import (
	"context"

	"skillfactory_task_30.8.1/internal/models"
)

func (s *serv) UpdateTaskById(ctx context.Context, task *models.UpdateTask) error {
	err := s.tasksRepository.UpdateTaskById(ctx, task)
	if err != nil {
		return err
	}

	return nil
}
