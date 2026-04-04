package store

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/variableway/innate/capture/internal/model"
)

// MarkdownStore implements file-based task storage using Markdown + YAML frontmatter.
type MarkdownStore struct {
	dataDir string
}

func NewMarkdownStore(dataDir string) *MarkdownStore {
	return &MarkdownStore{dataDir: dataDir}
}

func (s *MarkdownStore) taskDir() string {
	return filepath.Join(s.dataDir, "tasks")
}

func (s *MarkdownStore) taskPath(task *model.Task) string {
	return filepath.Join(
		s.taskDir(),
		task.CreatedAt.Format("2006"),
		task.CreatedAt.Format("01"),
		task.ID+".md",
	)
}

func (s *MarkdownStore) ensureDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0755)
}

func (s *MarkdownStore) CreateTask(ctx context.Context, task *model.Task) error {
	path := s.taskPath(task)
	if err := s.ensureDir(path); err != nil {
		return fmt.Errorf("failed to create task directory: %w", err)
	}

	data, err := s.encode(task)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write task file: %w", err)
	}

	task.FilePath = path
	return nil
}

func (s *MarkdownStore) GetTask(ctx context.Context, id string) (*model.Task, error) {
	path, err := s.findTaskFile(id)
	if err != nil {
		return nil, err
	}

	return s.decodeFile(path)
}

func (s *MarkdownStore) UpdateTask(ctx context.Context, task *model.Task) error {
	path := task.FilePath
	if path == "" {
		var err error
		path, err = s.findTaskFile(task.ID)
		if err != nil {
			return err
		}
	}

	data, err := s.encode(task)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (s *MarkdownStore) DeleteTask(ctx context.Context, id string) error {
	path, err := s.findTaskFile(id)
	if err != nil {
		return err
	}
	return os.Remove(path)
}

func (s *MarkdownStore) ListTasks(ctx context.Context, filter model.TaskFilter) ([]*model.Task, error) {
	tasksDir := s.taskDir()
	var tasks []*model.Task

	err := filepath.Walk(tasksDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		task, err := s.decodeFile(path)
		if err != nil {
			return nil // skip malformed files
		}

		if matchesFilter(task, &filter) {
			tasks = append(tasks, task)
		}
		return nil
	})

	return tasks, err
}

func (s *MarkdownStore) findTaskFile(id string) (string, error) {
	tasksDir := s.taskDir()
	var found string

	filepath.Walk(tasksDir, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil || info.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		if strings.TrimSuffix(base, ".md") == id {
			found = path
			return fmt.Errorf("found")
		}
		return nil
	})

	if found == "" {
		return "", fmt.Errorf("task %s not found", id)
	}

	return found, nil
}

func (s *MarkdownStore) encode(task *model.Task) ([]byte, error) {
	fmBytes, err := yaml.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal task: %w", err)
	}

	var sb strings.Builder
	sb.WriteString("---\n")
	sb.Write(fmBytes)
	sb.WriteString("---\n")

	if task.Body != "" {
		sb.WriteString("\n" + task.Body + "\n")
	} else {
		sb.WriteString("\n## Description\n\n" + task.Description + "\n")
		if len(task.Tags) > 0 {
			sb.WriteString("\n## Notes\n\n")
		}
	}

	return []byte(sb.String()), nil
}

func (s *MarkdownStore) decodeFile(path string) (*model.Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read task file: %w", err)
	}

	content := string(data)

	// Split by --- delimiters
	if !strings.HasPrefix(strings.TrimSpace(content), "---") {
		return nil, fmt.Errorf("invalid task file format: %s", path)
	}

	parts := strings.SplitN(strings.TrimPrefix(content, "---"), "---", 2)
	fmStr := strings.TrimSpace(parts[0])

	var task model.Task
	if err := yaml.Unmarshal([]byte(fmStr), &task); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	if len(parts) > 1 {
		task.Body = strings.TrimSpace(parts[1])
	}

	task.FilePath = path
	return &task, nil
}

func matchesFilter(task *model.Task, filter *model.TaskFilter) bool {
	if filter == nil {
		return true
	}
	if filter.Status != nil && task.Status != *filter.Status {
		return false
	}
	if filter.Stage != nil && task.Stage != *filter.Stage {
		return false
	}
	if filter.Priority != nil && task.Priority != *filter.Priority {
		return false
	}
	if filter.Source != nil && task.Source != *filter.Source {
		return false
	}
	if len(filter.Tags) > 0 {
		tagSet := make(map[string]bool)
		for _, t := range task.Tags {
			tagSet[t] = true
		}
		for _, ft := range filter.Tags {
			if !tagSet[ft] {
				return false
			}
		}
	}
	return true
}

// Ensure MarkdownStore implements Store interface at compile time.
var _ Store = (*MarkdownStore)(nil)

// Helper to get the current time (extracted for testability).
var now = time.Now
