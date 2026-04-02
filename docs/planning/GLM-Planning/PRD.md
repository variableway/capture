# Capture 工具产品需求文档（PRD）

## 1. 产品概述

### 1.1 产品定位
Capture 是一个想法捕捉与任务执行管理工具，旨在帮助用户快速记录工作或生活中产生的随机想法，并能够将这些想法转化为可执行的任务，通过 AI Agent 自动化执行。

### 1.2 核心价值
- 快速捕捉：通过多种渠道（CLI、TUI、飞书、微信）快速记录想法
- 结构化存储：将想法转化为结构化的任务文档存储在本地
- 自动化执行：通过 AI Agent（Claude Code、Codex 等）自动执行任务
- 全流程追踪：通过看板和多维表格追踪任务状态
- 多渠道通知：通过飞书/微信推送任务执行进度

## 2. 功能需求

### 2.1 想法捕捉模块

#### 2.1.1 CLI/TUI 捕捉
- **功能描述**：通过命令行工具或终端界面快速记录想法
- **输入内容**：
  - 想法标题
  - 想法描述
  - 来源场景（可选）
  - 相关链接或参考资料（可选）
- **优先级**：P0

#### 2.1.2 飞书 Bot 捕捉
- **功能描述**：通过飞书机器人接收用户发送的想法
- **输入内容**：
  - 文本消息
  - 图片（可选，需 OCR 提取文本）
  - 来源链接（自动记录对话上下文）
- **优先级**：P0

#### 2.1.3 微信 Bot 捕捉
- **功能描述**：通过微信机器人接收用户发送的想法
- **输入内容**：
  - 文本消息
  - 语音消息（可选，需语音转文字）
- **优先级**：P1

#### 2.1.4 想法结构化
- **功能描述**：将非结构化的想法转化为结构化文档
- **处理逻辑**：
  - 提取关键信息
  - 自动分类标签
  - 生成任务描述
  - 关联相关资源
- **优先级**：P0

### 2.2 本地存储模块

#### 2.2.1 文档存储
- **功能描述**：将想法以 Markdown 格式存储在本地文件系统
- **存储结构**：
  ```
  /capture-data/
    /ideas/
      /{year}-{month}/
        /{idea-id}.md
    /tasks/
      /{task-id}.md
    /templates/
      /task-template.md
  ```
- **优先级**：P0

#### 2.2.2 元数据管理
- **功能描述**：管理想法和任务的元数据
- **元数据字段**：
  - 创建时间
  - 更新时间
  - 状态（待处理、已转化、执行中、已完成、已取消）
  - 标签
  - 来源渠道
  - 优先级
- **优先级**：P0

### 2.3 任务执行模块

#### 2.3.1 AI Agent 配置
- **功能描述**：配置执行任务的 AI Agent
- **配置项**：
  - Agent 类型（Claude Code、Codex、OpenAI 等）
  - 模型选择
  - API 密钥管理
  - 执行参数（超时时间、重试次数等）
- **优先级**：P0

#### 2.3.2 任务执行引擎
- **功能描述**：根据任务描述调用 AI Agent 执行
- **执行流程**：
  1. 解析任务文档
  2. 调用配置的 AI Agent
  3. 监控执行进度
  4. 记录执行日志
  5. 更新任务状态
- **优先级**：P0

#### 2.3.3 手动执行
- **功能描述**：用户可以通过 CLI 手动触发任务执行
- **优先级**：P1

### 2.4 看板管理模块

#### 2.4.1 看板视图
- **功能描述**：通过 TUI 展示任务看板
- **看板列**：
  - 待处理
  - 已转化
  - 执行中
  - 已完成
  - 已取消
- **操作**：
  - 拖拽移动任务
  - 点击查看详情
  - 快捷键操作
- **优先级**：P0

#### 2.4.2 表格视图
- **功能描述**：以表格形式展示任务列表
- **列信息**：
  - 任务 ID
  - 标题
  - 状态
  - 优先级
  - 创建时间
  - 更新时间
- **操作**：
  - 排序
  - 筛选
  - 批量操作
- **优先级**：P1

#### 2.4.3 飞书多维表格同步
- **功能描述**：将任务状态同步到飞书多维表格
- **同步策略**：
  - 双向同步
  - 定时同步（可配置间隔）
  - 手动触发同步
- **优先级**：P1

### 2.5 通知模块

