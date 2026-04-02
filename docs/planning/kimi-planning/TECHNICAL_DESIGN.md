# Capture Tool - Technical Design Document

## 1. 系统架构

### 1.1 整体架构图

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              Capture Tool                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │   CLI Layer │  │   TUI Layer │  │  Bot Layer  │  │   API Layer │         │
│  │   (Typer)   │  │  (Textual)  │  │  (Webhook)  │  │  (FastAPI)  │         │
│  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘         │
│         └─────────────────┴─────────────────┴─────────────────┘              │
│                                    │                                         │
│                         ┌──────────▼──────────┐                              │
│                         │    Service Layer    │                              │
│                         │  ┌───────────────┐  │                              │
│                         │  │ Task Service  │  │                              │
│                         │  ├───────────────┤  │                              │
│                         │  │Execution Svc  │  │                              │
│                         │  ├───────────────┤  │                              │
│                         │  │  Sync Service │  │                              │
│                         │  ├───────────────┤  │                              │
│                         │  │Notify Service │  │                              │
│                         │  └───────────────┘  │                              │
│                         └──────────┬──────────┘                              │
│                                    │                                         │
│                         ┌──────────▼──────────┐                              │
│                         │    Core Layer       │                              │
│                         │  ┌───────────────┐  │                              │
│                         │  │  Task Model   │  │                              │
│                         │  ├───────────────┤  │                              │
│                         │  │   Storage     │  │                              │
│                         │  ├───────────────┤  │                              │
│                         │  │   Config      │  │                              │
│                         │  ├───────────────┤  │                              │
│                         │  │  Event Bus    │  │                              │
│                         │  └───────────────┘  │                              │
│                         └──────────┬──────────┘                              │
│                                    │                                         │
│  ┌─────────────────────────────────┼─────────────────────────────────┐     │
│  │                         Data Layer                                 │     │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────────┐  │     │
│  │  │  SQLite  │  │Markdown  │  │  Config  │  │  External APIs   │  │     │
│  │  │  (Index) │  │  Files   │  │   YAML   │  │  Feishu/WeChat   │  │     │
│  │  └──────────┘  └──────────┘  └──────────┘  └──────────────────┘  │     │
│  └──────────────────────────────────────────────────────────────────┘     │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.2 模块职责

| 模块 | 职责 | 技术选型 |
|------|------|----------|
| CLI Layer | 命令行接口 | Typer |
| TUI Layer | 终端交互界面 | Textual |
| Bot Layer | Bot Webhook 处理 | FastAPI + Webhook |
| Service Layer | 业务逻辑 | Python Classes |
| Core Layer | 核心模型与存储 | Pydantic + SQLAlchemy |
| Data Layer | 数据持久化 | SQLite + File System |

## 2. 技术选型

### 2.1 核心依赖

```
Python >= 3.9
├── typer >= 0.9.0          # CLI 框架
├── textual >= 0.41.0       # TUI 框架
├── fastapi >= 0.104.0      # API/Webhook 框架
├── uvicorn >= 0.24.0       # ASGI 服务器
├── sqlalchemy >= 2.0.0     # ORM
├── pydantic >= 2.5.0       # 数据验证
├── pydantic-settings >= 2.0.0  # 配置管理
├── aiohttp >= 3.9.0        # 异步 HTTP 客户端
├── python-dotenv >= 1.0.0  # 环境变量
├── rich >= 13.7.0          # 终端美化
├── pyyaml >= 6.0.1         # YAML 解析
├── watchfiles >= 0.21.0    # 文件监控
└── pytest >= 7.4.0         # 测试框架
```

### 2.2 选型理由

| 技术 | 选型理由 |
|------|----------|
| **Typer** | 基于 Click，支持类型提示，自动生成帮助文档 |
| **Textual** | 功能强大的 TUI 框架，支持组件化开发，有完善的文档 |
| **FastAPI** | 高性能异步框架，自动生成 OpenAPI 文档 |
| **SQLAlchemy 2.0** | 现代 ORM，支持异步操作，类型安全 |
| **Pydantic v2** | 高性能数据验证，与 FastAPI 完美集成 |
| **Rich** | 美化终端输出，支持表格、进度条、Markdown 渲染 |

## 3. 目录结构

