package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/mdqni/dmark-todo/internal/domain"
)

type PostgresTaskRepo struct {
	DB *sql.DB
}

func NewPostgresTaskRepo(db *sql.DB) (*PostgresTaskRepo, error) {
	repo := &PostgresTaskRepo{DB: db}
	if err := repo.initSchema(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *PostgresTaskRepo) initSchema(ctx context.Context) error {
	query := `
CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	completed BOOLEAN DEFAULT FALSE,
	priority VARCHAR(10) DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high')),
	due_date TIMESTAMP,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);`

	_, err := r.DB.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}
	return nil
}

func (r *PostgresTaskRepo) AddTask(ctx context.Context, task domain.Task) error {
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO tasks (title, description, completed, priority, due_date) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.Description, task.Completed, task.Priority, task.DueDate,
	)

	return err
}

func (r *PostgresTaskRepo) GetTaskByID(ctx context.Context, id int) (domain.Task, error) {
	var task domain.Task
	query := `SELECT id, title, completed, priority, due_date, description, created_at, updated_at FROM tasks WHERE id=$1`
	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&task.ID, &task.Title, &task.Completed, &task.Priority, &task.DueDate, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, fmt.Errorf("task with id %d not found", id)
		}
		return task, err
	}
	return task, nil
}

func (r *PostgresTaskRepo) DeleteTask(ctx context.Context, id int) error {
	res, err := r.DB.ExecContext(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("no task found to delete")
	}
	return nil
}

func (r *PostgresTaskRepo) UpdateTask(ctx context.Context, task domain.Task) error {
	res, err := r.DB.ExecContext(ctx,
		"UPDATE tasks SET title=$1, completed=$2, priority=$3, due_date=$4, description=$5, updated_at=NOW() WHERE id=$6",
		task.Title, task.Completed, task.Priority, task.DueDate, task.Description, task.ID,
	)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("no task found to update")
	}
	return nil
}

func (r *PostgresTaskRepo) GetTasks(ctx context.Context) ([]domain.Task, error) {
	rows, err := r.DB.QueryContext(ctx,
		"SELECT id, title, completed, priority, due_date, description, created_at, updated_at FROM tasks ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.Priority, &t.DueDate, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *PostgresTaskRepo) GetTasksFiltered(ctx context.Context, status string, dateFilter string) ([]domain.Task, error) {
	query := "SELECT id, title, completed, priority, due_date, description, created_at, updated_at FROM tasks"
	conditions := []string{}
	args := []interface{}{}

	if status == "active" {
		conditions = append(conditions, "completed = FALSE")
	} else if status == "done" {
		conditions = append(conditions, "completed = TRUE")
	}

	switch dateFilter {
	case "today":
		conditions = append(conditions, "due_date::date = CURRENT_DATE")
	case "week":
		conditions = append(conditions, "due_date::date BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '7 days'")
	case "overdue":
		conditions = append(conditions, "due_date::date < CURRENT_DATE")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY created_at ASC"

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.Priority, &t.DueDate, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}
