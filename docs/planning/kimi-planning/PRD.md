# Capture Tool - Product Requirements Document

## 1. 产品概述

Capture 是一个用于捕捉、管理和执行随机想法的智能化任务管理工具。它支持通过多种渠道（TUI、CLI、Bot）记录想法，并通过可配置的 AI Agent 自动执行，同时提供可视化的看板管理和多渠道通知。

## 2. 用户故事

1. 作为用户，我可以通过 CLI 快速记录工作中产生的随机想法为待办事项
2. 作为用户，我可以通过飞书/微信 Bot 发送消息记录随机想法，并保留上下文
3. 作为用户，我希望所有想法都以本地文件形式存储，便于管理和版本控制
4. 作为用户，我可以通过 CLI 或 Bot 触发想法的执行
5. 作为用户，我可以通过看板查看所有任务的状态和进度
6. 作为用户，我希望通过 Bot 接收任务执行状态和完成通知
7. 作为用户，我可以配置使用不同的 AI Agent (Claude Code, Codex 等) 和模型
8. 作为用户，我可以通过看板或表格管理任务状态
9. 作为用户，我希望任务状态可以与飞书多维表格同步

## 3. 功能需求

### 3.1 核心功能模块

#### 3.1.1 想法捕获模块 (Capture Module)

| 功能点 | 描述 | 优先级 |
|--------|------|--------|
| CLI 捕获 | 通过命令行快速记录想法，支持标题、描述、标签 | P0 |
| TUI 捕获 | 提供交互式终端界面，支持表单式输入 | P1 |
| Bot 捕获 | 接收飞书/微信 Bot 消息，自动解析内容 | P0 |
| 上下文记录 | 自动记录想法来源、时间、场景、触发方式 | P1 |
| 文件存储 | 将想法持久化为本地 Markdown/YAML 文件 | P0 |

#### 3.1.2 任务管理模块 (Task Management Module)

| 功能点 | 描述 | 优先级 |
|--------|------|--------|
| 任务状态机 | 待办/进行中/已完成/已取消/已归档 | P0 |
| 标签系统 | 支持多标签分类，便于筛选和检索 | P1 |
| 优先级管理 | 高/中/低优先级标记 | P1 |
| 任务搜索 | 支持关键词、标签、状态、时间范围搜索 | P1 |
| 任务编辑 | 修改任务内容、状态、优先级 | P1 |
| 任务删除 | 软删除支持，保留历史记录 | P2 |

#### 3.1.3 看板模块 (Kanban Module)

| 功能点 | 描述 | 优先级 |
|--------|------|--------|
| 状态看板 | 按状态列展示任务卡片 | P0 |
| 拖拽操作 | 支持拖拽变更任务状态 | P1 |
| 任务筛选 | 按标签、优先级、时间筛选 | P1 |
| 详情弹窗 | 点击卡片查看/编辑任务详情 | P1 |
| 快捷操作 | 快速标记完成、添加标签 | P2 |

#### 3.1.4 执行引擎模块 (Execution Module)

| 功能点 | 描述 | 优先级 |
|--------|------|--------|
| 本地执行 | 在本地环境执行想法（代码生成、文件操作等） | P0 |
| AI Agent 集成 | 支持调用 Claude Code、Codex、Kimi CLI 等 | P0 |
| 模型配置 | 可配置使用的 AI 模型（GPT-4, Claude-3, Kimi 等） | P1 |
| 执行沙箱 | 隔离的执行环境，确保安全 | P2 |
| 执行日志 | 记录执行过程的完整日志 | P0 |
| 回滚机制 | 执行失败时的回滚能力 | P2 |

#### 3.1.5 通知模块 (Notification Module)

| 功能点 | 描述 | 优先级 |
|--------|------|--------|
| 状态通知 | 任务状态变更时发送通知 | P1 |
| 执行结果通知 | 任务执行完成后发送结果摘要 | P1 |
| 定时提醒 | 可配置的任务提醒 | P2 |
| 飞书通知 | 通过飞书 Bot 发送通知 | P1 |
| 微信通知 | 通过微信 Bot 发送通知 | P2 |
| 邮件通知 | 通过邮件发送通知 | P3 |

#### 3.1.6 同步模块 (Sync Module)

| 功能点 | 描述 | 优先级 |
|--------|------|--------|
| 飞书多维表格同步 | 双向同步任务状态到飞书多维表格 | P1 |
| 增量同步 | 只同步变更的数据，提高效率 | P2 |
| 冲突处理 | 处理本地和远程的数据冲突 | P2 |
| 手动同步 | 支持手动触发全量同步 | P1 |

