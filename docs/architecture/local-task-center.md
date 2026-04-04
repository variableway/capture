# Local Task Center 架构设计

## 架构目标

构建一个轻量级、个人本地优先的任务中心，统一承接：

- idea 收集
- 任务分析与规划
- AI Agent 分派
- 执行跟踪
- 结果回流

## 总体分层

### 1. Input Ports

负责接收任务输入：

- CLI
- TUI
- 飞书 Bot
- 后续 Web
- 后续 Desktop

### 2. Task Center Core

核心领域层，负责：

- Task 生命周期
- stage 流转
- status 管理
- dispatch 分派
- execution 记录

### 3. Persistence

负责本地存储：

- Markdown：完整记录与可读归档
- SQLite：索引、过滤、看板查询

### 4. Integration Layer

负责外部集成：

- 飞书 Bot
- 飞书 Bitable
- 后续 Agent Runtime Adapter

### 5. Observability Layer

负责可视化与反馈：

- TUI 看板
- 后续 Web Dashboard
- 通知系统

## 核心对象

### Task

统一任务对象，包含：

- 基础信息：title、description、tags、priority
- 业务状态：todo、in_progress、done、cancelled、archived
- 流程阶段：inbox、mindstorm、analysis、planning、prd、tasks、dispatch、execution、review
- 分派信息：agent、model、repo、worktree、terminal
- 执行信息：执行状态、开始时间、结束时间、结果摘要

### Stage 与 Status 的边界

- stage 表示任务当前处于哪一个流程阶段
- status 表示任务当前的执行状态

示例：

- 一个任务可以处于 `stage=analysis` 且 `status=todo`
- 一个任务可以处于 `stage=execution` 且 `status=in_progress`
- 一个任务可以处于 `stage=review` 且 `status=done`

## 模块划分

### 1. Intake Module

负责收集输入并转成统一 Task。

### 2. Workflow Module

负责 stage 与 status 的规则与切换。

### 3. Dispatch Module

负责把任务下发到指定 agent 与仓库执行环境。

### 4. Execution Tracking Module

负责记录 terminal、工作目录、开始/结束时间、结果。

### 5. Storage Module

负责 Markdown 与 SQLite 双写。

### 6. Dashboard Module

负责给 TUI/Web 输出可观察视图。

## 数据流

1. 用户从 CLI / Bot 输入 idea
2. Intake Module 创建 Task，默认进入 `stage=inbox`
3. 用户或系统逐步推进到 `mindstorm/analysis/planning/prd/tasks`
4. Dispatch Module 记录 agent、repo、terminal，阶段进入 `dispatch`
5. Execution Tracking Module 标记 `execution`
6. 完成后进入 `review`
7. 确认完成后 status 进入 `done`

## 部署建议

### 当前阶段

- 单机本地运行
- 本地文件 + SQLite
- 飞书作为远程输入通道

### 下一阶段

- 增加本地 HTTP API 给 Web/Desktop 复用
- Web 只读或轻编辑
- Desktop 复用本地 API 与 terminal 能力

## 本次实现对应范围

本次实现先覆盖：

- stage 模型
- dispatch 元数据
- CLI 管理命令
- 文档化模块设计

不在本次实现范围：

- 统一 Agent 执行器
- Web Dashboard
- Desktop 客户端
- 多机调度
