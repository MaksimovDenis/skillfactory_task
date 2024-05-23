package app

import (
	"context"
	"log"

	db "skillfactory_task/internal/client"
	"skillfactory_task/internal/client/db/pg"
	"skillfactory_task/internal/client/db/transaction"
	"skillfactory_task/internal/config"
	"skillfactory_task/internal/handler"
	"skillfactory_task/internal/repository"
	tasksRepository "skillfactory_task/internal/repository/tasks"
	"skillfactory_task/internal/service"
	tasksService "skillfactory_task/internal/service/tasks"
)

type serviceProvider struct {
	pgConfig     config.PGConfig
	serverConfig config.ServerConfig

	dbClient        db.Client
	txManager       db.TxManger
	tasksRepository repository.TasksRepository

	tasksService service.TasksService

	handler *handler.Handler
}

func newServicProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) ServerConfig() config.ServerConfig {
	if s.serverConfig == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			log.Fatalf("failed to get server config: %s", err.Error())
		}

		s.serverConfig = cfg
	}

	return s.serverConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManger {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionsManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) TasksRepository(ctx context.Context) repository.TasksRepository {
	if s.tasksRepository == nil {
		s.tasksRepository = tasksRepository.NewRepository(s.DBClient(ctx))
	}

	return s.tasksRepository
}

func (s *serviceProvider) TasksService(ctx context.Context) service.TasksService {
	if s.tasksService == nil {
		s.tasksService = tasksService.NewService(
			s.TasksRepository(ctx),
			s.TxManager(ctx),
		)
	}
	return s.tasksService
}

func (s *serviceProvider) TasksHandler(ctx context.Context) *handler.Handler {
	if s.handler == nil {
		s.handler = handler.NewHandler(s.TasksService(ctx))
	}
	return s.handler
}
