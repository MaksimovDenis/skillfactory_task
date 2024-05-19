package tasks

import (
	"context"

	"Skillfactory_task_30.8.1/internal/models"
)

func (s *serv) Get(ctx context.Context, userID int) (*models.Tasks, error) {
	task, err := s.tasksRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return task, nil
}
