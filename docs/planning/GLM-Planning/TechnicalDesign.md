# Capture 工具技术方案文档

## 1. 技术架构

### 1.1 整体架构

```
┌─────────────────────────────────────────────────────────────┐
│                        用户交互层                            │
├─────────────┬──────────────┬─────────────┬─────────────────┤
│  CLI/TUI    │  飞书 Bot    │  微信 Bot   │  Web Dashboard  │
└─────────────┴──────────────┴─────────────┴─────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                        API 网关层                            │
│              (路由、认证、限流、日志)                        │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                        核心服务层                            │
├─────────────┬──────────────┬─────────────┬─────────────────┤
│ 想法捕捉服务│ 任务管理服务 │ 执行引擎服务│  通知服务       │
├─────────────┴──────────────┴─────────────┴─────────────────┤
│                      同步服务                               │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                        数据持久层                            │
├──────────────────────┬──────────────────────────────────────┤
│   SQLite 数据库      │        本地文件系统                  │
└──────────────────────┴──────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                    外部集成层                                │
├──────────────────┬──────────────────┬──────────────────────┤
│  AI Agent 集成   │  飞书 API 集成   │  微信 API 集成       │
└──────────────────┴──────────────────┴──────────────────────┘
```

### 1.2 模块职责

#### 1.2.1 用户交互层
- **CLI/TUI**：提供命令行和终端界面交互
- **飞书 Bot**：接收飞书消息并处理
- **微信 Bot**：接收微信消息并处理
- **Web Dashboard**：可选的 Web 管理界面

#### 1.2.2 API 网关层
- 路由分发
- 请求认证
- 限流控制
- 日志记录

#### 1.2.3 核心服务层
- **想法捕捉服务**：处理想法的创建、存储、转化
- **任务管理服务**：管理任务的生命周期
- **执行引擎服务**：调用 AI Agent 执行任务
- **通知服务**：发送任务状态通知
- **同步服务**：与飞书多维表格同步

#### 1.2.4 数据持久层
- **SQLite**：存储结构化数据（想法、任务、配置等）
- **本地文件系统**：存储 Markdown 文档

#### 1.2.5 外部集成层
- **AI Agent 集成**：Claude Code、Codex、OpenAI 等
- **飞书 API 集成**：消息接收、多维表格操作
- **微信 API 集成**：消息接收、发送

## 2. 技术选型

### 2.1 编程语言
**选择：Python 3.10+**

**理由**：
- 丰富的 AI/ML 生态（OpenAI SDK、Anthropic SDK）
- 优秀的 CLI 库（Click、Rich、Textual）
- 良好的异步支持（asyncio）
- 快速开发效率

### 2.2 核心框架和库

#### 2.2.1 CLI/TUI 框架
- **Click**：命令行参数解析
- **Rich**：终端富文本显示
- **Textual**：TUI 框架（用于看板视图）

#### 2.2.2 Web 框架
- **FastAPI**：轻量级异步 Web 框架
- **Uvicorn**：ASGI 服务器

#### 2.2.3 数据库
- **SQLite**：嵌入式数据库
- **SQLAlchemy**：ORM 框架
- **Alembic**：数据库迁移工具

#### 2.2.4 异步任务
- **asyncio**：异步 I/O
- **Celery**（可选）：任务队列（用于长时间执行的任务）

#### 2.2.5 AI Agent SDK
- **Anthropic SDK**：Claude API 客户端
- **OpenAI SDK**：OpenAI API 客户端

#### 2.2.6 第三方集成
- **feishu-sdk**：飞书开放平台 SDK
- **wechatpy**：微信 SDK

#### 2.2.7 工具库
- **Pydantic**：数据验证
- **PyYAML**：配置文件解析
- **cryptography**：加密解密
- **python-dotenv**：环境变量管理

### 2.3 项目结构

```
capture/
├── capture/
│   ├── __init__.py
│   ├── cli/                    # CLI 相关代码
│   │   ├── __init__.py
│   │   ├── main.py            # CLI 入口
│   │   ├── idea.py            # 想法相关命令
│   │   ├── task.py            # 任务相关命令
│   │   └── kanban.py          # 看板命令
│   ├── tui/                    # TUI 相关代码
│   │   ├── __init__.py
│   │   ├── app.py             # TUI 应用
│   │   └── widgets/           # TUI 组件
│   ├── api/                    # API 相关代码
│   │   ├── __init__.py
│   │   ├── server.py          # FastAPI 服务
│   │   ├── routes/            # 路由
│   │   │   ├── ideas.py
│   │   │   ├── tasks.py
│   │   │   ├── feishu.py
│   │   │   └── wechat.py
│   │   └── middleware/        # 中间件
│   ├── services/               # 核心服务
│   │   ├── __init__.py
│   │   ├── idea_service.py
│   │   ├── task_service.py
│   │   ├── execution_engine.py
│   │   ├── notification_service.py
│   │   └── sync_service.py
│   ├── agents/                 # AI Agent 集成
│   │   ├── __init__.py
│   │   ├── base.py            # Agent 基类
│   │   ├── claude_code.py
│   │   ├── codex.py
│   │   └── openai_agent.py
│   ├── integrations/           # 第三方集成
│   │   ├── __init__.py
│   │   ├── feishu.py
│   │   └── wechat.py
│   ├── models/                 # 数据模型
│   │   ├── __init__.py
│   │   ├── idea.py
│   │   ├── task.py
│   │   ├── config.py
│   │   └── sync_log.py
│   ├── storage/                # 存储层
│   │   ├── __init__.py
│   │   ├── database.py
│   │   └── file_storage.py
│   ├── utils/                  # 工具函数
│   │   ├── __init__.py
│   │   ├── crypto.py
│   │   └── logger.py
│   └── config.py               # 配置管理
├── tests/                      # 测试代码
│   ├── __init__.py
│   ├── test_cli.py
│   ├── test_services.py
│   └── test_api.py
├── capture-data/               # 数据目录
│   ├── ideas/
│   ├── tasks/
│   ├── templates/
│   ├── capture.db
│   └── logs/
├── docs/                       # 文档
│   ├── quickstart.md
│   ├── configuration.md
│   └── api.md
├── capture.yaml                # 配置文件
├── .env.example                # 环境变量示例
├── requirements.txt            # 依赖列表
├── setup.py                    # 安装配置
└── README.md                   # 项目说明
```

