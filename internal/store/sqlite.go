package store

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"

	"github.com/variableway/innate/capture/internal/model"
)

// SQLiteStore provides an index layer over tasks for fast querying.
type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dataDir string) (*SQLiteStore, error) {
	dbPath := filepath.Join(dataDir, "capture.db")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	s := &SQLiteStore{db: db}
	if err := s.migrate(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return s, nil
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}

func (s *SQLiteStore) migrate() error {
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'todo',
		priority TEXT DEFAULT 'medium',
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		source TEXT DEFAULT 'cli',
		file_path TEXT NOT NULL,
		feishu_record_id TEXT DEFAULT ''
	);
	CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
	CREATE INDEX IF NOT EXISTS idx_tasks_priority ON tasks(priority);
	CREATE INDEX IF NOT EXISTS idx_tasks_created_at ON tasks(created_at);

	CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS task_tags (
		task_id TEXT NOT NULL,
		tag TEXT NOT NULL,
		PRIMARY KEY (task_id, tag),
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS sync_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sync_type TEXT NOT NULL,
		direction TEXT NOT NULL,
		started_at DATETIME NOT NULL,
		completed_at DATETIME,
		status TEXT,
		records_count INTEGER DEFAULT 0,
		error_message TEXT
	);
	`
	_, err := s.db.Exec(schema)
	return err
}

func (s *SQLiteStore) CreateTask(ctx context.Context, task *model.Task) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`INSERT INTO tasks (id, title, status, priority, created_at, updated_at, source, file_path, feishu_record_id)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		task.ID, task.Title, string(task.Status), string(task.Priority),
		task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		task.Source, task.FilePath, task.Sync.FeishuRecordID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}

	for _, tag := range task.Tags {
		_, _ = tx.ExecContext(ctx, `INSERT OR IGNORE INTO tags (name) VALUES (?)`, tag)
		_, err = tx.ExecContext(ctx, `INSERT INTO task_tags (task_id, tag) VALUES (?, ?)`, task.ID, tag)
		if err != nil {
			return fmt.Errorf("failed to insert tag: %w", err)
		}
	}

	return tx.Commit()
}

func (s *SQLiteStore) GetTask(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task
	var status, priority, source, filePath, feishuRecordID string
	var createdAt, updatedAt string

	err := s.db.QueryRowContext(ctx,
		`SELECT id, title, status, priority, created_at, updated_at, source, file_path, feishu_record_id
		 FROM tasks WHERE id = ?`, id,
	).Scan(&task.ID, &task.Title, &status, &priority, &createdAt, &updatedAt, &source, &filePath, &feishuRecordID)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task %s not found", id)
	}
	if err != nil {
		return nil, err
	}

	task.Status = model.TaskStatus(status)
	task.Priority = model.TaskPriority(priority)
	task.Source = source
	task.FilePath = filePath
	task.Sync.FeishuRecordID = feishuRecordID

	// Parse timestamps
	task.CreatedAt, _ = parseTime(createdAt)
	task.UpdatedAt, _ = parseTime(updatedAt)

	// Load tags
	tags, err := s.getTags(ctx, id)
	if err == nil {
		task.Tags = tags
	}

	return &task, nil
}

func (s *SQLiteStore) UpdateTask(ctx context.Context, task *model.Task) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`UPDATE tasks SET title=?, status=?, priority=?, updated_at=?, file_path=?, feishu_record_id=? WHERE id=?`,
		task.Title, string(task.Status), string(task.Priority),
		task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		task.FilePath, task.Sync.FeishuRecordID, task.ID,
	)
	if err != nil {
		return err
	}

	// Re-sync tags
	_, _ = tx.ExecContext(ctx, `DELETE FROM task_tags WHERE task_id = ?`, task.ID)
	for _, tag := range task.Tags {
		_, _ = tx.ExecContext(ctx, `INSERT OR IGNORE INTO tags (name) VALUES (?)`, tag)
		_, _ = tx.ExecContext(ctx, `INSERT INTO task_tags (task_id, tag) VALUES (?, ?)`, task.ID, tag)
	}

	return tx.Commit()
}

func (s *SQLiteStore) DeleteTask(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM task_tags WHERE task_id = ?`, id)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `DELETE FROM tasks WHERE id = ?`, id)
	return err
}

func (s *SQLiteStore) ListTasks(ctx context.Context, filter model.TaskFilter) ([]*model.Task, error) {
	query := `SELECT id, title, status, priority, created_at, updated_at, source, file_path, feishu_record_id FROM tasks`
	var conditions []string
	var args []interface{}

	if filter.Status != nil {
		conditions = append(conditions, "status = ?")
		args = append(args, string(*filter.Status))
	}
	if filter.Priority != nil {
		conditions = append(conditions, "priority = ?")
		args = append(args, string(*filter.Priority))
	}
	if filter.Source != nil {
		conditions = append(conditions, "source = ?")
		args = append(args, string(*filter.Source))
	}
	if len(filter.Tags) > 0 {
		conditions = append(conditions, fmt.Sprintf(
			"id IN (SELECT task_id FROM task_tags WHERE tag IN (%s) GROUP BY task_id HAVING COUNT(DISTINCT tag) = %d)",
			strings.Repeat("?,", len(filter.Tags))[0:len(filter.Tags)*2-1],
			len(filter.Tags),
		))
		for _, t := range filter.Tags {
			args = append(args, t)
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY created_at DESC"

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var task model.Task
		var status, priority, source, filePath, feishuRecordID string
		var createdAt, updatedAt string

		if err := rows.Scan(&task.ID, &task.Title, &status, &priority, &createdAt, &updatedAt, &source, &filePath, &feishuRecordID); err != nil {
			continue
		}

		task.Status = model.TaskStatus(status)
		task.Priority = model.TaskPriority(priority)
		task.Source = source
		task.FilePath = filePath
		task.Sync.FeishuRecordID = feishuRecordID
		task.CreatedAt, _ = parseTime(createdAt)
		task.UpdatedAt, _ = parseTime(updatedAt)

		tags, _ := s.getTags(ctx, task.ID)
		task.Tags = tags

		tasks = append(tasks, &task)
	}

	return tasks, rows.Err()
}

func (s *SQLiteStore) getTags(ctx context.Context, taskID string) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT tag FROM task_tags WHERE task_id = ?`, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func parseTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	formats := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse time: %s", s)
}
