package repository

import (
	"context"

	"github.com/mdqni/dmark-todo/internal/domain"
)

type TaskRepo interface {
	AddTask(ctx context.Context, task domain.Task) error
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, task domain.Task) error
	GetTaskByID(ctx context.Context, id int) (domain.Task, error)
	GetTasks(ctx context.Context) ([]domain.Task, error)

	GetTasksFiltered(ctx context.Context, status string, dateFilter string) ([]domain.Task, error)
}