### 3.2 功能详细设计

#### 3.2.1 CLI 接口设计

```
capture add <title> [options]     # 添加新想法
  -d, --description <desc>        # 描述
  -t, --tags <tags>               # 标签，逗号分隔
  -p, --priority <level>          # 优先级: high/medium/low
  
capture list [options]            # 列出任务
  -s, --status <status>           # 按状态筛选
  -t, --tag <tag>                 # 按标签筛选
  -p, --priority <level>          # 按优先级筛选
  
capture show <id>                 # 查看任务详情
capture edit <id>                 # 编辑任务
capture delete <id>               # 删除任务
capture execute <id>              # 执行任务
capture status <id> <status>      # 修改状态

capture kanban                    # 启动 TUI 看板
capture sync                      # 手动触发同步
capture config                    # 配置管理
```

#### 3.2.2 Bot 消息格式设计

**飞书 Bot 消息解析规则：**

```
# 记录想法
@CaptureBot 记录：优化项目构建脚本，减少构建时间
标签：#优化 #构建
优先级：高

# 执行任务
@CaptureBot 执行 TASK-001

# 查询任务
@CaptureBot 列出待办任务
@CaptureBot 查找标签为"优化"的任务
```

**消息元数据自动捕获：**
- 发送者信息
- 发送时间
- 对话上下文（引用消息）
- 消息来源群组/频道

#### 3.2.3 文件存储格式

**任务文件结构：**

```
~/.capture/
├── tasks/
│   ├── 2024/
│   │   ├── 04/
│   │   │   ├── TASK-001.md       # 任务文件
│   │   │   └── TASK-002.md
│   │   └── 05/
│   └── index.yaml                # 任务索引
├── config.yaml                   # 全局配置
├── logs/                         # 执行日志
└── templates/                    # 任务模板
```

**任务文件格式 (Markdown + Frontmatter):**

```yaml
---
id: TASK-001
title: "优化项目构建脚本"
description: "减少构建时间，提高开发效率"
status: todo  # todo, in_progress, done, cancelled, archived
priority: high  # high, medium, low
tags: ["优化", "构建", "devops"]
created_at: 2024-04-02T10:30:00+08:00
updated_at: 2024-04-02T10:30:00+08:00
source: cli  # cli, tui, feishu_bot, wechat_bot
context:
  trigger: "工作中突然产生的想法"
  location: "办公室"
  related_to: ""
execution:
  agent: claude-code
  model: claude-3-sonnet
  result: null
  logs: []
  started_at: null
  completed_at: null
sync:
  feishu_bitable:
    record_id: ""
    last_sync: null
---

# 任务详情

## 执行计划
1. 分析当前构建流程
2. 识别性能瓶颈
3. 优化脚本
4. 测试验证

## 执行结果

## 备注
```

## 4. 非功能需求

### 4.1 性能需求

| 指标 | 目标值 | 说明 |
|------|--------|------|
| 任务创建响应时间 | < 500ms | 本地文件操作 |
| 任务列表加载 | < 1s | 1000 个任务以内 |
| 看板初始化 | < 2s | 包含渲染 |
| Bot 消息响应 | < 3s | 从接收到确认 |
| 同步操作 | < 5s | 单次同步 |

### 4.2 可靠性需求

| 指标 | 目标值 | 说明 |
|------|--------|------|
| 数据持久化 | 100% | 所有任务必须落盘 |
| 执行失败恢复 | 支持 | 可重试失败的任务 |
| 配置备份 | 自动 | 定期备份配置 |
| 日志保留 | 30天 | 执行日志保留期限 |

### 4.3 安全需求

| 需求 | 说明 | 优先级 |
|------|------|--------|
| 本地数据加密 | 敏感配置加密存储 | P2 |
| Bot 认证 | 验证 Bot 请求来源 | P1 |
| 执行权限控制 | 限制执行脚本的权限范围 | P1 |
| 敏感信息过滤 | 自动过滤日志中的敏感信息 | P2 |

### 4.4 兼容性需求

| 需求 | 说明 | 优先级 |
|------|------|--------|
| 跨平台 | 支持 macOS、Linux、Windows | P1 |
| Python 版本 | 支持 Python 3.9+ | P0 |
| 终端兼容 | 支持常见终端模拟器 | P1 |
| 编辑器兼容 | 与 VS Code、Vim 等集成 | P2 |

### 4.5 可扩展性需求

| 需求 | 说明 | 优先级 |
|------|------|--------|
| 插件系统 | 支持自定义插件扩展 | P2 |
| 多 Bot 支持 | 易于添加新的 Bot 适配器 | P1 |
| 多 Agent 支持 | 易于集成新的 AI Agent | P1 |
| 自定义模板 | 支持自定义任务模板 | P2 |

