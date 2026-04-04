package model

import "time"

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
	StatusCancelled  TaskStatus = "cancelled"
	StatusArchived   TaskStatus = "archived"
)

func ValidStatuses() []TaskStatus {
	return []TaskStatus{StatusTodo, StatusInProgress, StatusDone, StatusCancelled, StatusArchived}
}

func IsValidStatus(s string) bool {
	for _, v := range ValidStatuses() {
		if TaskStatus(s) == v {
			return true
		}
	}
	return false
}

type TaskStage string

const (
	StageInbox     TaskStage = "inbox"
	StageMindstorm TaskStage = "mindstorm"
	StageAnalysis  TaskStage = "analysis"
	StagePlanning  TaskStage = "planning"
	StagePRD       TaskStage = "prd"
	StageTasks     TaskStage = "tasks"
	StageDispatch  TaskStage = "dispatch"
	StageExecution TaskStage = "execution"
	StageReview    TaskStage = "review"
)

func ValidStages() []TaskStage {
	return []TaskStage{
		StageInbox,
		StageMindstorm,
		StageAnalysis,
		StagePlanning,
		StagePRD,
		StageTasks,
		StageDispatch,
		StageExecution,
		StageReview,
	}
}

func IsValidStage(s string) bool {
	for _, v := range ValidStages() {
		if TaskStage(s) == v {
			return true
		}
	}
	return false
}

type TaskPriority string

const (
	PriorityHigh   TaskPriority = "high"
	PriorityMedium TaskPriority = "medium"
	PriorityLow    TaskPriority = "low"
)

func ValidPriorities() []TaskPriority {
	return []TaskPriority{PriorityHigh, PriorityMedium, PriorityLow}
}

func IsValidPriority(p string) bool {
	for _, v := range ValidPriorities() {
		if TaskPriority(p) == v {
			return true
		}
	}
	return false
}

type TaskContext struct {
	Trigger   string `yaml:"trigger" json:"trigger"`
	Location  string `yaml:"location" json:"location"`
	RelatedTo string `yaml:"related_to" json:"related_to"`
}

type TaskExecution struct {
	Agent       string     `yaml:"agent" json:"agent"`
	Model       string     `yaml:"model" json:"model"`
	Result      string     `yaml:"result" json:"result"`
	Status      string     `yaml:"exec_status" json:"exec_status"`
	StartedAt   *time.Time `yaml:"started_at" json:"started_at"`
	CompletedAt *time.Time `yaml:"completed_at" json:"completed_at"`
}

type TaskDispatch struct {
	Agent           string     `yaml:"agent" json:"agent"`
	Model           string     `yaml:"model" json:"model"`
	Repository      string     `yaml:"repository" json:"repository"`
	Worktree        string     `yaml:"worktree" json:"worktree"`
	TerminalSession string     `yaml:"terminal_session" json:"terminal_session"`
	AssignedAt      *time.Time `yaml:"assigned_at" json:"assigned_at"`
}

type TaskSync struct {
	FeishuRecordID string     `yaml:"feishu_record_id" json:"feishu_record_id"`
	LastSyncedAt   *time.Time `yaml:"last_synced_at" json:"last_synced_at"`
}

type Task struct {
	ID          string        `yaml:"id" json:"id"`
	Title       string        `yaml:"title" json:"title"`
	Description string        `yaml:"description" json:"description"`
	Status      TaskStatus    `yaml:"status" json:"status"`
	Stage       TaskStage     `yaml:"stage" json:"stage"`
	Priority    TaskPriority  `yaml:"priority" json:"priority"`
	Tags        []string      `yaml:"tags" json:"tags"`
	CreatedAt   time.Time     `yaml:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `yaml:"updated_at" json:"updated_at"`
	Source      string        `yaml:"source" json:"source"` // cli, tui, feishu_bot
	Context     TaskContext   `yaml:"context" json:"context"`
	Dispatch    TaskDispatch  `yaml:"dispatch" json:"dispatch"`
	Execution   TaskExecution `yaml:"execution" json:"execution"`
	Sync        TaskSync      `yaml:"sync" json:"sync"`
	FilePath    string        `yaml:"-" json:"-"`
	Body        string        `yaml:"-" json:"-"` // Markdown body content
}

func NewTask(title string) *Task {
	now := time.Now()
	return &Task{
		Title:     title,
		Status:    StatusTodo,
		Stage:     StageInbox,
		Priority:  PriorityMedium,
		Tags:      []string{},
		CreatedAt: now,
		UpdatedAt: now,
		Source:    "cli",
	}
}

// ValidTransitions defines allowed status transitions.
var validTransitions = map[TaskStatus][]TaskStatus{
	StatusTodo:       {StatusInProgress, StatusDone, StatusCancelled},
	StatusInProgress: {StatusDone, StatusCancelled, StatusTodo},
	StatusDone:       {StatusArchived},
	StatusCancelled:  {StatusTodo, StatusArchived},
	StatusArchived:   {},
}

func CanTransition(from, to TaskStatus) bool {
	allowed, ok := validTransitions[from]
	if !ok {
		return false
	}
	for _, s := range allowed {
		if s == to {
			return true
		}
	}
	return false
}

type TaskFilter struct {
	Status   *TaskStatus
	Stage    *TaskStage
	Priority *TaskPriority
	Tags     []string
	Source   *string
}
