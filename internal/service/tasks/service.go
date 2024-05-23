package tasks

import (
	db "skillfactory_task/internal/client"
	"skillfactory_task/internal/repository"
	"skillfactory_task/internal/service"
)

type serv struct {
	tasksRepository repository.TasksRepository
	txManger        db.TxManger
}

func NewService(
	tasksRepository repository.TasksRepository,
	txManager db.TxManger,
) service.TasksService {
	return &serv{
		tasksRepository: tasksRepository,
		txManger:        txManager,
	}
}
