package tasks

import (
	"context"

	"skillfactory_task/internal/models"
)

func (s *serv) GetTasks(ctx context.Context) (*[]models.Tasks, error) {
	tasks, err := s.tasksRepository.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

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

func (s *serv) GetUserById(ctx context.Context, id int) (*models.User, error) {
	user, err := s.tasksRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *serv) GetTaskById(ctx context.Context, id int) (*models.Tasks, error) {
	task, err := s.tasksRepository.GetTaskById(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