## 5. 数据库设计

### 5.1 本地存储设计 (SQLite + 文件系统)

由于 Capture 工具以本地文件为主存储，使用 SQLite 作为索引和查询优化。

**表结构：**

```sql
-- 任务主表
CREATE TABLE tasks (
    id TEXT PRIMARY KEY,                    -- TASK-XXXXX
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'todo',    -- todo, in_progress, done, cancelled, archived
    priority TEXT DEFAULT 'medium',         -- high, medium, low
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    source TEXT,                            -- cli, tui, feishu_bot, wechat_bot
    file_path TEXT NOT NULL,                -- 对应的 Markdown 文件路径
    execution_agent TEXT,                   -- 执行使用的 Agent
    execution_status TEXT,                  -- pending, running, success, failed
    feishu_record_id TEXT                   -- 飞书多维表格记录 ID
);

-- 标签表
CREATE TABLE tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    color TEXT,                             -- 标签颜色
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 任务-标签关联表
CREATE TABLE task_tags (
    task_id TEXT NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (task_id, tag_id),
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- 执行历史表
CREATE TABLE execution_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id TEXT NOT NULL,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    status TEXT,                            -- success, failed, cancelled
    output TEXT,                            -- 执行输出
    error_message TEXT,
    agent_used TEXT,
    model_used TEXT,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);

-- 同步记录表
CREATE TABLE sync_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sync_type TEXT NOT NULL,                -- feishu_bitable, etc.
    direction TEXT NOT NULL,                -- push, pull, bidirectional
    started_at TIMESTAMP NOT NULL,
    completed_at TIMESTAMP,
    status TEXT,                            -- success, failed
    records_count INTEGER,
    error_message TEXT
);

-- 配置表
CREATE TABLE config (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_priority ON tasks(priority);
CREATE INDEX idx_tasks_created_at ON tasks(created_at);
CREATE INDEX idx_tasks_source ON tasks(source);
CREATE INDEX idx_execution_logs_task_id ON execution_logs(task_id);
```

### 5.2 配置结构设计

**config.yaml:**

```yaml
# 应用配置
app:
  name: "Capture"
  version: "1.0.0"
  data_dir: "~/.capture"
  
# 默认设置
defaults:
  priority: "medium"
  tags: []
  editor: "vim"  # 编辑任务时使用的编辑器
  
# AI Agent 配置
agents:
  claude-code:
    enabled: true
    command: "claude"
    default_model: "claude-3-sonnet-20240229"
    timeout: 300
  codex:
    enabled: false
    command: "codex"
    default_model: "gpt-4o"
    timeout: 300
  kimi-cli:
    enabled: false
    command: "kimi"
    default_model: "kimi-latest"
    timeout: 300
    
# Bot 配置
bots:
  feishu:
    enabled: true
    app_id: "${FEISHU_APP_ID}"
    app_secret: "${FEISHU_APP_SECRET}"
    encrypt_key: "${FEISHU_ENCRYPT_KEY}"
    verification_token: "${FEISHU_VERIFICATION_TOKEN}"
    webhook_url: ""
  wechat:
    enabled: false
    webhook_url: ""
    
# 通知配置
notifications:
  enabled: true
  on_status_change: true
  on_execution_complete: true
  channels:
    - feishu
    
# 同步配置
sync:
  feishu_bitable:
    enabled: false
    app_token: "${FEISHU_APP_TOKEN}"
    table_id: "${FEISHU_TABLE_ID}"
    sync_interval: 3600  # 秒，0 表示仅手动同步
    auto_sync: false
    
# 执行配置
execution:
  default_agent: "claude-code"
  dry_run: false
  confirm_before_execute: true
  max_execution_time: 600  # 秒
  allowed_commands: []     # 允许执行的命令白名单
  blocked_commands: []     # 禁止执行的命令黑名单
```

## 6. 接口设计

### 6.1 内部 API 设计 (Python)

```python
# 核心服务接口
class TaskService:
    def create(self, title: str, **kwargs) -> Task
    def get(self, task_id: str) -> Task
    def update(self, task_id: str, **kwargs) -> Task
    def delete(self, task_id: str) -> bool
    def list(self, filters: TaskFilter) -> List[Task]
    def search(self, query: str) -> List[Task]

class ExecutionService:
    def execute(self, task_id: str, agent: str = None) -> ExecutionResult
    def get_logs(self, task_id: str) -> List[ExecutionLog]
    def cancel(self, execution_id: str) -> bool

class SyncService:
    def sync_to_feishu(self, task_ids: List[str] = None) -> SyncResult
    def sync_from_feishu(self) -> SyncResult
    def resolve_conflict(self, task_id: str, resolution: str) -> Task

class NotificationService:
    def send(self, message: NotificationMessage, channels: List[str])
    def notify_status_change(self, task: Task, old_status: str)
    def notify_execution_complete(self, task: Task, result: ExecutionResult)

class BotService:
    def handle_message(self, message: BotMessage) -> str
    def register_handler(self, intent: str, handler: Callable)
```

