package service

import (
	"github.com/mdqni/dmark-todo/internal/repository"
)

type Services struct {
	Tasks *TaskService
}

func NewServices(repos repository.TaskRepo) *Services {
	return &Services{
		Tasks: NewTaskService(repos),
	}
}
