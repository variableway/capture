package store

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/variableway/innate/capture/internal/model"
)

func TestMarkdownStore_CreateAndGet(t *testing.T) {
	dir := t.TempDir()
	s := NewMarkdownStore(dir)

	task := &model.Task{
		ID:        "TASK-00001",
		Title:     "Test task",
		Status:    model.StatusTodo,
		Priority:  model.PriorityHigh,
		Tags:      []string{"test"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Source:    "cli",
	}

	ctx := context.Background()
	if err := s.CreateTask(ctx, task); err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	// Verify file was created
	path := s.taskPath(task)
	if _, err := filepath.Glob(path); err != nil {
		t.Errorf("task file not found at %s", path)
	}

	got, err := s.GetTask(ctx, "TASK-00001")
	if err != nil {
		t.Fatalf("GetTask failed: %v", err)
	}

	if got.ID != task.ID {
		t.Errorf("ID = %q, want %q", got.ID, task.ID)
	}
	if got.Title != task.Title {
		t.Errorf("Title = %q, want %q", got.Title, task.Title)
	}
}

func TestMarkdownStore_Update(t *testing.T) {
	dir := t.TempDir()
	s := NewMarkdownStore(dir)

	task := &model.Task{
		ID:        "TASK-00001",
		Title:     "Original",
		Status:    model.StatusTodo,
		Priority:  model.PriorityMedium,
		Tags:      []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Source:    "cli",
	}

	ctx := context.Background()
	s.CreateTask(ctx, task)

	task.Title = "Updated"
	task.Status = model.StatusInProgress
	task.UpdatedAt = time.Now()

	if err := s.UpdateTask(ctx, task); err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	got, _ := s.GetTask(ctx, "TASK-00001")
	if got.Title != "Updated" {
		t.Errorf("Title = %q, want %q", got.Title, "Updated")
	}
}

func TestMarkdownStore_Delete(t *testing.T) {
	dir := t.TempDir()
	s := NewMarkdownStore(dir)

	task := &model.Task{
		ID:        "TASK-00001",
		Title:     "To delete",
		Status:    model.StatusTodo,
		Priority:  model.PriorityMedium,
		Tags:      []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Source:    "cli",
	}

	ctx := context.Background()
	s.CreateTask(ctx, task)

	if err := s.DeleteTask(ctx, "TASK-00001"); err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}

	_, err := s.GetTask(ctx, "TASK-00001")
	if err == nil {
		t.Error("expected error after delete, got nil")
	}
}

func TestMarkdownStore_List(t *testing.T) {
	dir := t.TempDir()
	s := NewMarkdownStore(dir)

	ctx := context.Background()
	for i, title := range []string{"Task A", "Task B", "Task C"} {
		task := &model.Task{
			ID:        fmt.Sprintf("TASK-0000%d", i+1),
			Title:     title,
			Status:    model.StatusTodo,
			Priority:  model.PriorityMedium,
			Tags:      []string{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Source:    "cli",
		}
		s.CreateTask(ctx, task)
	}

	tasks, err := s.ListTasks(ctx, model.TaskFilter{})
	if err != nil {
		t.Fatalf("ListTasks failed: %v", err)
	}

	if len(tasks) != 3 {
		t.Errorf("ListTasks returned %d tasks, want 3", len(tasks))
	}
}
