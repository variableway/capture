# Capture

> 捕捉灵感，管理任务 — 通过 Terminal 或飞书 Bot 快速记录想法，保存为 Markdown 文件并同步到飞书多维表格。

## 功能特性

- **CLI 命令行** — 快速创建、查看、编辑、删除任务
- **Task Center 工作流** — 支持从 inbox、mindstorm、analysis 到 dispatch、execution、review 的阶段化流转
- **TUI 看板** — 终端交互式看板界面，可视化任务状态
- **飞书 Bot** — 通过飞书机器人接收消息，自动创建任务
- **飞书多维表格同步** — 双向同步任务到飞书 Bitable
- **Markdown 存储** — 任务以 Markdown + YAML frontmatter 格式存储
- **Agent 分派记录** — 为任务记录 agent、model、repo、worktree、terminal 等执行上下文
- **双模式 Bot** — 支持 Webhook 和 WebSocket 两种连接模式

## 安装

```bash
# 从源码构建
git clone https://github.com/variableway/capture.git
cd capture
go build -o capture .

# 或直接安装
go install ./...
```

## 快速开始

```bash
# 初始化
capture init

# 创建任务
capture add "优化项目构建脚本"
capture add "学习 Go 语言" -d "完成官方教程" -t "学习,Go" -p high
capture add "设计本地任务中心" --stage analysis

# 查看任务列表
capture list
capture list --status todo
capture list --stage analysis

# 查看任务详情
capture show TASK-00001

# 更新任务
capture edit TASK-00001 --title "新的标题" --priority high
capture edit TASK-00001 --stage planning

# 修改状态
capture status TASK-00001 in_progress
capture status TASK-00001 done

# 修改任务阶段
capture stage TASK-00001 mindstorm
capture stage TASK-00001 dispatch

# 分配给 AI Agent
capture assign TASK-00001 --agent codex --model gpt-5 --repo ~/workspace/project --worktree ~/workspace/project --terminal term-1

# 删除任务
capture delete TASK-00001
capture delete TASK-00001 --force

# 启动 TUI 看板
capture kanban

# 同步到飞书多维表格
capture sync --direction push
capture sync --direction bidirectional
```

## 飞书 Bot 配置

### 1. 创建飞书应用

1. 访问 [飞书开放平台](https://open.feishu.cn/app)
2. 点击「创建企业自建应用」
3. 记录 `App ID` 和 `App Secret`

### 2. 启用机器人功能

1. 在应用管理页面，找到「应用功能」→「机器人」
2. 开启机器人功能

### 3. 配置权限

添加以下权限：
- `im:message` — 获取与发送消息
- `im:message.receive_v1` — 接收消息事件
- `bitable:app` — 读写多维表格
- `bitable:app:readonly` — 读取多维表格

### 4. 配置事件订阅

#### WebSocket 模式（推荐，本地开发）

1. 在「事件订阅」页面选择「使用长连接接收事件」
2. 添加事件：`im.message.receive_v1`
3. 无需公网 URL

#### Webhook 模式（生产部署）

1. 在「事件订阅」页面配置请求 URL：`https://your-domain.com/webhook/feishu`
2. 添加事件：`im.message.receive_v1`
3. 记录 `Verification Token` 和 `Encrypt Key`

### 5. 配置多维表格（可选）

1. 创建一个飞书多维表格
2. 添加以下列：
   - `task_id` (文本)
   - `title` (文本)
   - `status` (单选: todo, in_progress, done, cancelled, archived)
   - `priority` (单选: high, medium, low)
   - `source` (文本)
   - `created_at` (日期)
   - `updated_at` (日期)
3. 记录 `App Token`（URL 中 `/base/` 后面的部分）和 `Table ID`

### 6. 启动 Bot

```bash
# 设置环境变量
export FEISHU_APP_ID="cli_xxx"
export FEISHU_APP_SECRET="xxx"

# WebSocket 模式（推荐本地开发）
capture bot serve --mode websocket

# Webhook 模式（需要公网 URL）
export FEISHU_VERIFICATION_TOKEN="xxx"
export FEISHU_ENCRYPT_KEY="xxx"
capture bot serve --mode webhook --port 8080
```

### 7. 使用飞书 Bot

在飞书中给机器人发送消息：

```
记录 优化项目构建脚本 #优化
列出
删除 TASK-00001
帮助
```

## 配置

配置文件位于 `~/.capture/config.yaml`：

```yaml
app:
  name: Capture
  version: 0.1.0
  data_dir: ~/.capture

defaults:
  priority: medium
  editor: vim

feishu:
  app_id: ${FEISHU_APP_ID}
  app_secret: ${FEISHU_APP_SECRET}

bitable:
  enabled: false
  app_token: ${FEISHU_BITABLE_APP_TOKEN}
  table_id: ${FEISHU_BITABLE_TABLE_ID}

bot:
  mode: websocket
  port: 8080
```

## 任务存储格式

任务以 Markdown 文件存储在 `~/.capture/tasks/YYYY/MM/TASK-NNNNN.md`：

```markdown
---
id: TASK-00001
title: "优化项目构建脚本"
status: todo
stage: analysis
priority: high
tags: [优化, 构建]
created_at: 2026-04-03T10:30:00+08:00
updated_at: 2026-04-03T10:30:00+08:00
source: cli
dispatch:
  agent: "codex"
  model: "gpt-5"
  repository: "/Users/demo/workspace/project"
  worktree: "/Users/demo/workspace/project"
  terminal_session: "term-1"
  assigned_at: 2026-04-03T11:00:00+08:00
sync:
  feishu_record_id: ""
  last_synced_at: null
---

## Description

减少构建时间，提高开发效率

## Notes

(备注)
```

## 开发

```bash
# 安装依赖
go mod tidy

# 构建
go build -o capture .

# 运行测试
go test ./...

# 运行测试（详细输出）
go test ./... -v
```

## 技术栈

| 组件 | 技术 |
|------|------|
| CLI | Cobra |
| TUI | bubbletea + lipgloss |
| 配置 | Viper |
| 存储 | Markdown + SQLite (modernc.org/sqlite) |
| 飞书 SDK | oapi-sdk-go/v3 |
| 测试 | testify |

## License

MIT