#### 2.5.1 飞书通知
- **功能描述**：通过飞书 Bot 推送任务状态变更
- **通知场景**：
  - 任务创建
  - 任务开始执行
  - 任务执行成功
  - 任务执行失败
  - 任务需要人工介入
- **优先级**：P0

#### 2.5.2 微信通知
- **功能描述**：通过微信 Bot 推送任务状态变更
- **优先级**：P1

## 3. 非功能需求

### 3.1 性能要求
- CLI/TUI 响应时间 < 500ms
- 想法捕捉响应时间 < 2s
- 任务状态同步延迟 < 5s
- 支持同时管理 1000+ 任务

### 3.2 可靠性
- 本地数据持久化，支持离线操作
- 任务执行失败自动重试（最多 3 次）
- 数据备份机制（可选配置）

### 3.3 安全性
- API 密钥本地加密存储
- 敏感信息（如飞书/微信 Token）不记录到日志
- 支持 HTTPS 通信（飞书/微信 API）

### 3.4 可扩展性
- 插件化架构，支持自定义 AI Agent
- 支持自定义通知渠道
- 支持自定义存储后端（本地、云存储）

### 3.5 易用性
- CLI 提供友好的帮助文档
- TUI 支持键盘快捷键
- 提供示例配置文件
- 提供快速开始指南

## 4. 接口设计

### 4.1 CLI 接口

#### 4.1.1 捕捉想法
```bash
capture idea create --title "想法标题" --description "想法描述" --source "来源场景"
capture idea create --from-file idea.md
```

#### 4.1.2 查看任务
```bash
capture task list --status pending --priority high
capture task show <task-id>
```

#### 4.1.3 执行任务
```bash
capture task execute <task-id>
capture task execute --all --status pending
```

#### 4.1.4 看板操作
```bash
capture kanban
capture kanban move <task-id> --to completed
```

#### 4.1.5 配置管理
```bash
capture config set agent.type claude-code
capture config set agent.model claude-3-opus
capture config set notification.feishu.enabled true
```

### 4.2 Bot 接口

#### 4.2.1 飞书 Bot API
- **Webhook 端点**：`POST /api/v1/feishu/webhook`
- **消息格式**：
  ```json
  {
    "msg_type": "text",
    "content": {
      "text": "想法内容"
    },
    "context": {
      "user_id": "ou_xxx",
      "chat_id": "oc_xxx",
      "timestamp": 1234567890
    }
  }
  ```

#### 4.2.2 微信 Bot API
- **Webhook 端点**：`POST /api/v1/wechat/webhook`
- **消息格式**：
  ```json
  {
    "msgtype": "text",
    "text": {
      "content": "想法内容"
    },
    "fromusername": "user_xxx",
    "timestamp": 1234567890
  }
  ```

### 4.3 内部 API

#### 4.3.1 想法管理 API
```
POST   /api/v1/ideas              # 创建想法
GET    /api/v1/ideas              # 获取想法列表
GET    /api/v1/ideas/:id          # 获取想法详情
PUT    /api/v1/ideas/:id          # 更新想法
DELETE /api/v1/ideas/:id          # 删除想法
POST   /api/v1/ideas/:id/convert  # 转化为任务
```

#### 4.3.2 任务管理 API
```
POST   /api/v1/tasks              # 创建任务
GET    /api/v1/tasks              # 获取任务列表
GET    /api/v1/tasks/:id          # 获取任务详情
PUT    /api/v1/tasks/:id          # 更新任务
DELETE /api/v1/tasks/:id          # 删除任务
POST   /api/v1/tasks/:id/execute  # 执行任务
POST   /api/v1/tasks/:id/cancel   # 取消任务
```

#### 4.3.3 同步 API
```
POST   /api/v1/sync/feishu        # 同步到飞书多维表格
GET    /api/v1/sync/status        # 获取同步状态
```

## 5. 数据库设计

### 5.1 想法表（ideas）
| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | VARCHAR(36) | 想法 ID | PRIMARY KEY |
| title | VARCHAR(255) | 标题 | NOT NULL |
| description | TEXT | 描述 | |
| source | VARCHAR(255) | 来源场景 | |
| channel | VARCHAR(50) | 来源渠道（cli/feishu/wechat） | NOT NULL |
| context | JSON | 上下文信息（链接、图片等） | |
| tags | JSON | 标签列表 | |
| status | VARCHAR(20) | 状态（pending/converted） | NOT NULL |
| priority | VARCHAR(10) | 优先级（low/medium/high） | DEFAULT 'medium' |
| created_at | TIMESTAMP | 创建时间 | NOT NULL |
| updated_at | TIMESTAMP | 更新时间 | NOT NULL |