## 3. 核心流程设计

### 3.1 想法捕捉流程

```
用户输入 → 输入解析 → 结构化处理 → 存储 → 通知
```

**详细步骤**：
1. 接收用户输入（CLI/飞书/微信）
2. 解析输入内容，提取标题、描述、来源等
3. 调用 AI 进行结构化处理（可选）
4. 保存到数据库和文件系统
5. 发送通知（可选）

### 3.2 想法转化为任务流程

```
选择想法 → 确认转化 → 生成任务文档 → 创建任务记录 → 更新想法状态
```

**详细步骤**：
1. 用户选择待转化的想法
2. 确认转化细节（优先级、Agent 类型等）
3. 生成任务 Markdown 文档
4. 在数据库中创建任务记录
5. 更新想法状态为 "converted"

### 3.3 任务执行流程

```
选择任务 → 加载任务文档 → 调用 AI Agent → 监控执行 → 更新状态 → 通知用户
```

**详细步骤**：
1. 用户或系统选择待执行的任务
2. 从文件系统加载任务文档
3. 根据配置调用对应的 AI Agent
4. 实时监控执行进度和输出
5. 更新任务状态和执行日志
6. 发送执行结果通知

### 3.4 飞书多维表格同步流程

```
定时触发 → 查询变更任务 → 调用飞书 API → 记录同步日志
```

**详细步骤**：
1. 定时器触发同步任务
2. 查询自上次同步以来变更的任务
3. 调用飞书多维表格 API 更新记录
4. 记录同步结果到 sync_logs 表

## 4. 数据流设计

### 4.1 想法数据流

```
用户 → CLI/Bot → 想法捕捉服务 → 数据库 + 文件存储
```

### 4.2 任务执行数据流

```
用户 → CLI/API → 任务管理服务 → 执行引擎 → AI Agent → 文件存储
```

### 4.3 通知数据流

```
状态变更 → 通知服务 → 飞书/微信 API → 用户
```

### 4.4 同步数据流

```
定时器 → 同步服务 → 数据库 + 飞书 API
```

## 5. API 设计

### 5.1 RESTful API 规范

- 基础路径：`/api/v1`
- 认证方式：API Key（通过 `X-API-Key` Header）
- 响应格式：JSON
- 错误处理：统一错误格式

```json
{
  "error": {
    "code": "INVALID_REQUEST",
    "message": "请求参数错误",
    "details": "title 字段不能为空"
  }
}
```

### 5.2 关键 API 示例

#### 5.2.1 创建想法
```http
POST /api/v1/ideas
Content-Type: application/json
X-API-Key: xxx

{
  "title": "优化登录流程",
  "description": "当前登录流程太复杂，需要简化",
  "source": "产品讨论会",
  "channel": "feishu",
  "priority": "high"
}
```

#### 5.2.2 执行任务
```http
POST /api/v1/tasks/{task_id}/execute
Content-Type: application/json
X-API-Key: xxx

{
  "agent_type": "claude-code",
  "model": "claude-3-opus"
}
```

#### 5.2.3 获取看板数据
```http
GET /api/v1/kanban
X-API-Key: xxx
```

## 6. 安全设计

### 6.1 API Key 管理
- 用户首次运行时自动生成 API Key
- API Key 存储在本地配置文件中
- 支持手动重新生成

### 6.2 敏感信息加密
- 使用 `cryptography` 库进行加密
- 加密密钥存储在用户主目录（~/.capture/key）
- 敏感配置项（如飞书 App Secret）加密存储

### 6.3 飞书/微信验证
- 飞书：验证签名（signature、timestamp、nonce）
- 微信：验证 Token 和 EncodingAESKey

## 7. 性能优化

### 7.1 异步处理
- 使用 `asyncio` 进行异步 I/O
- 任务执行使用后台任务队列

### 7.2 缓存策略
- 配置信息缓存到内存
- 任务列表缓存（5 分钟有效期）

### 7.3 数据库优化
- 为常用查询字段添加索引
- 定期清理过期日志

## 8. 部署方案

### 8.1 本地部署
- 作为 Python 包安装
- CLI 直接运行
- API Server 作为后台服务运行

### 8.2 Docker 部署（可选）
```dockerfile
FROM python:3.10-slim

WORKDIR /app
COPY . .
RUN pip install -e .

EXPOSE 8000
CMD ["capture", "server", "start"]
```

## 9. 监控与日志

### 9.1 日志系统
- 使用 Python logging 模块
- 日志级别：DEBUG、INFO、WARNING、ERROR
- 日志文件按天轮转

### 9.2 性能监控
- 记录关键操作耗时
- 统计 API 调用频率
- 监控 AI Agent 执行时间

## 10. 测试策略

### 10.1 单元测试
- 覆盖核心服务逻辑
- 使用 pytest 框架
- Mock 外部依赖

### 10.2 集成测试
- 测试 API 端点
- 测试数据库操作
- 测试文件存储

### 10.3 端到端测试
- 测试完整的想法捕捉流程
- 测试任务执行流程
