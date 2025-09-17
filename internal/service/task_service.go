package service

import (
	"context"

	"github.com/mdqni/dmark-todo/internal/domain"
	"github.com/mdqni/dmark-todo/internal/repository"
)

type TaskService struct {
	Repo repository.TaskRepo
}

func NewTaskService(repo repository.TaskRepo) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) AddTask(ctx context.Context, task domain.Task) error {
	return s.Repo.AddTask(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.Repo.DeleteTask(ctx, id)
}

func (s *TaskService) ToggleDone(ctx context.Context, id int) error {
	task, err := s.Repo.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}
	task.Completed = !task.Completed
	return s.Repo.UpdateTask(ctx, task)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id int) (domain.Task, error) {
	return s.Repo.GetTaskByID(ctx, id)
}

func (s *TaskService) ListTasksFiltered(ctx context.Context, status, dateFilter, sortBy string) ([]domain.Task, error) {
	return s.Repo.GetTasksFiltered(ctx, status, dateFilter, sortBy)
}
