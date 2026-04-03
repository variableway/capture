package bitable

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/store"
)

// SyncEngine handles bidirectional sync between local tasks and Feishu Bitable.
type SyncEngine struct {
	bitableClient *Client
	store         store.Store
}

func NewSyncEngine(bitableClient *Client, store store.Store) *SyncEngine {
	return &SyncEngine{
		bitableClient: bitableClient,
		store:         store,
	}
}

type SyncResult struct {
	Pushed  int
	Pulled  int
	Errors  []string
}

// PushToBitable syncs local tasks to Feishu Bitable.
func (e *SyncEngine) PushToBitable(ctx context.Context, tasks []*model.Task) (*SyncResult, error) {
	result := &SyncResult{}

	for _, task := range tasks {
		fields := TaskToFields(task)

		if task.Sync.FeishuRecordID == "" {
			// Create new record
			recordID, err := e.bitableClient.CreateRecord(ctx, fields)
			if err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("%s: create failed: %v", task.ID, err))
				continue
			}
			task.Sync.FeishuRecordID = recordID
			now := time.Now()
			task.Sync.LastSyncedAt = &now
			_ = e.store.UpdateTask(ctx, task)
			result.Pushed++
		} else {
			// Update existing record
			if err := e.bitableClient.UpdateRecord(ctx, task.Sync.FeishuRecordID, fields); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("%s: update failed: %v", task.ID, err))
				continue
			}
			now := time.Now()
			task.Sync.LastSyncedAt = &now
			_ = e.store.UpdateTask(ctx, task)
			result.Pushed++
		}
	}

	return result, nil
}

// PullFromBitable syncs Feishu Bitable records to local tasks.
func (e *SyncEngine) PullFromBitable(ctx context.Context) (*SyncResult, error) {
	result := &SyncResult{}

	records, err := e.bitableClient.ListRecords(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	for _, record := range records {
		if record.Fields == nil {
			continue
		}

		fields := make(map[string]interface{})
		for k, v := range record.Fields {
			fields[k] = v
		}

		task := FieldsToTask(fields)

		// Check if task exists locally
		existing, err := e.store.GetTask(ctx, task.ID)
		if err != nil {
			// Task doesn't exist locally, create it
			task.Source = "feishu_bitable"
			if record.RecordId != nil {
				task.Sync.FeishuRecordID = *record.RecordId
			}
			if err := e.store.CreateTask(ctx, task); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("%s: create failed: %v", task.ID, err))
				continue
			}
			result.Pulled++
		} else {
			// Task exists, update if remote is newer
			if task.UpdatedAt.After(existing.UpdatedAt) {
				existing.Title = task.Title
				existing.Status = task.Status
				existing.Priority = task.Priority
				existing.UpdatedAt = time.Now()
				_ = e.store.UpdateTask(ctx, existing)
				result.Pulled++
			}
		}
	}

	return result, nil
}

// Sync performs bidirectional sync.
func (e *SyncEngine) Sync(ctx context.Context) (*SyncResult, error) {
	log.Println("Starting bidirectional sync...")

	// Push local changes
	tasks, err := e.store.ListTasks(ctx, model.TaskFilter{})
	if err != nil {
		return nil, err
	}

	var unsyncedTasks []*model.Task
	for _, t := range tasks {
		if t.Sync.LastSyncedAt == nil || t.UpdatedAt.After(*t.Sync.LastSyncedAt) {
			unsyncedTasks = append(unsyncedTasks, t)
		}
	}

	pushResult, err := e.PushToBitable(ctx, unsyncedTasks)
	if err != nil {
		return nil, err
	}

	// Pull remote changes
	pullResult, err := e.PullFromBitable(ctx)
	if err != nil {
		return pushResult, err
	}

	// Merge results
	pushResult.Pulled = pullResult.Pulled
	pushResult.Errors = append(pushResult.Errors, pullResult.Errors...)

	return pushResult, nil
}
