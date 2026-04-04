# Local Task Center UI 严格评估

## 评估目标

评估 Local Task Center 是否应该同时实现 CLI/TUI、Web、Desktop 三种 UI 形态，并给出严格的优先级建议。

## 评估前提

- 产品目标是个人使用、轻量级、本地优先
- 当前已有 Go CLI/TUI 代码基础
- 当前尚未形成稳定的本地 API 层
- 当前最关键的问题不是“展示不足”，而是“任务中心工作流与执行跟踪尚未完全闭环”

## 评估维度

- 与个人场景的匹配度
- 与现有代码基础的复用度
- 研发复杂度
- 对执行观察能力的支持度
- 后续扩展到多端的一致性

## UI 形态逐项评估

### 1. CLI/TUI

#### 优势

- 与当前仓库能力最一致
- 与个人本地优先工作方式最匹配
- 适合快速录入、批量操作、状态推进
- 最接近 agent 执行现场，适合观察 terminal 与 worktree

#### 劣势

- 信息密度高时可视化能力有限
- 聚合筛选和统计能力不如 Web
- 对非终端型使用方式不够友好

#### 结论

必须优先实现，属于 P0 主界面。

### 2. Web

#### 优势

- 最适合展示 dashboard、聚合统计、跨任务筛选
- 最适合展示 agent / repo / stage 的多维观察
- 用户理解成本低，适合作为总览入口

#### 劣势

- 当前仓库还没有稳定 API 层
- 如果过早实现，会出现前端先行、核心不稳的问题
- 对 terminal 深交互支持天然较弱

#### 结论

应该实现，但应在 CLI/TUI 工作流稳定后进入，属于 P1。

### 3. Desktop

#### 优势

- 可以整合快捷捕捉、系统通知、本地 terminal、托盘入口
- 更适合做“个人工作台”

#### 劣势

- 研发成本最高
- 需要考虑打包、升级、权限、平台差异
- 如果没有稳定 API 和 Web 信息架构，Desktop 很容易变成重复实现

#### 结论

长期有价值，但不应作为近期目标，属于 P2。

## 严格结论

### 不建议

- 三端同时推进
- 先做 Desktop
- 在没有本地 API 之前直接铺开 Web 和 Desktop

### 建议

1. 先把 CLI/TUI 做成真正的 Task Center 主控台
2. 再为 Web 抽象本地 API 和只读/轻编辑能力
3. 最后再决定 Desktop 是 Electron、Tauri 还是完全不做

## 推荐职责分配

### CLI/TUI 负责

- 快速录入
- 阶段推进
- 分派 agent
- 执行观察
- review 操作

### Web 负责

- dashboard
- 聚合筛选
- 执行总览
- 历史回顾
- 轻量编辑

### Desktop 负责

- 快捷唤起
- 系统托盘
- 本地通知中心
- 集成 terminal 容器

## 推荐实施顺序

### Phase 1

- CLI/TUI 工作流闭环
- stage 看板
- dispatch 面板
- execution 观察摘要

### Phase 2

- 本地 API
- Web dashboard
- repo / agent / review 聚合视图

### Phase 3

- Desktop 增强壳层
- 快捷录入
- 通知与本地 terminal 集成

## 最终建议

如果只能做一个 UI，做 CLI/TUI。

如果做两个 UI，做 CLI/TUI + Web。

如果做完整产品路线，顺序必须是 CLI/TUI → Web → Desktop。
