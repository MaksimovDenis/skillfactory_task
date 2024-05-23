package tasks

import (
	"context"
	"fmt"
	"strings"

	db "skillfactory_task/internal/client"
	"skillfactory_task/internal/models"
	"skillfactory_task/internal/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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
func (r *repo) CreateTask(ctx context.Context, task *models.CreateTask) (int, error) {
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
		return nil, status.Errorf(codes.Internal, "Internal server error")
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
		Name:     "task_repository.GetTaskByAuthorName",
		QueryRaw: query,
	}

	var tasks models.Tasks
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&tasks.Id, &tasks.Opened, &tasks.Closed, &tasks.AuthorId, &tasks.AssignedId, &tasks.Title, &tasks.Content)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
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
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	fmt.Println(tasks)
	return &tasks, nil
}

// Возвращает автора по ID
func (r *repo) GetUserById(ctx context.Context, id int) (*models.User, error) {
	builder := sq.Select(idUserColumn, nameUserColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableUsers).
		Where(sq.Eq{idUserColumn: id})

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "task_repository.GetUserById",
		QueryRaw: query,
	}

	var user models.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	fmt.Println(user)
	return &user, nil
}

// Возвращает задачу по ID
func (r *repo) GetTaskById(ctx context.Context, id int) (*models.Tasks, error) {
	builder := sq.Select(idTasksColumn, openedTasksColumn, closedTasksColumn, authorIdTasksColumns, assignedIdTasksColumn, titleTasksColumn, contentTasksColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableTasks).
		Where(sq.Eq{idTasksColumn: id})

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "task_repository.GetTaskById",
		QueryRaw: query,
	}

	var task models.Tasks
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&task.Id, &task.Opened, &task.Closed, &task.AuthorId, &task.AssignedId, &task.Title, &task.Content)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	fmt.Println(task)
	return &task, nil
}

// Обнавляет задачу по ID
func (r *repo) UpdateTaskById(ctx context.Context, task *models.UpdateTask) error {
	builder := sq.Update("tasks").
		PlaceholderFormat(sq.Dollar).
		Set("author_id", task.AuthorId).
		Set("assigned_id", task.AssignedId).
		Set("title", task.Title).
		Set("content", task.Content).
		Where(sq.Eq{"id": task.Id})

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "task_repository.UpdateTaskById",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal server error")
	}

	rowAffcted := res.RowsAffected()

	if rowAffcted == 0 {
		return errors.New("Task not found")
	}

	return nil
}

// Удаляет задачу по ID
func (r *repo) DeleteTaskById(ctx context.Context, id int) error {
	builder := sq.Delete("tasks").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "task_repository.DeleteTaskById",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal server error")
	}

	rowAffcted := res.RowsAffected()

	if rowAffcted == 0 {
		return errors.New("Task not found")
	}

	return nil
}

// Удаляет задачу по ID из таблицы tasks_labels
func (r *repo) DeleteTaskLabelById(ctx context.Context, id int) error {
	builder := sq.Delete("tasks_labels").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"task_id": id})

	query, args, err := builder.ToSql()
	fmt.Println(query, args)
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "task_repository.DeleteTaskLabelById",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal server error")
	}

	rowAffcted := res.RowsAffected()

	if rowAffcted == 0 {
		return nil
	}

	return nil
}
