package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Handler - функция, которая выполняется в транзакции
type Handler func(ctx context.Context) error

// Client для работы с БД
type Client interface {
	DB() DB
	Close() error
}

// TxManager менеджер транзакций, который выполняет указанный пользователем обработчик в транзакции
type TxManger interface {
	ReadCommited(ctx context.Context, f Handler) error
}

// Transactor интерфейс для работы с транзакциями
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Pinger интерфейс для проверки соединения с БД
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB интерфейс для работы с БД
type DB interface {
	Transactor
	Pinger
	Close()
}
