package dto

import (
	"time"

	"github.com/mdqni/dmark-todo/internal/domain"
)

type TaskDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"` // ← string для фронта
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToDTO(task domain.Task) TaskDTO {
	return TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Completed:   task.Completed,
		Priority:    task.Priority,
		DueDate:     formatTime(task.DueDate),
		Description: task.Description,
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}
}

func FromDTO(dto TaskDTO) domain.Task {
	var due *time.Time
	if dto.DueDate != "" {
		if t, err := time.Parse("2006-01-02", dto.DueDate); err == nil {
			due = &t
		}
	}

	created, _ := time.Parse(time.RFC3339, dto.CreatedAt)
	updated, _ := time.Parse(time.RFC3339, dto.UpdatedAt)

	return domain.Task{
		ID:          dto.ID,
		Title:       dto.Title,
		Completed:   dto.Completed,
		Priority:    dto.Priority,
		DueDate:     due,
		Description: dto.Description,
		CreatedAt:   created,
		UpdatedAt:   updated,
	}
}

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}