```
capture/
├── pyproject.toml              # 项目配置
├── README.md
├── src/
│   └── capture/
│       ├── __init__.py
│       ├── __main__.py         # 入口点
│       ├── cli.py              # CLI 入口
│       ├── config.py           # 配置管理
│       ├── constants.py        # 常量定义
│       │
│       ├── models/             # 数据模型
│       │   ├── __init__.py
│       │   ├── task.py         # 任务模型
│       │   ├── execution.py    # 执行记录模型
│       │   └── config.py       # 配置模型
│       │
│       ├── services/           # 业务服务
│       │   ├── __init__.py
│       │   ├── task_service.py
│       │   ├── execution_service.py
│       │   ├── sync_service.py
│       │   ├── notification_service.py
│       │   └── bot_service.py
│       │
│       ├── storage/            # 存储层
│       │   ├── __init__.py
│       │   ├── database.py     # SQLite 操作
│       │   ├── file_storage.py # 文件存储
│       │   └── repository.py   # 仓储模式
│       │
│       ├── agents/             # AI Agent 集成
│       │   ├── __init__.py
│       │   ├── base.py         # Agent 基类
│       │   ├── claude_code.py  # Claude Code 集成
│       │   ├── codex.py        # OpenAI Codex 集成
│       │   └── kimi_cli.py     # Kimi CLI 集成
│       │
│       ├── bots/               # Bot 适配器
│       │   ├── __init__.py
│       │   ├── base.py         # Bot 基类
│       │   ├── feishu.py       # 飞书 Bot
│       │   └── wechat.py       # 微信 Bot
│       │
│       ├── tui/                # TUI 界面
│       │   ├── __init__.py
│       │   ├── app.py          # TUI 应用
│       │   ├── kanban.py       # 看板组件
│       │   ├── task_card.py    # 任务卡片
│       │   └── widgets.py      # 通用组件
│       │
│       ├── sync/               # 同步模块
│       │   ├── __init__.py
│       │   ├── base.py
│       │   └── feishu_bitable.py
│       │
│       ├── utils/              # 工具函数
│       │   ├── __init__.py
│       │   ├── id_generator.py # ID 生成器
│       │   ├── validators.py   # 验证器
│       │   └── helpers.py      # 辅助函数
│       │
│       └── web/                # Web 服务
│           ├── __init__.py
│           ├── server.py       # FastAPI 应用
│           └── routes.py       # 路由
│
├── tests/                      # 测试
│   ├── __init__.py
│   ├── test_task_service.py
│   ├── test_execution.py
│   └── conftest.py
│
└── docs/                       # 文档
    └── api.md
```

## 4. 核心类设计

### 4.1 任务模型

```python
from datetime import datetime
from enum import Enum
from typing import List, Optional, Dict, Any
from pydantic import BaseModel, Field

class TaskStatus(str, Enum):
    TODO = "todo"
    IN_PROGRESS = "in_progress"
    DONE = "done"
    CANCELLED = "cancelled"
    ARCHIVED = "archived"

class TaskPriority(str, Enum):
    HIGH = "high"
    MEDIUM = "medium"
    LOW = "low"

class TaskContext(BaseModel):
    trigger: Optional[str] = None
    location: Optional[str] = None
    related_to: Optional[str] = None
    source_message: Optional[Dict[str, Any]] = None

class TaskExecution(BaseModel):
    agent: Optional[str] = None
    model: Optional[str] = None
    result: Optional[str] = None
    logs: List[str] = Field(default_factory=list)
    started_at: Optional[datetime] = None
    completed_at: Optional[datetime] = None
    status: Optional[str] = None  # pending, running, success, failed

class TaskSync(BaseModel):
    feishu_bitable_record_id: Optional[str] = None
    last_sync: Optional[datetime] = None

class Task(BaseModel):
    id: str = Field(..., description="Task ID (TASK-XXXXX)")
    title: str
    description: Optional[str] = None
    status: TaskStatus = TaskStatus.TODO
    priority: TaskPriority = TaskPriority.MEDIUM
    tags: List[str] = Field(default_factory=list)
    created_at: datetime = Field(default_factory=datetime.utcnow)
    updated_at: datetime = Field(default_factory=datetime.utcnow)
    source: str = "cli"  # cli, tui, feishu_bot, wechat_bot
    context: TaskContext = Field(default_factory=TaskContext)
    execution: TaskExecution = Field(default_factory=TaskExecution)
    sync: TaskSync = Field(default_factory=TaskSync)
    file_path: Optional[str] = None
    
    class Config:
        from_attributes = True
```

### 4.2 服务接口

