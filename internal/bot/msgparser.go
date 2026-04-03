package bot

import (
	"regexp"
	"strings"
)

type Intent string

const (
	IntentCreate Intent = "create_task"
	IntentList   Intent = "list_tasks"
	IntentShow   Intent = "show_task"
	IntentDelete Intent = "delete_task"
	IntentUpdate Intent = "update_task"
	IntentHelp   Intent = "help"
	IntentUnknown Intent = "unknown"
)

type ParsedIntent struct {
	Intent Intent
	Params map[string]string
}

var intentPatterns = []struct {
	pattern string
	intent  Intent
}{
	{`(?i)(记录|添加|新建|创建|create|add|新)`, IntentCreate},
	{`(?i)(列出|查看|列表|list|show|所有|全部)`, IntentList},
	{`(?i)(删除|移除|delete|remove|删)`, IntentDelete},
	{`(?i)(更新|修改|update|edit|改)`, IntentUpdate},
	{`(?i)(帮助|help|\?)`, IntentHelp},
}

var taskIDPattern = regexp.MustCompile(`(TASK-\d+)`)
var tagPattern = regexp.MustCompile(`#(\S+)`)
var priorityHighPattern = regexp.MustCompile(`(?i)(优先级[：:]\s*(高|high)|紧急|urgent)`)
var priorityLowPattern = regexp.MustCompile(`(?i)(优先级[：:]\s*(低|low))`)

// ParseMessage detects intent and extracts parameters from a message.
func ParseMessage(message string) ParsedIntent {
	message = strings.TrimSpace(message)

	result := ParsedIntent{
		Intent: IntentUnknown,
		Params: make(map[string]string),
	}

	// Detect intent
	for _, ip := range intentPatterns {
		matched, _ := regexp.MatchString(ip.pattern, message)
		if matched {
			result.Intent = ip.intent
			break
		}
	}

	// If no intent matched and message looks like plain text, treat as create
	if result.Intent == IntentUnknown && len(message) > 0 {
		result.Intent = IntentCreate
	}

	// Extract task ID
	if matches := taskIDPattern.FindStringSubmatch(message); len(matches) > 1 {
		result.Params["task_id"] = matches[1]
	}

	// Extract tags
	tags := tagPattern.FindAllStringSubmatch(message, -1)
	if len(tags) > 0 {
		var tagList []string
		for _, t := range tags {
			tagList = append(tagList, t[1])
		}
		result.Params["tags"] = strings.Join(tagList, ",")
	}

	// Extract priority
	if priorityHighPattern.MatchString(message) {
		result.Params["priority"] = "high"
	} else if priorityLowPattern.MatchString(message) {
		result.Params["priority"] = "low"
	}

	// Extract content (strip commands and metadata)
	content := message
	for _, cmd := range []string{"记录", "添加", "新建", "创建", "create", "add"} {
		content = strings.TrimPrefix(content, cmd)
	}
	content = strings.TrimSpace(content)
	// Remove task ID, tags, priority from content
	content = taskIDPattern.ReplaceAllString(content, "")
	content = tagPattern.ReplaceAllString(content, "")
	content = priorityHighPattern.ReplaceAllString(content, "")
	content = priorityLowPattern.ReplaceAllString(content, "")
	content = strings.TrimSpace(content)
	if content != "" {
		result.Params["content"] = content
	}

	return result
}
