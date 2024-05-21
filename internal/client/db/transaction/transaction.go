package transaction

import (
	"context"

	db "Skillfactory_task_30.8.1/internal/client"
	"Skillfactory_task_30.8.1/internal/client/db/pg"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type manager struct {
	db db.Transactor
}

// Менеджер транзакций, который удовлетворяет интерфейск db.TxManager
func NewTransactionsManager(db db.Transactor) db.TxManger {
	return &manager{
		db: db,
	}
}

// transaction основная функция, которая выполняет указанный пользователем обработчик в транзакции
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	// Если это вложенная транзакция, пропускаем инициацию новой транзакции и выполняем обработчик.
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	// Стратуем новоую транзакцию
	tx, err = m.db.BeginTx(ctx, opts)

	// Настраиваем функцию отсрочки для отката или коммита транзакции
	defer func() {
		// Восстановление после паники
		if r := recover(); r != nil {
			err = errors.Errorf("panica recovered: %v", r)
		}

		// Откатываем транзакцию, если произошла ошибка
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}
			return
		}

		// Ошибок не было, коммитим транзакцию
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrapf(err, "tx commit failed")
			}
		}
	}()

	// Выполните код внутри транзакциию
	// Если функция терпит неудачу, вощвращаем ошибку, и функция отсрочки выполняет откат
	// или в противном случае транзакция коммитится
	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
