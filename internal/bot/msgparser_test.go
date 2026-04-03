package bot

import (
	"testing"
)

func TestParseMessage_Create(t *testing.T) {
	tests := []struct {
		name          string
		message       string
		wantIntent    Intent
		wantContent   string
		wantPriority  string
	}{
		{
			name: "simple create in Chinese", message: "记录 优化项目构建脚本",
			wantIntent: IntentCreate, wantContent: "优化项目构建脚本",
		},
		{
			name: "create with tags", message: "添加 学习新技术 #学习 #成长",
			wantIntent: IntentCreate, wantContent: "学习新技术",
		},
		{
			name: "create with high priority", message: "新建 紧急修复bug 优先级：高",
			wantIntent: IntentCreate, wantPriority: "high",
		},
		{
			name: "create in English", message: "create new feature idea",
			wantIntent: IntentCreate, wantContent: "new feature idea",
		},
		{
			name: "plain text defaults to create", message: "随便写点什么想法",
			wantIntent: IntentCreate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseMessage(tt.message)
			if result.Intent != tt.wantIntent {
				t.Errorf("ParseMessage(%q).Intent = %v, want %v", tt.message, result.Intent, tt.wantIntent)
			}
			if tt.wantContent != "" {
				if result.Params["content"] != tt.wantContent {
					t.Errorf("ParseMessage(%q).Params[content] = %q, want %q", tt.message, result.Params["content"], tt.wantContent)
				}
			}
			if tt.wantPriority != "" {
				if result.Params["priority"] != tt.wantPriority {
					t.Errorf("ParseMessage(%q).Params[priority] = %q, want %q", tt.message, result.Params["priority"], tt.wantPriority)
				}
			}
		})
	}
}

func TestParseMessage_List(t *testing.T) {
	result := ParseMessage("列出所有任务")
	if result.Intent != IntentList {
		t.Errorf("ParseMessage(\"列出所有任务\").Intent = %v, want %v", result.Intent, IntentList)
	}
}

func TestParseMessage_Delete(t *testing.T) {
	result := ParseMessage("删除 TASK-00001")
	if result.Intent != IntentDelete {
		t.Errorf("ParseMessage(\"删除 TASK-00001\").Intent = %v, want %v", result.Intent, IntentDelete)
	}
	if result.Params["task_id"] != "TASK-00001" {
		t.Errorf("task_id = %q, want %q", result.Params["task_id"], "TASK-00001")
	}
}

func TestParseMessage_Help(t *testing.T) {
	result := ParseMessage("帮助")
	if result.Intent != IntentHelp {
		t.Errorf("ParseMessage(\"帮助\").Intent = %v, want %v", result.Intent, IntentHelp)
	}
}

func TestParseMessage_Tags(t *testing.T) {
	result := ParseMessage("添加 学习新技术 #学习 #成长")
	if result.Params["tags"] != "学习,成长" {
		t.Errorf("tags = %q, want %q", result.Params["tags"], "学习,成长")
	}
}
