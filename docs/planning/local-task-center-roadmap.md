# Local Task Center 规划

## 阶段一：基础工作流

目标：让 Capture 真正具备任务中心的基本流程表达能力。

- 新增 stage 字段
- 新增 assign 命令
- 新增 stage 命令
- list/show/TUI 展示阶段与分派摘要
- 输出分析、架构与模块 spec

## 阶段二：执行跟踪

目标：从“记录分派”升级为“可观察执行”。

- 增加 execution session 标识
- 增加执行日志索引
- 增加心跳与最后活动时间
- 支持 terminal/repo/worktree 维度查询

## 阶段三：Dashboard

目标：形成真正的任务中心界面。

- TUI 增加基于 stage 的视图
- 增加今日活跃任务视图
- 增加按 agent / repo 聚合视图
- 增加失败任务与阻塞任务视图

## 阶段四：多端统一

目标：让 Web 与 Desktop 共享同一任务中心能力。

- 增加本地 API
- Web 端只读与轻编辑
- Desktop 端增强 terminal 入口
- Bot 与 Web 状态保持一致

## 建议优先级

### P0

- stage 工作流
- dispatch 元数据
- CLI 能力补齐
- 存储兼容性

### P1

- execution 跟踪
- TUI 增强
- Dashboard 只读页

### P2

- Desktop
- 高级调度
- 自动重试与回滚
