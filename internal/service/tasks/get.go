package tasks

import (
	"context"

	"Skillfactory_task_30.8.1/internal/models"
)

func (s *serv) GetTaskByAuthorName(ctx context.Context, authorName string) (*models.Tasks, error) {
	task, err := s.tasksRepository.GetTaskByAuthorName(ctx, authorName)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *serv) GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error) {
	tasks, err := s.tasksRepository.GetTaskByLabel(ctx, label)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
