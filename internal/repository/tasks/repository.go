package tasks

import (
	"context"

	db "Skillfactory_task_30.8.1/internal/client"
	"Skillfactory_task_30.8.1/internal/models"
	"Skillfactory_task_30.8.1/internal/repository"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "tasks"

	idColumn         = "id"
	openedColumn     = "opened"
	closedColumn     = "closed"
	authorIdColumns  = "author_id"
	assignedIdColumn = "assigned_id"
	titleColumn      = "title"
	contentColumn    = "content"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TasksRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, userID int) (*models.Tasks, error) {
	builder := sq.Select(idColumn, openedColumn, closedColumn, authorIdColumns, assignedIdColumn, titleColumn, contentColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{authorIdColumns: userID}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "task_repository.Get",
		QueryRaw: query,
	}

	var tasks models.Tasks
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&tasks.Id, &tasks.Opened, &tasks.Closed, &tasks.AuthorId, &tasks.AssignedId, &tasks.Title, &tasks.Content)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}