```python
from abc import ABC, abstractmethod
from typing import List, Optional

class ITaskService(ABC):
    @abstractmethod
    async def create(self, title: str, **kwargs) -> Task:
        pass
    
    @abstractmethod
    async def get(self, task_id: str) -> Optional[Task]:
        pass
    
    @abstractmethod
    async def update(self, task_id: str, **kwargs) -> Task:
        pass
    
    @abstractmethod
    async def delete(self, task_id: str) -> bool:
        pass
    
    @abstractmethod
    async def list(self, status: Optional[TaskStatus] = None, 
                   tags: Optional[List[str]] = None) -> List[Task]:
        pass

class IExecutionService(ABC):
    @abstractmethod
    async def execute(self, task_id: str, agent: Optional[str] = None) -> ExecutionResult:
        pass
    
    @abstractmethod
    async def get_logs(self, task_id: str) -> List[ExecutionLog]:
        pass
    
    @abstractmethod
    async def cancel(self, execution_id: str) -> bool:
        pass
```

### 4.3 Agent 基类

```python
from abc import ABC, abstractmethod
from dataclasses import dataclass
from typing import Dict, Any, AsyncIterator

@dataclass
class ExecutionResult:
    success: bool
    output: str
    error: Optional[str] = None
    artifacts: Dict[str, Any] = None

class BaseAgent(ABC):
    name: str
    supported_models: List[str]
    
    def __init__(self, config: Dict[str, Any]):
        self.config = config
    
    @abstractmethod
    async def execute(self, task: Task, model: Optional[str] = None) -> ExecutionResult:
        """执行任务，返回执行结果"""
        pass
    
    @abstractmethod
    async def execute_stream(self, task: Task, model: Optional[str] = None) -> AsyncIterator[str]:
        """流式执行，返回输出流"""
        pass
    
    @abstractmethod
    def validate_task(self, task: Task) -> bool:
        """验证任务是否适合该 Agent"""
        pass
```

## 5. 数据流设计

### 5.1 任务创建流程

```
┌─────────┐    ┌──────────┐    ┌──────────────┐    ┌──────────────┐
│  User   │───▶│ CLI/Bot  │───▶│ TaskService  │───▶│   Validate   │
└─────────┘    └──────────┘    └──────────────┘    └──────────────┘
                                                        │
                         ┌──────────────────────────────┘
                         ▼
┌─────────┐    ┌──────────────┐    ┌──────────────┐    ┌──────────┐
│   ID    │◀───│  Generate    │◀───│  Build Task  │◀───│ Sanitize │
│Generator│    │   Metadata   │    │    Object    │    │   Input  │
└────┬────┘    └──────────────┘    └──────────────┘    └──────────┘
     │
     ▼
┌──────────────────────────────────────────────────────────────────┐
│                         Parallel Save                             │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────────────┐  │
│  │Write to File│    │Insert to DB │    │Sync to Feishu (opt) │  │
│  │  (Markdown) │    │  (SQLite)   │    │                     │  │
│  └─────────────┘    └─────────────┘    └─────────────────────┘  │
└──────────────────────────────────────────────────────────────────┘
     │
     ▼
┌──────────────┐    ┌──────────┐
│   Return     │───▶│  Notify  │
│  Task ID     │    │  User    │
└──────────────┘    └──────────┘
```

### 5.2 任务执行流程

```
┌─────────┐    ┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│  User   │───▶│Execute Command│───▶│ Load Task   │───▶│Check Permission│
└─────────┘    └──────────────┘    └──────────────┘    └──────────────┘
                                                            │
                         ┌──────────────────────────────────┘
                         ▼
┌────────────────┐    ┌──────────────┐    ┌──────────────┐    ┌──────────┐
│  Update Status │◀───│   Execute    │◀───│ Select Agent │◀───│ Validate │
│  to Running    │    │    Agent     │    │  by Config   │    │   Task   │
└───────┬────────┘    └──────────────┘    └──────────────┘    └──────────┘
        │
        ▼
┌──────────────────────────────────────────────────────────────────────┐
│                          Execution Loop                               │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────────┐   │
│  │Stream    │───▶│Capture   │───▶│Log to    │───▶│Check Timeout │   │
│  │Output    │    │Output    │    │File & DB │    │& Cancel?     │   │
│  └──────────┘    └──────────┘    └──────────┘    └──────────────┘   │
└──────────────────────────────────────────────────────────────────────┘
        │
        ▼
┌──────────────────────────────────────────────────────────────────────┐
│                         Completion Handling                           │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────────────┐   │
│  │Update Status │    │  Write Log   │    │   Notify User/Bot    │   │
│  │(done/failed) │    │  (Markdown)  │    │   with Summary       │   │
│  └──────────────┘    └──────────────┘    └──────────────────────┘   │
└──────────────────────────────────────────────────────────────────────┘
```

