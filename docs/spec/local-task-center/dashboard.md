# Dashboard Module Spec

## 目标

把任务中心的状态变成可观察界面。

## 消费数据

- status
- stage
- priority
- tags
- dispatch.agent
- dispatch.repository
- execution.exec_status

## 视图需求

- 按 status 的 Kanban 视图
- 按 stage 的流程视图
- 按 agent 的执行视图
- 按 repo 的工作区视图

## 当前实现基础

- 现有 TUI 已支持 status 看板
- 本次增强展示 stage 与 agent 摘要
- 更完整的多端 UI 设计见 `ui-overview.md`、`ui-cli-tui.md`、`ui-web.md`、`ui-desktop.md`

## 下一步

- 增加 stage 维度切换
- 增加阻塞任务视图
- 增加 review 队列
- 增加今日活跃任务视图
