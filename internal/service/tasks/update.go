package tasks

import (
	"context"

	"skillfactory_task/internal/models"
)

func (s *serv) UpdateTaskById(ctx context.Context, task *models.UpdateTask) error {
	err := s.tasksRepository.UpdateTaskById(ctx, task)
	if err != nil {
		return err
	}

	return nil
}
