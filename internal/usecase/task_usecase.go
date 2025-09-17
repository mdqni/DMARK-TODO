package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/mdqni/dmark-todo/internal/domain"
	"github.com/mdqni/dmark-todo/internal/dto"
	"github.com/mdqni/dmark-todo/internal/service"
)

type TaskUseCase struct {
	Service *service.TaskService
}

func NewTaskUseCase(service *service.TaskService) *TaskUseCase {
	return &TaskUseCase{Service: service}
}

func (uc *TaskUseCase) AddTask(title, description, priority, dueDate string) error {
	ctx := context.Background()

	var parsedDate *time.Time
	if dueDate != "" {
		t, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return err
		}
		parsedDate = &t
	}

	task := domain.Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		DueDate:     parsedDate,
		Completed:   false,
	}

	return uc.Service.AddTask(ctx, task)
}

func (uc *TaskUseCase) DeleteTask(idStr string) error {
	ctx := context.Background()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	return uc.Service.DeleteTask(ctx, id)
}

func (uc *TaskUseCase) ToggleDone(idStr string) error {
	ctx := context.Background()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	return uc.Service.ToggleDone(ctx, id)
}
func (u *TaskUseCase) ListTasks(status string, dateFilter string) ([]dto.TaskDTO, error) {
	tasks, err := u.Service.ListTasksFiltered(context.Background(), status, dateFilter)
	if err != nil {
		return nil, err
	}

	var dtos []dto.TaskDTO
	for _, t := range tasks {
		dtos = append(dtos, dto.ToDTO(t))
	}

	return dtos, nil
}
