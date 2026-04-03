package store

import (
	"context"

	"github.com/variableway/innate/capture/internal/model"
)

// Store defines the interface for task persistence operations.
type Store interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTask(ctx context.Context, id string) (*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context, filter model.TaskFilter) ([]*model.Task, error)
}
