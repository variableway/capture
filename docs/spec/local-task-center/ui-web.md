# Web UI Spec

## 定位

Web 是 Local Task Center 的观察与管理界面，不承担最深的 terminal 交互。

## 核心模块

### 1. Overview Dashboard

- 今日活跃任务
- 按 stage 分布
- 按 agent 分布
- 按 repo 分布

### 2. Stage Board

- 可视化流程板
- 快速定位阻塞阶段
- 查看 review 队列

### 3. Agent View

- 某个 agent 当前执行哪些任务
- 每个任务在哪个 repo / terminal
- 执行状态摘要

### 4. Repository View

- 每个 repo 当前承载的任务
- 哪些任务在同一 worktree
- 哪些任务长期阻塞

### 5. Task Detail

- 查看完整任务信息
- 轻量编辑 stage / status / dispatch

## MVP 范围

- dashboard 只读
- task detail 轻编辑
- 按 stage、agent、repo 过滤

## 依赖前提

- 必须先有稳定 Local API
- 必须先稳定 Task Center 核心模型

## 不建议提前做的内容

- 前端先行的复杂拖拽
- 脱离本地 API 的直接存储访问
- 深度 terminal 嵌入
