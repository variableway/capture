# Storage Module Spec

## 目标

在保持本地优先的前提下，同时满足可读归档和快速查询。

## 结构

- Markdown：完整任务记录
- SQLite：索引与筛选

## 必须持久化的字段

- id
- title
- status
- stage
- priority
- source
- tags
- dispatch.agent
- dispatch.repository
- dispatch.assigned_at

## 设计原则

1. Markdown 是 source of truth
2. SQLite 是查询索引
3. 新字段优先保证 Markdown 完整性
4. SQLite 只保存高频查询字段

## 兼容性要求

- 新字段必须兼容旧任务文件
- SQLite 迁移必须支持已有本地数据库

## 本次实现

- 新增 stage 索引字段
- 新增 dispatch 索引摘要
- list 支持按 stage 过滤
