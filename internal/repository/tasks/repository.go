package tasks

import (
	"context"
	"fmt"
	"strings"

	db "Skillfactory_task_30.8.1/internal/client"
	"Skillfactory_task_30.8.1/internal/models"
	"Skillfactory_task_30.8.1/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	tableTasks       = "tasks AS t"
	tableUsers       = "users AS u"
	tableLabels      = "labels AS l"
	tableTasksLabels = "tasks_labels AS tl"

	idTasksColumn         = "t.id"
	openedTasksColumn     = "t.opened"
	closedTasksColumn     = "t.closed"
	authorIdTasksColumns  = "t.author_id"
	assignedIdTasksColumn = "t.assigned_id"
	titleTasksColumn      = "t.title"
	contentTasksColumn    = "t.content"

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

// Возвращает список всех задач
func (r *repo) CreateTask(ctx context.Context, task *models.CreateTasks) (int, error) {
	builder := sq.Insert("tasks").
		PlaceholderFormat(sq.Dollar).
		Columns("author_id", "assigned_id", "title", "content").
		Values(task.AuthorId, task.AssignedId, task.Title, task.Content).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "task_repository.CreateTask",
		QueryRaw: query,
	}

	var taskID int

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&taskID)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "Internal server error")
	}

	return taskID, nil
}

// Возвращает список всех задач
func (r *repo) GetTasks(ctx context.Context) (*[]models.Tasks, error) {
	builder := sq.Select(idTasksColumn, openedTasksColumn, closedTasksColumn, authorIdTasksColumns, assignedIdTasksColumn, titleTasksColumn, contentTasksColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasks).
		OrderBy(idTasksColumn)

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
}

// Возвращает список задач по имени автора
func (r *repo) GetTaskByAuthorName(ctx context.Context, authorName string) (*models.Tasks, error) {
	builder := sq.Select(idTasksColumn, openedTasksColumn, closedTasksColumn, authorIdTasksColumns, assignedIdTasksColumn, titleTasksColumn, contentTasksColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasks).Join(tableUsers + " ON " + idUserColumn + "=" + authorIdTasksColumns).
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
func (r *repo) GetTaskByLabel(ctx context.Context, label string) (*[]models.Tasks, error) {
	builder := sq.Select(idTasksColumn, openedTasksColumn, closedTasksColumn, authorIdTasksColumns, assignedIdTasksColumn, titleTasksColumn, contentTasksColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasksLabels).Join(tableTasks + " ON " + idTasksColumn + "=tl.task_id").
		Join(tableLabels + " ON " + idLableColumn + "=tl.label_id").
		Where(sq.Eq{"LOWER(" + nameLabelColumn + ")": strings.ToLower(label)})

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
}
