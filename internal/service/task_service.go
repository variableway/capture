package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/store"
	"github.com/variableway/innate/capture/pkg/idgen"
)

type TaskService struct {
	store   store.Store
	dataDir string
}

func NewTaskService(store store.Store, dataDir string) *TaskService {
	return &TaskService{store: store, dataDir: dataDir}
}

func (s *TaskService) Create(ctx context.Context, title string, opts ...TaskOption) (*model.Task, error) {
	if strings.TrimSpace(title) == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	task := model.NewTask(title)
	for _, opt := range opts {
		opt(task)
	}

	id, err := idgen.Next(s.dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ID: %w", err)
	}
	task.ID = id

	if err := s.store.CreateTask(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *TaskService) Get(ctx context.Context, id string) (*model.Task, error) {
	return s.store.GetTask(ctx, id)
}

func (s *TaskService) Update(ctx context.Context, id string, opts ...TaskOption) (*model.Task, error) {
	task, err := s.store.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(task)
	}
	task.UpdatedAt = time.Now()

	if err := s.store.UpdateTask(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	if _, err := s.store.GetTask(ctx, id); err != nil {
		return err
	}
	return s.store.DeleteTask(ctx, id)
}

func (s *TaskService) List(ctx context.Context, filter model.TaskFilter) ([]*model.Task, error) {
	return s.store.ListTasks(ctx, filter)
}

func (s *TaskService) SetStatus(ctx context.Context, id string, status model.TaskStatus) (*model.Task, error) {
	task, err := s.store.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	if !model.CanTransition(task.Status, status) {
		return nil, fmt.Errorf("cannot transition from %s to %s", task.Status, status)
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	if err := s.store.UpdateTask(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to update status: %w", err)
	}

	return task, nil
}

// TaskOption is a functional option for modifying a task.
type TaskOption func(*model.Task)

func WithDescription(desc string) TaskOption {
	return func(t *model.Task) { t.Description = desc }
}

func WithPriority(p model.TaskPriority) TaskOption {
	return func(t *model.Task) { t.Priority = p }
}

func WithTags(tags []string) TaskOption {
	return func(t *model.Task) { t.Tags = tags }
}

func WithSource(source string) TaskOption {
	return func(t *model.Task) { t.Source = source }
}

func WithContext(ctx model.TaskContext) TaskOption {
	return func(t *model.Task) { t.Context = ctx }
}
