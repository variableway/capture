package store

import (
	"context"
	"fmt"
	"log"

	"github.com/variableway/innate/capture/internal/model"
)

// DualStore coordinates writes to both Markdown files and SQLite index.
// Markdown is the source of truth; SQLite provides fast querying.
type DualStore struct {
	markdown *MarkdownStore
	sqlite   *SQLiteStore
}

func NewDualStore(dataDir string) (*DualStore, error) {
	md := NewMarkdownStore(dataDir)
	sq, err := NewSQLiteStore(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to init SQLite: %w", err)
	}

	return &DualStore{markdown: md, sqlite: sq}, nil
}

func (s *DualStore) Close() error {
	if s.sqlite != nil {
		return s.sqlite.Close()
	}
	return nil
}

func (s *DualStore) CreateTask(ctx context.Context, task *model.Task) error {
	// Write Markdown first (source of truth)
	if err := s.markdown.CreateTask(ctx, task); err != nil {
		return fmt.Errorf("markdown write failed: %w", err)
	}

	// Then write to SQLite index
	if err := s.sqlite.CreateTask(ctx, task); err != nil {
		log.Printf("Warning: SQLite index write failed for %s: %v", task.ID, err)
		// Markdown write succeeded, SQLite failure is non-fatal
	}

	return nil
}

func (s *DualStore) GetTask(ctx context.Context, id string) (*model.Task, error) {
	// Try SQLite first for speed, fall back to Markdown
	task, err := s.sqlite.GetTask(ctx, id)
	if err != nil {
		return s.markdown.GetTask(ctx, id)
	}

	// Enrich with full data from Markdown
	mdTask, err := s.markdown.GetTask(ctx, id)
	if err == nil {
		task.Body = mdTask.Body
		task.Description = mdTask.Description
		task.Context = mdTask.Context
		task.Execution = mdTask.Execution
		task.Sync = mdTask.Sync
	}

	return task, nil
}

func (s *DualStore) UpdateTask(ctx context.Context, task *model.Task) error {
	if err := s.markdown.UpdateTask(ctx, task); err != nil {
		return fmt.Errorf("markdown update failed: %w", err)
	}

	if err := s.sqlite.UpdateTask(ctx, task); err != nil {
		log.Printf("Warning: SQLite index update failed for %s: %v", task.ID, err)
	}

	return nil
}

func (s *DualStore) DeleteTask(ctx context.Context, id string) error {
	if err := s.markdown.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("markdown delete failed: %w", err)
	}

	if err := s.sqlite.DeleteTask(ctx, id); err != nil {
		log.Printf("Warning: SQLite index delete failed for %s: %v", id, err)
	}

	return nil
}

func (s *DualStore) ListTasks(ctx context.Context, filter model.TaskFilter) ([]*model.Task, error) {
	// Use SQLite for fast listing
	return s.sqlite.ListTasks(ctx, filter)
}

// RebuildIndex reconstructs the SQLite index from Markdown files.
func (s *DualStore) RebuildIndex(ctx context.Context) error {
	allFilter := model.TaskFilter{}
	tasks, err := s.markdown.ListTasks(ctx, allFilter)
	if err != nil {
		return fmt.Errorf("failed to read markdown tasks: %w", err)
	}

	for _, task := range tasks {
		existing, _ := s.sqlite.GetTask(ctx, task.ID)
		if existing != nil {
			_ = s.sqlite.UpdateTask(ctx, task)
		} else {
			_ = s.sqlite.CreateTask(ctx, task)
		}
	}

	return nil
}

// Ensure DualStore implements Store interface.
var _ Store = (*DualStore)(nil)
