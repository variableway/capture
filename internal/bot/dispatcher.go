package bot

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/service"
)

// Handler processes parsed bot messages.
type Handler func(ctx context.Context, intent ParsedIntent, sender string) (string, error)

// Dispatcher routes parsed messages to appropriate handlers.
type Dispatcher struct {
	taskSvc *service.TaskService
}

func NewDispatcher(taskSvc *service.TaskService) *Dispatcher {
	return &Dispatcher{
		taskSvc: taskSvc,
	}
}

// ProcessMessage handles an incoming bot message and returns a response.
func (d *Dispatcher) ProcessMessage(ctx context.Context, message, sender string) string {
	intent := ParseMessage(message)
	log.Printf("Bot message from %s: intent=%s, params=%v", sender, intent.Intent, intent.Params)

	switch intent.Intent {
	case IntentCreate:
		return d.handleCreate(ctx, intent, sender)
	case IntentList:
		return d.handleList(ctx, intent)
	case IntentDelete:
		return d.handleDelete(ctx, intent)
	case IntentHelp:
		return handleHelp()
	default:
		return "未知命令。发送 \"帮助\" 查看可用命令。"
	}
}

func (d *Dispatcher) handleCreate(ctx context.Context, intent ParsedIntent, sender string) string {
	content := intent.Params["content"]
	if content == "" {
		return "请提供要记录的内容。例如：记录 优化项目构建脚本"
	}

	opts := []service.TaskOption{
		service.WithSource("feishu_bot"),
		service.WithContext(model.TaskContext{
			Trigger: "feishu_message",
		}),
	}

	if p, ok := intent.Params["priority"]; ok {
		opts = append(opts, service.WithPriority(model.TaskPriority(p)))
	}
	if tags, ok := intent.Params["tags"]; ok {
		opts = append(opts, service.WithTags(strings.Split(tags, ",")))
	}

	task, err := d.taskSvc.Create(ctx, content, opts...)
	if err != nil {
		return fmt.Sprintf("创建失败: %v", err)
	}

	return fmt.Sprintf("已创建: %s - %s", task.ID, task.Title)
}

func (d *Dispatcher) handleList(ctx context.Context, intent ParsedIntent) string {
	filter := model.TaskFilter{}
	tasks, err := d.taskSvc.List(ctx, filter)
	if err != nil {
		return fmt.Sprintf("查询失败: %v", err)
	}

	if len(tasks) == 0 {
		return "暂无任务。"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("共 %d 个任务:\n", len(tasks)))
	for _, t := range tasks {
		sb.WriteString(fmt.Sprintf("  %s [%s] %s\n", t.ID, t.Status, t.Title))
	}
	return sb.String()
}

func (d *Dispatcher) handleDelete(ctx context.Context, intent ParsedIntent) string {
	taskID := intent.Params["task_id"]
	if taskID == "" {
		return "请指定要删除的任务ID。例如：删除 TASK-00001"
	}

	if err := d.taskSvc.Delete(ctx, taskID); err != nil {
		return fmt.Sprintf("删除失败: %v", err)
	}
	return fmt.Sprintf("已删除: %s", taskID)
}

func handleHelp() string {
	return `Capture Bot 命令：
- 记录 <内容>  — 创建新任务
- 列出 — 查看所有任务
- 删除 <TASK-ID> — 删除任务
- 帮助 — 显示此帮助

示例：
  记录 优化项目构建脚本 #优化
  列出
  删除 TASK-00001`
}
