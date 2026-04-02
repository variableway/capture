# Capture Tool - 规划文档

本文档目录包含 Capture 工具的完整规划，基于 PRD/overview.md 中的需求设计。

## 文档结构

```
kimi-planning/
├── README.md                # 本文档 - 规划概览
├── PRD.md                   # 产品需求文档
├── TECHNICAL_DESIGN.md      # 技术设计文档
└── TASKS.md                 # 任务分解与实施计划
```

## 快速导航

### 1. [产品需求文档 (PRD.md)](./PRD.md)
包含完整的产品需求定义：
- **功能需求**: 6 大核心模块（捕获、任务管理、看板、执行引擎、通知、同步）
- **非功能需求**: 性能、可靠性、安全、兼容性、可扩展性
- **接口设计**: CLI 命令、Bot Webhook、TUI 界面、内部 API
- **数据库设计**: SQLite 表结构、配置结构、文件存储格式

### 2. [技术设计文档 (TECHNICAL_DESIGN.md)](./TECHNICAL_DESIGN.md)
包含详细的技术实现方案：
- **系统架构**: 分层架构图与模块职责
- **技术选型**: Python 技术栈及选型理由
- **目录结构**: 项目代码组织方式
- **核心类设计**: Task 模型、Service 接口、Agent 基类
- **数据流设计**: 任务创建、执行、同步的完整流程
- **关键技术点**: 双写一致性、消息解析、TUI 渲染、执行隔离

### 3. [任务分解文档 (TASKS.md)](./TASKS.md)
包含详细的实施计划：
- **83 个具体任务**，按模块和优先级划分
- **6 个阶段**的迭代计划 (MVP → 执行引擎 → TUI → Bot → 通知/同步 → 优化)
- **预估工时**: P0 任务约 50 小时，全部任务约 166 小时
- **依赖关系**: 明确任务间的先后依赖

## 核心功能一览

| 功能模块 | 描述 | 优先级 |
|----------|------|--------|
| CLI 任务管理 | 命令行添加、查看、编辑、删除任务 | P0 |
| 本地文件存储 | 任务以 Markdown + Frontmatter 形式存储 | P0 |
| AI Agent 执行 | 集成 Claude Code/Codex 执行任务 | P0 |
| TUI 看板 | 终端可视化任务看板，支持拖拽 | P1 |
| 飞书 Bot | 通过飞书 Bot 记录和执行任务 | P1 |
| 消息通知 | 任务状态变更和执行结果通知 | P1 |
| 飞书同步 | 与飞书多维表格双向同步 | P1 |

## 技术栈

- **CLI**: [Typer](https://typer.tiangolo.com/) - 基于 Python 类型提示的 CLI 框架
- **TUI**: [Textual](https://textual.textualize.io/) - Python TUI 框架
- **Web**: [FastAPI](https://fastapi.tiangolo.com/) - 高性能异步 Web 框架
- **ORM**: [SQLAlchemy 2.0](https://docs.sqlalchemy.org/) - 现代异步 ORM
- **验证**: [Pydantic v2](https://docs.pydantic.dev/) - 数据验证
- **终端**: [Rich](https://rich.readthedocs.io/) - 终端美化

## 快速开始 (开发计划)

### Phase 1: MVP (Week 1-2)
实现基础任务管理：
```bash
# 初始化
capture init

# 添加任务
capture add "优化项目构建脚本" -d "减少构建时间" -t 优化,构建 -p high

# 查看任务
capture list
capture show TASK-001

# 编辑任务
capture edit TASK-001

# 删除任务
capture delete TASK-001
```

### Phase 2: 执行引擎 (Week 2-3)
集成 AI Agent 执行任务：
```bash
# 执行任务
capture execute TASK-001

# 查看执行日志
capture logs TASK-001
```

### Phase 3: TUI 看板 (Week 3-4)
可视化任务管理：
```bash
# 启动看板
capture kanban
```

### Phase 4: Bot 集成 (Week 4-5)
飞书 Bot 交互：
```bash
# 启动 Bot 服务
capture bot serve --port 8080
```

飞书内使用：
```
@CaptureBot 记录：优化项目构建脚本
标签：#优化 #构建
优先级：高
```

### Phase 5: 同步与通知 (Week 5-6)
```bash
# 手动同步
capture sync
```

## 项目状态追踪

使用 [TASKS.md](./TASKS.md) 追踪开发进度：

- ⬜ 未开始
- 🔄 进行中
- ✅ 已完成

建议每周更新一次任务状态。

## 参考资源

- [Typer 文档](https://typer.tiangolo.com/)
- [Textual 文档](https://textual.textualize.io/)
- [FastAPI 文档](https://fastapi.tiangolo.com/)
- [飞书开放平台](https://open.feishu.cn/)
- [Claude Code](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview)

---

**规划版本**: 1.0  
**创建日期**: 2024-04-02
