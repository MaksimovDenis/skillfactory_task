package tasks

import (
	db "Skillfactory_task_30.8.1/internal/client"
	"Skillfactory_task_30.8.1/internal/repository"
	"Skillfactory_task_30.8.1/internal/service"
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