### 5.3 同步流程 (飞书多维表格)

```
┌──────────┐    ┌──────────────┐    ┌─────────────────────────────┐
│  Trigger │───▶│ SyncService  │───▶│ 1. Query local changes      │
│(Manual/  │    │              │    │ 2. Query Feishu changes     │
│ Cron)    │    │              │    │ 3. Build diff               │
└──────────┘    └──────────────┘    └─────────────────────────────┘
                                              │
                    ┌─────────────────────────┴────────────────────┐
                    ▼                                            ▼
┌──────────────────────────┐                        ┌──────────────────────────┐
│     Push to Feishu       │                        │     Pull from Feishu     │
│  (local changes exist)   │                        │  (remote changes exist)  │
└──────────┬───────────────┘                        └──────────┬───────────────┘
           │                                                   │
           ▼                                                   ▼
┌──────────────────────────┐                        ┌──────────────────────────┐
│ - Create new records     │                        │ - Create local tasks     │
│ - Update existing records│                        │ - Update local tasks     │
│ - Delete records (soft)  │                        │ - Handle conflicts       │
└──────────┬───────────────┘                        └──────────┬───────────────┘
           │                                                   │
           └───────────────────────┬───────────────────────────┘
                                   ▼
                    ┌──────────────────────────┐
                    │  Conflict Resolution     │
                    │  (manual/auto rules)     │
                    └──────────┬───────────────┘
                               │
                               ▼
                    ┌──────────────────────────┐
                    │  Update sync timestamps  │
                    │  Record sync log         │
                    └──────────────────────────┘
```

## 6. 关键技术点

### 6.1 文件与数据库双写一致性

使用 **Event Sourcing + Write-Ahead Log** 模式：

```python
class StorageManager:
    def __init__(self):
        self.db = SQLiteStorage()
        self.file_store = FileStorage()
        self.wal = WriteAheadLog()
    
    async def save_task(self, task: Task):
        # 1. 写入 WAL
        entry = self.wal.append("SAVE_TASK", task.model_dump())
        
        try:
            # 2. 并行写入文件和数据库
            await asyncio.gather(
                self.file_store.save(task),
                self.db.save(task)
            )
            
            # 3. 标记 WAL 完成
            entry.mark_complete()
            
        except Exception as e:
            # 4. 失败时从 WAL 恢复或重试
            await self._recover_from_wal(entry)
            raise
```

### 6.2 Bot 消息解析

使用简单的意图识别 + 参数提取：

```python
class MessageParser:
    INTENTS = {
        r"记录|添加|新建|create|add": "create_task",
        r"执行|运行|execute|run": "execute_task",
        r"列出|查看|list|show": "list_tasks",
        r"删除|remove|delete": "delete_task",
        r"更新|修改|update|edit": "update_task",
    }
    
    def parse(self, message: str) -> ParsedIntent:
        # 1. 识别意图
        intent = self._detect_intent(message)
        
        # 2. 提取参数
        params = self._extract_params(message)
        
        return ParsedIntent(intent=intent, params=params)
    
    def _extract_params(self, message: str) -> Dict[str, Any]:
        params = {}
        
        # 标签提取: #标签名
        params["tags"] = re.findall(r'#(\w+)', message)
        
        # 优先级提取: 优先级：高/中/低
        priority_match = re.search(r'优先级[：:]\s*(高|中|低)', message)
        if priority_match:
            params["priority"] = priority_match.group(1)
        
        # 任务 ID 提取: TASK-XXXXX
        task_id_match = re.search(r'(TASK-\d+)', message)
        if task_id_match:
            params["task_id"] = task_id_match.group(1)
        
        return params
```

### 6.3 TUI 看板渲染

```python
from textual.app import App, ComposeResult
from textual.widgets import DataTable, Header, Footer, Static
from textual.containers import Horizontal, Vertical

class KanbanBoard(App):
    CSS = """
    .column { width: 33%; height: 100%; border: solid green; }
    .task-card { height: auto; margin: 1; padding: 1; }
    """
    
    def compose(self) -> ComposeResult:
        yield Header()
        yield Horizontal(
            Vertical(Static("TODO"), id="col-todo", classes="column"),
            Vertical(Static("IN PROGRESS"), id="col-inprogress", classes="column"),
            Vertical(Static("DONE"), id="col-done", classes="column"),
        )
        yield Footer()
    
    def on_mount(self):
        self.load_tasks()
    
    async def load_tasks(self):
        tasks = await self.task_service.list()
        for task in tasks:
            card = TaskCard(task)
            self.query_one(f"#col-{task.status}").mount(card)
```

