# Local Task Center UI 架构设计

## 目标

让 CLI/TUI、Web、Desktop 三类界面共享同一套任务中心核心能力，而不是三套各自演进的系统。

## 核心原则

1. 核心能力统一，UI 只是不同的呈现层
2. CLI/TUI 优先，Web 其次，Desktop 最后
3. 所有 UI 都围绕同一 Task Center Core 工作
4. Desktop 不直接绑定存储层，应复用本地 API

## 分层结构

### 1. Core Domain

统一负责：

- task
- stage
- status
- dispatch
- execution
- review

### 2. Application Services

统一负责：

- 创建任务
- 更新任务
- 分派任务
- 推进阶段
- 获取 dashboard 数据

### 3. Local API

作为 Web 与 Desktop 的共享入口，负责：

- 查询任务列表
- 查询 dashboard
- 更新 stage / status
- 提交 dispatch
- 查询 execution 状态

### 4. UI Surfaces

- CLI：命令驱动
- TUI：终端看板与详情界面
- Web：dashboard 与聚合界面
- Desktop：本地增强壳层

## UI 模块映射

### CLI/TUI

- Capture Entry
- Workflow Board
- Dispatch Panel
- Execution Monitor
- Review Queue

### Web

- Overview Dashboard
- Stage Board
- Agent View
- Repository View
- Review Center

### Desktop

- Quick Capture
- Local Dashboard Shell
- Notification Center
- Embedded Terminal Entry

## 数据流

1. 用户从 CLI/TUI、Web、Desktop 发起操作
2. 操作进入 Application Services
3. Services 调用 Store 与 Integration Layer
4. 更新结果回流到各 UI
5. Web/Desktop 通过 Local API 读取统一状态

## 为什么不让 Desktop 直接读 SQLite

- 会绕过统一业务规则
- 会让 Desktop 绑定本地存储细节
- 后续一旦存储或模型演进，Desktop 改动成本高

## 为什么 Web 不应直接先做前端

- 当前最关键的是核心工作流还在收敛
- 没有稳定 API 时，前端会反复重构
- 会导致“UI 很完整，核心流程不稳定”的假繁荣

## 推荐架构演进

### 当前

- CLI + TUI 直接调用 service

### 下一步

- 抽象 Local API
- Web 与 Desktop 全部经由 Local API

### 最终

- CLI/TUI 仍可直连 service
- Web/Desktop 标准化走 Local API
- Notification 与 Bot 作为外围入口/出口