### 5.2 任务表（tasks）
| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | VARCHAR(36) | 任务 ID | PRIMARY KEY |
| idea_id | VARCHAR(36) | 关联想法 ID | FOREIGN KEY |
| title | VARCHAR(255) | 标题 | NOT NULL |
| description | TEXT | 任务描述 | |
| status | VARCHAR(20) | 状态（pending/executing/completed/failed/cancelled） | NOT NULL |
| priority | VARCHAR(10) | 优先级（low/medium/high） | DEFAULT 'medium' |
| agent_type | VARCHAR(50) | AI Agent 类型 | |
| agent_model | VARCHAR(50) | AI 模型 | |
| execution_log | TEXT | 执行日志 | |
| error_message | TEXT | 错误信息 | |
| created_at | TIMESTAMP | 创建时间 | NOT NULL |
| updated_at | TIMESTAMP | 更新时间 | NOT NULL |
| started_at | TIMESTAMP | 开始执行时间 | |
| completed_at | TIMESTAMP | 完成时间 | |

### 5.3 配置表（configs）
| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| key | VARCHAR(255) | 配置键 | PRIMARY KEY |
| value | TEXT | 配置值 | |
| encrypted | BOOLEAN | 是否加密 | DEFAULT FALSE |
| updated_at | TIMESTAMP | 更新时间 | NOT NULL |

### 5.4 同步记录表（sync_logs）
| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | VARCHAR(36) | 记录 ID | PRIMARY KEY |
| task_id | VARCHAR(36) | 任务 ID | FOREIGN KEY |
| platform | VARCHAR(50) | 平台（feishu/wechat） | NOT NULL |
| operation | VARCHAR(50) | 操作类型（create/update/delete） | NOT NULL |
| status | VARCHAR(20) | 同步状态（success/failed） | NOT NULL |
| error_message | TEXT | 错误信息 | |
| synced_at | TIMESTAMP | 同步时间 | NOT NULL |

## 6. 配置文件设计

### 6.1 主配置文件（capture.yaml）
```yaml
storage:
  type: local
  path: ./capture-data

database:
  type: sqlite
  path: ./capture-data/capture.db

agent:
  default_type: claude-code
  default_model: claude-3-opus
  timeout: 3600
  retry: 3

notification:
  feishu:
    enabled: true
    webhook_url: ${FEISHU_WEBHOOK_URL}
    app_id: ${FEISHU_APP_ID}
    app_secret: ${FEISHU_APP_SECRET}
  wechat:
    enabled: false
    webhook_url: ${WECHAT_WEBHOOK_URL}

sync:
  feishu_bitable:
    enabled: true
    app_token: ${FEISHU_BITABLE_APP_TOKEN}
    table_id: ${FEISHU_BITABLE_TABLE_ID}
    interval: 300

logging:
  level: info
  path: ./capture-data/logs
```

## 7. 用户故事

### 7.1 快速捕捉想法
作为用户，我希望通过 CLI 快速记录一个想法，以便不会遗忘。

### 7.2 移动端捕捉
作为用户，我希望通过飞书或微信发送消息来记录想法，以便随时随地捕捉灵感。

### 7.3 自动执行任务
作为用户，我希望系统自动将想法转化为任务并执行，以便节省时间。

### 7.4 查看任务进度
作为用户，我希望通过看板查看所有任务的状态，以便了解整体进度。

### 7.5 接收通知
作为用户，我希望在任务完成或失败时收到通知，以便及时了解执行情况。

## 8. 发布计划

### 8.1 MVP 版本（v0.1.0）
- CLI/TUI 捕捉想法
- 本地文件存储
- 基础看板视图
- 手动执行任务

### 8.2 Alpha 版本（v0.2.0）
- 飞书 Bot 集成
- AI Agent 执行任务
- 飞书通知

### 8.3 Beta 版本（v0.3.0）
- 飞书多维表格同步
- 微信 Bot 集成
- 高级看板功能

### 8.4 正式版本（v1.0.0）
- 完整功能
- 文档完善
- 性能优化
