# CLI/TUI UI Spec

## 定位

CLI/TUI 是 Local Task Center 的主控界面，承担从任务录入到执行观察的核心链路。

## 核心模块

### 1. Capture Entry

- 快速录入任务
- 设置 priority、stage、tags
- 支持从 terminal 快速进入 inbox

### 2. Task Inbox

- 查看新输入任务
- 按来源过滤
- 快速推进到 mindstorm / analysis

### 3. Workflow Board

- 按 status 查看
- 按 stage 查看
- 支持快速切换任务阶段

### 4. Dispatch Panel

- 绑定 agent
- 绑定 model
- 绑定 repo / worktree / terminal

### 5. Execution Monitor

- 查看执行状态
- 查看最近活动
- 查看执行摘要

### 6. Review Queue

- 查看待确认任务
- 确认结果是否回流
- 推进到 done 或重新进入 planning

## MVP 范围

- `capture add`
- `capture list --stage`
- `capture stage`
- `capture assign`
- TUI 展示 stage 与 agent 摘要

## 后续增强

- TUI stage 视图切换
- repo/agent 聚合页
- terminal 快捷跳转
