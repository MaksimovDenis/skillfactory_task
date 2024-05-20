package tasks

import (
	"context"
	"fmt"
	"strings"

	db "Skillfactory_task_30.8.1/internal/client"
	"Skillfactory_task_30.8.1/internal/models"
	"Skillfactory_task_30.8.1/internal/repository"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableTasks       = "tasks"
	tableUsers       = "users"
	tableLabels      = "labels"
	tableTasksLabels = "tasks_labels"

	idColumn         = "t.id"
	openedColumn     = "t.opened"
	closedColumn     = "t.closed"
	authorIdColumns  = "t.author_id"
	assignedIdColumn = "t.assigned_id"
	titleColumn      = "t.title"
	contentColumn    = "t.content"

	idUserColumn   = "u.id"
	nameUserColumn = "u.name"

	idLableColumn   = "l.id"
	nameLabelColumn = "l.name"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TasksRepository {
	return &repo{db: db}
}

// Возвращает список задач по имени автора
func (r *repo) GetTaskByAuthorName(ctx context.Context, authorName string) (*models.Tasks, error) {
	builder := sq.Select(idColumn, openedColumn, closedColumn, authorIdColumns, assignedIdColumn, titleColumn, contentColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasks + " AS t").Join(tableUsers + " AS u ON " + idUserColumn + "=" + authorIdColumns).
		Where(sq.Eq{"LOWER(" + nameUserColumn + ")": strings.ToLower(authorName)}).
		Limit(1)

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
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

// Возвращает список всех задач по лейблу
/*func (r *repo) GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error) {
	builder := sq.Select(idColumn, openedColumn, closedColumn, authorIdColumns, assignedIdColumn, titleColumn, contentColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasksLabels + " AS tl").Join(tableTasks + " AS t ON " + idColumn + "=tl.task_id").
		Join(tableLabels + " AS l ON " + idLableColumn + "=tl.label_id").
		Where(sq.Eq{nameLabelColumn: strings.ToLower(label)})

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "task_repository.GetTaskByLabel",
		QueryRaw: query,
	}

	var tasks []models.Tasks
	err = r.db.DB().ScanAllContext(ctx, &tasks, q, args...)
	if err != nil {
		return nil, err
	}

	fmt.Println(tasks)
	return &tasks, nil
}*/

// Возвращает список всех задач по лейблу
func (r *repo) GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error) {
	builder := sq.Select(idColumn, openedColumn, closedColumn, authorIdColumns, assignedIdColumn, titleColumn, contentColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasks + " AS t").Join(tableUsers + " AS u ON " + idUserColumn + "=" + authorIdColumns).
		Where(sq.Eq{"LOWER(" + nameUserColumn + ")": strings.ToLower(label)}).
		Limit(1)

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "task_repository.Get",
		QueryRaw: query,
	}
	var tasks []models.Tasks
	err = r.db.DB().ScanAllContext(ctx, &tasks, q, args...)
	if err != nil {
		return nil, err
	}

	fmt.Println(tasks)
	return &tasks, nil
}
