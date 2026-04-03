package model

import (
	"testing"
)

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		status  string
		want    bool
	}{
		{"todo", true},
		{"in_progress", true},
		{"done", true},
		{"cancelled", true},
		{"archived", true},
		{"unknown", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			if got := IsValidStatus(tt.status); got != tt.want {
				t.Errorf("IsValidStatus(%q) = %v, want %v", tt.status, got, tt.want)
			}
		})
	}
}

func TestIsValidPriority(t *testing.T) {
	tests := []struct {
		priority string
		want     bool
	}{
		{"high", true},
		{"medium", true},
		{"low", true},
		{"urgent", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.priority, func(t *testing.T) {
			if got := IsValidPriority(tt.priority); got != tt.want {
				t.Errorf("IsValidPriority(%q) = %v, want %v", tt.priority, got, tt.want)
			}
		})
	}
}

func TestCanTransition(t *testing.T) {
	tests := []struct {
		name    string
		from    TaskStatus
		to      TaskStatus
		want    bool
	}{
		{"todo to in_progress", StatusTodo, StatusInProgress, true},
		{"todo to done", StatusTodo, StatusDone, true},
		{"todo to cancelled", StatusTodo, StatusCancelled, true},
		{"todo to archived", StatusTodo, StatusArchived, false},
		{"in_progress to done", StatusInProgress, StatusDone, true},
		{"in_progress to todo", StatusInProgress, StatusTodo, true},
		{"done to archived", StatusDone, StatusArchived, true},
		{"done to todo", StatusDone, StatusTodo, false},
		{"archived to todo", StatusArchived, StatusTodo, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanTransition(tt.from, tt.to); got != tt.want {
				t.Errorf("CanTransition(%s, %s) = %v, want %v", tt.from, tt.to, got, tt.want)
			}
		})
	}
}

func TestNewTask(t *testing.T) {
	task := NewTask("test task")
	if task.Title != "test task" {
		t.Errorf("NewTask title = %q, want %q", task.Title, "test task")
	}
	if task.Status != StatusTodo {
		t.Errorf("NewTask status = %q, want %q", task.Status, StatusTodo)
	}
	if task.Priority != PriorityMedium {
		t.Errorf("NewTask priority = %q, want %q", task.Priority, PriorityMedium)
	}
	if task.Source != "cli" {
		t.Errorf("NewTask source = %q, want %q", task.Source, "cli")
	}
}