### 6.2 Bot Webhook 接口

**飞书 Bot Webhook:**

```
POST /webhook/feishu
Content-Type: application/json

{
  "schema": "2.0",
  "header": {
    "event_id": "xxxx",
    "token": "verification_token",
    "create_time": "1234567890",
    "event_type": "im.message.receive_v1"
  },
  "event": {
    "message": {
      "message_id": "msg_xxx",
      "content": "{\"text\": \"@CaptureBot 记录：xxx\"}",
      "chat_type": "group",
      "chat_id": "chat_xxx"
    },
    "sender": {
      "sender_id": {
        "union_id": "xxx"
      }
    }
  }
}
```

**响应格式：**

```json
{
  "content": "任务已创建：TASK-001 - 优化项目构建脚本"
}
```

### 6.3 TUI 界面设计

**看板界面布局：**

```
┌─────────────────────────────────────────────────────────────────┐
│ Capture Kanban                                      [? Help]    │
├─────────────────────────────────────────────────────────────────┤
│ [All] [Todo] [In Progress] [Done] [Search: ______] [Filter ▼]   │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐             │
│  │ 📋 TODO (3) │  │ 🔄 IN       │  │ ✅ DONE (2) │             │
│  │             │  │    PROGRESS │  │             │             │
│  │ ┌─────────┐ │  │    (1)      │  │ ┌─────────┐ │             │
│  │ │TASK-001 │ │  │             │  │ │TASK-004 │ │             │
│  │ │High ⚡  │ │  │ ┌─────────┐ │  │ │Low      │ │             │
│  │ │#优化    │ │  │ │TASK-003 │ │  │ │#文档    │ │             │
│  │ └─────────┘ │  │ │Medium   │ │  │ └─────────┘ │             │
│  │             │  │ │#功能    │ │  │             │             │
│  │ ┌─────────┐ │  │ └─────────┘ │  │ ┌─────────┐ │             │
│  │ │TASK-002 │ │  │             │  │ │TASK-005 │ │             │
│  │ │Low      │ │  │             │  │ └─────────┘ │             │
│  │ │#重构    │ │  │             │  │             │             │
│  │ └─────────┘ │  │             │  │             │             │
│  └─────────────┘  └─────────────┘  └─────────────┘             │
├─────────────────────────────────────────────────────────────────┤
│ [A] Add  [E] Edit  [D] Delete  [X] Execute  [Q] Quit            │
└─────────────────────────────────────────────────────────────────┘
```

## 7. 错误处理与日志

### 7.1 错误码定义

| 错误码 | 描述 | 处理方式 |
|--------|------|----------|
| CAP-001 | 任务不存在 | 提示用户检查任务 ID |
| CAP-002 | 文件写入失败 | 重试或提示检查权限 |
| CAP-003 | 执行超时 | 自动终止，记录状态 |
| CAP-004 | Agent 调用失败 | 记录错误，通知用户 |
| CAP-005 | 同步失败 | 记录错误，支持手动重试 |
| CAP-006 | 配置错误 | 启动时校验，提示修复 |
| CAP-007 | Bot 认证失败 | 拒绝请求，记录日志 |

### 7.2 日志级别

- **DEBUG**: 详细的调试信息
- **INFO**: 常规操作记录
- **WARNING**: 警告信息
- **ERROR**: 错误信息
- **CRITICAL**: 严重错误

## 8. 部署与安装

### 8.1 安装方式

```bash
# pip 安装
pip install capture-tool

# 初始化
capture init

# 配置
capture config set editor vim
capture config set agents.claude-code.enabled true
```

### 8.2 环境变量

```bash
CAPTURE_DATA_DIR          # 数据目录
CAPTURE_CONFIG_FILE       # 配置文件路径
FEISHU_APP_ID            # 飞书应用 ID
FEISHU_APP_SECRET        # 飞书应用密钥
OPENAI_API_KEY           # OpenAI API Key (for Codex)
ANTHROPIC_API_KEY        # Anthropic API Key (for Claude)
```

---

**文档版本**: 1.0  
**最后更新**: 2024-04-02