### 6.4 Agent 执行隔离

使用子进程隔离执行：

```python
import asyncio
import tempfile
import os

class AgentRunner:
    async def run(self, agent: BaseAgent, task: Task) -> ExecutionResult:
        # 1. 创建临时工作目录
        with tempfile.TemporaryDirectory() as work_dir:
            # 2. 写入任务描述为文件
            task_file = os.path.join(work_dir, "TASK.md")
            await self._write_task_file(task, task_file)
            
            # 3. 构建命令
            cmd = self._build_command(agent, task_file, work_dir)
            
            # 4. 在子进程中执行
            proc = await asyncio.create_subprocess_shell(
                cmd,
                stdout=asyncio.subprocess.PIPE,
                stderr=asyncio.subprocess.PIPE,
                cwd=work_dir,
                limit=1024 * 1024  # 1MB buffer
            )
            
            # 5. 流式读取输出
            stdout_chunks = []
            async for chunk in self._read_stream(proc.stdout):
                stdout_chunks.append(chunk)
                await self._emit_progress(chunk)
            
            # 6. 等待完成
            returncode = await proc.wait()
            
            return ExecutionResult(
                success=returncode == 0,
                output="".join(stdout_chunks),
                error=await proc.stderr.read() if returncode != 0 else None
            )
```

## 7. 配置管理

### 7.1 配置加载优先级 (从高到低)

1. 命令行参数
2. 环境变量 (`CAPTURE_*`)
3. 用户配置文件 (`~/.capture/config.yaml`)
4. 项目配置文件 (`./.capture/config.yaml`)
5. 默认配置

### 7.2 配置热重载

```python
from watchfiles import watch

class ConfigManager:
    def __init__(self, config_path: str):
        self.config_path = config_path
        self.config = self._load()
        self._start_watcher()
    
    def _start_watcher(self):
        async def watch_config():
            async for changes in watch(self.config_path):
                self.config = self._load()
                self._emit_config_changed()
        
        asyncio.create_task(watch_config())
```

## 8. 测试策略

### 8.1 测试金字塔

```
                    ┌─────────┐
                    │  E2E   │  (5%)  - 关键流程端到端测试
                    │ Tests  │
                   ┌┴─────────┴┐
                   │ Integration│ (25%) - 服务层集成测试
                   │   Tests    │
                  ┌┴────────────┴┐
                  │    Unit      │ (70%) - 模型、工具函数单元测试
                  │    Tests     │
                  └──────────────┘
```

### 8.2 关键测试用例

| 模块 | 测试类型 | 测试内容 |
|------|----------|----------|
| TaskService | 单元测试 | CRUD 操作、状态流转 |
| Storage | 单元测试 | 文件/数据库双写一致性 |
| ExecutionService | 集成测试 | Agent 调用、超时处理 |
| Bot Parser | 单元测试 | 意图识别、参数提取 |
| SyncService | 集成测试 | 双向同步、冲突处理 |
| TUI | E2E | 看板交互、状态更新 |

## 9. 部署方案

### 9.1 开发模式

```bash
# 克隆仓库
git clone <repo>
cd capture

# 安装开发依赖
pip install -e ".[dev]"

# 初始化
capture init

# 运行 TUI
capture kanban

# 启动 Bot 服务
capture bot serve --port 8080
```

### 9.2 生产部署 (Bot 服务)

```bash
# 使用 Docker
docker build -t capture-bot .
docker run -d \
  -p 8080:8080 \
  -v ~/.capture:/data \
  -e FEISHU_APP_ID=xxx \
  -e FEISHU_APP_SECRET=xxx \
  capture-bot

# 或使用 systemd
systemctl enable capture-bot
systemctl start capture-bot
```

### 9.3 飞书 Bot 配置

1. 在飞书开放平台创建企业自建应用
2. 启用机器人功能
3. 配置事件订阅 Webhook: `https://your-domain.com/webhook/feishu`
4. 配置权限: `im:chat:readonly`, `im:message:send`
5. 发布应用

---

**文档版本**: 1.0  
**最后更新**: 2024-04-02
