package bitable

import (
	"github.com/variableway/innate/capture/internal/model"
)

// TaskToFields converts a Task model to Bitable record fields.
func TaskToFields(task *model.Task) map[string]interface{} {
	return map[string]interface{}{
		"task_id":    task.ID,
		"title":      task.Title,
		"status":     string(task.Status),
		"priority":   string(task.Priority),
		"source":     task.Source,
		"created_at": task.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": task.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// FieldsToTask converts Bitable record fields to a Task model.
func FieldsToTask(fields map[string]interface{}) *model.Task {
	task := &model.Task{
		Status:   model.StatusTodo,
		Priority: model.PriorityMedium,
	}

	if v, ok := fields["task_id"].(string); ok {
		task.ID = v
	}
	if v, ok := fields["title"].(string); ok {
		task.Title = v
	}
	if v, ok := fields["status"].(string); ok && model.IsValidStatus(v) {
		task.Status = model.TaskStatus(v)
	}
	if v, ok := fields["priority"].(string); ok && model.IsValidPriority(v) {
		task.Priority = model.TaskPriority(v)
	}
	if v, ok := fields["source"].(string); ok {
		task.Source = v
	}

	return task
}
