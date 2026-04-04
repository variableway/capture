# Local Task Center 可行性分析

## 目标

把 Capture 从“任务收集工具”扩展为“个人本地任务收口中心”，覆盖以下链路：

1. 收集任务：Terminal、飞书 Bot、后续 Web/Desktop
2. 组织任务：从 idea 进入 mindstorm、analysis、planning、prd、tasks
3. 下发任务：把任务分派给不同 AI Agent 和代码仓库
4. 跟踪执行：知道任务在哪个 repo、哪个 terminal、哪个 agent 正在执行
5. 汇总反馈：任务完成后回流到统一看板

## 现状基础

- 已有 CLI，可作为本地控制面
- 已有 TUI，可作为第一版轻量级看板
- 已有 Markdown + SQLite 双写存储，可继续承载工作流数据
- 已有飞书 Bot 输入通道，可继续作为远程采集入口
- 已有 Bitable 同步，可作为外部可视化补充

## 可行性结论

结论：可行，且适合按“轻量个人版 → 扩展多端版”分阶段推进。

原因：

1. 当前仓库已经具备任务对象、存储层、CLI/TUI、Bot、同步等核心基础
2. 本地优先架构天然适合个人使用，复杂度和运维成本低
3. 工作流阶段和分派元数据可以直接叠加到现有 Task 模型
4. 真正复杂的部分不是存储，而是执行编排和可视化观察，需要拆阶段实现

## 主要风险

### 1. 执行编排复杂度高

- 多个 Agent 运行方式不同
- 不同仓库的执行上下文不同
- terminal 会话、日志、心跳、失败重试都需要抽象

结论：第一阶段只做“记录分派”和“执行状态跟踪”，不直接做统一执行器。

### 2. 多端同时推进成本高

- CLI/TUI、Web、Desktop 的交互需求不同
- 轻量个人产品不适合一开始就做完整多端

结论：先以 CLI/TUI 为主，Web 做观察面，Desktop 后置。

### 3. 状态模型容易混乱

- 业务状态和流程阶段不是同一层概念
- 如果混在一起，会导致看板含义不清晰

结论：保留现有 status，新增独立 stage。

## 建议路线

### Phase 1：Task Center 基础层

- 引入 stage 工作流
- 引入 dispatch 分派信息
- CLI 支持 stage 与 assign
- TUI 展示阶段与分派摘要

### Phase 2：执行观察层

- 增加 execution session、日志索引、心跳
- 支持 repo/worktree/terminal 维度追踪
- 增加 dashboard 数据视图

### Phase 3：多端接入层

- Web 只做只读观察和简单操作
- Desktop 作为本地增强终端入口
- Bot 继续承担远程输入与提醒

## MVP 范围

本次建议落地范围：

- 在现有任务模型中增加 stage
- 支持为任务记录 agent/model/repo/worktree/terminal
- 输出完整的分析、架构、规划与模块 spec 文档
- 为后续 execution/dashboard/web/desktop 预留结构
