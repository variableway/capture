# Assignment Module Spec

## 目标

把任务与具体的 AI Agent 执行上下文绑定。

## 输入

- task id
- agent
- model
- repository
- worktree
- terminal session

## 输出

- dispatch 元数据写入任务
- stage 自动进入 `dispatch`

## 数据结构

- dispatch.agent
- dispatch.model
- dispatch.repository
- dispatch.worktree
- dispatch.terminal_session
- dispatch.assigned_at

## 接口

- `capture assign <id> --agent <name> --model <model> --repo <path> --worktree <path> --terminal <session>`

## 规则

1. agent 为必填
2. repo/worktree/terminal 为推荐字段
3. 允许重复分派，最新一次覆盖旧值

## MVP

- 只记录分派元数据
- 不直接启动执行器
