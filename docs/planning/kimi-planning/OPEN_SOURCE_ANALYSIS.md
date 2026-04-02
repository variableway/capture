# Capture Tool - 开源组件参考分析报告

本文档对 Capture 工具各模块可参考的开源项目进行深入分析，每个模块精选 3-5 个代表性项目，涵盖架构设计、功能实现和技术选型等方面。

---

## 目录

1. [任务管理模块](#1-任务管理模块)
2. [CLI/TUI 界面模块](#2-clitui-界面模块)
3. [Web/Desktop UI 看板模块](#3-webdesktop-ui-看板模块)
4. [Vibe Coding 看板模块（AI Agent 专用）](#4-vibe-coding-看板模块ai-agent-专用)
5. [AI Agent 执行模块](#5-ai-agent-执行模块)
6. [Bot 集成模块](#6-bot-集成模块)
7. [通知系统模块](#7-通知系统模块)
8. [数据同步模块](#8-数据同步模块)
9. [架构设计参考](#9-架构设计参考)
10. [总结与建议](#10-总结与建议)

---

## 1. 任务管理模块

### 1.1 Taskwarrior

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/GothenburgBitFactory/taskwarrior |
| **开发语言** | C++ |
| **Stars** | 8.5k+ |
| **许可证** | MIT |

**项目简介**
Taskwarrior 是一个功能强大的命令行任务管理工具，自 2008 年发布以来，已成为 CLI 任务管理领域的标杆项目。它以简单的工作流、丰富的功能和极高的性能著称。

**核心功能**
- 自然语言风格的命令设计 (`task add`, `task list`, `task done`)
- 强大的过滤和查询系统 (支持日期、标签、项目、优先级等)
- 虚拟标签系统 (自动计算的 `TODAY`, `OVERDUE`, `BLOCKED` 等)
- 任务依赖关系管理
- 用户自定义属性和报告
- Hooks API 扩展机制
- Taskserver 多设备同步

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **命令设计** | 使用自然语言风格的命令，降低学习成本 |
| **过滤系统** | 强大的表达式过滤，如 `task project:Work priority:H due.before:today` |
| **虚拟标签** | 自动计算任务状态标签，减少用户手动维护 |
| **数据存储** | 纯文本文件存储，便于版本控制和同步 |
| **扩展机制** | Hooks API 允许在任务生命周期中插入自定义脚本 |

**与 Capture 的关联**
- 参考其命令设计风格，设计 `capture add/list/done` 等直观命令
- 借鉴其过滤表达式，实现强大的任务搜索功能
- 学习其 Hooks 机制，设计事件驱动的插件系统

---

### 1.2 dstask

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/naggie/dstask |
| **开发语言** | Go |
| **Stars** | 1.2k+ |
| **许可证** | MIT |

**项目简介**
dstask 是受 Taskwarrior 启发的个人任务跟踪器，但使用 Git 进行数据同步而非专用协议。它的设计理念是帮助用户专注，特别适合喜欢终端和 Git 的开发者。

**核心功能**
- Git 驱动的数据同步和版本控制
- 每个任务的 Markdown 笔记支持 (`note` 命令)
- 强大的上下文系统（自动应用过滤/标签）
- 单静态二进制文件分发
- zsh/bash 自动补全
- 内置导入工具（GitHub Issues、Taskwarrior）

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **Git 同步** | 使用 Git 进行数据同步，无需专用服务器 |
| **Markdown 笔记** | 每个任务独立的 Markdown 文件，支持丰富格式 |
| **上下文系统** | 自动应用当前上下文到查询和新任务 |
| **单文件分发** | 单二进制文件，易于安装和分发 |
| **URL 提取** | `open` 命令自动提取并打开任务中的 URL |

**与 Capture 的关联**
- 直接借鉴其 "Git + Markdown" 的存储方案
- 参考其上下文系统设计，实现工作场景自动捕获
- 学习其笔记功能，设计任务的详细描述格式

---

### 1.3 todo.txt

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/todotxt/todo.txt-cli |
| **开发语言** | Shell/Bash |
| **Stars** | 3.5k+ |
| **许可证** | GPL-3.0 |

**项目简介**
todo.txt 是一个极简主义的任务管理方案，核心是一个纯文本文件格式规范。它不依赖特定软件，任何文本编辑器都可以管理任务。

**核心功能**
- 简单的文本格式规范：`优先级 任务内容 @上下文 +项目 due:日期`
- 纯文本存储，完全透明和可移植
- 丰富的社区工具和插件生态
- 跨平台支持（任何支持文本编辑的平台）

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **格式规范** | 简单、可读的文本格式，便于解析和版本控制 |
| **极简哲学** | 核心功能最小化，通过工具生态扩展 |
| **跨平台** | 不绑定特定平台或工具 |
| **社区生态** | 标准化的格式促进了丰富的第三方工具 |

**与 Capture 的关联**
- 参考其格式设计，定义 Capture 的任务文件格式
- 学习其极简哲学，保持核心功能的简洁性
- 借鉴其生态思路，设计可扩展的插件机制

---

### 1.4 TaskLite

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/ad-si/TaskLite |
| **开发语言** | Haskell |
| **Stars** | 800+ |
| **许可证** | MIT |

**项目简介**
TaskLite 是一个基于 YAML 文件的任务管理器，专为与 Git 同步而设计。它强调数据的可移植性和与版本控制的集成。

**核心功能**
- YAML 文件格式，人类可读且易于解析
- Git 原生集成，数据同步通过 Git 完成
- SQLite 后端支持（可选）
- 任务依赖和阻塞关系
- 丰富的元数据支持（创建时间、关闭时间、UUID 等）

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **YAML 格式** | 比 JSON 更易读，比纯文本更结构化 |
| **UUID 标识** | 使用 UUID 作为任务唯一标识，避免冲突 |
| **Git 集成** | 数据存储天然适合 Git 版本控制 |
| **元数据丰富** | 记录完整的时间线和状态变更 |

**与 Capture 的关联**
- 参考其 YAML + Git 的方案设计存储层
- 学习其 UUID 设计，实现全局唯一任务 ID
- 借鉴其元数据设计，记录任务的完整生命周期

---

### 1.5 综合对比

| 项目 | 存储格式 | 同步方式 | 学习曲线 | 扩展性 |
|------|----------|----------|----------|--------|
| Taskwarrior | 纯文本 | Taskserver | 中等 | 高 (Hooks) |
| dstask | YAML + Git | Git | 低 | 中 |
| todo.txt | 纯文本 | 任意文件同步 | 极低 | 高 (生态) |
| TaskLite | YAML + Git | Git | 中等 | 中 |

---

## 2. CLI/TUI 界面模块

### 2.1 Textual (Python TUI 框架)

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/Textualize/textual |
| **开发语言** | Python |
| **Stars** | 25k+ |
| **许可证** | MIT |

**项目简介**
Textual 是一个用于 Python 的现代 TUI（终端用户界面）框架，由 Rich 库的开发者创建。它使用声明式 UI 和 CSS 样式，让开发者能够创建复杂的交互式终端应用。

**核心功能**
- 声明式 UI 构建（类似 React）
- CSS 样式的界面布局
- 丰富的内置组件（DataTable、Tree、Input、Button 等）
- 响应式设计，自动适应终端大小
- 支持鼠标交互
- 内置动画和过渡效果

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **组件化设计** | 声明式组件构建复杂的 TUI 界面 |
| **CSS 样式** | 使用 CSS 进行界面布局和样式定义 |
| **事件系统** | 完善的消息和事件传递机制 |
| **响应式布局** | 自动适应不同终端尺寸 |

**与 Capture 的关联**
- 作为 Capture TUI 看板的首选技术框架
- 参考其组件设计，构建看板的列、卡片等组件
- 学习其事件系统，实现拖拽等交互

---

### 2.2 kanban-tui (Zaloog)

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/Zaloog/kanban-tui |
| **开发语言** | Python |
| **Stars** | 500+ |
| **许可证** | MIT |

**项目简介**
kanban-tui 是一个基于 Textual 框架的看板任务管理器，支持多种后端（SQLite、Jira、Claude）。它是目前功能最完整的 Python TUI 看板应用之一。

**核心功能**
- 多后端架构（SQLite、Jira、Claude 任务文件）
- 任务依赖管理
- 多看板支持
- 可自定义列
- 任务依赖阻止移动
- CLI 接口支持（适合 Agent 调用）
- Web 模式（通过浏览器访问）

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **多后端设计** | 抽象存储层，支持不同数据源 |
| **任务依赖** | 依赖关系阻止任务状态变更 |
| **CLI + TUI 双模式** | 既支持交互式 TUI，也支持命令式 CLI |
| **Web 模式** | Textual 应用通过浏览器访问 |

**与 Capture 的关联**
- 直接参考其看板布局和交互设计
- 学习其多后端架构，设计 Capture 的存储抽象
- 借鉴其 CLI + TUI 双模式设计

---

### 2.3 taskwarrior-tui

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/kdheepak/taskwarrior-tui |
| **开发语言** | Rust |
| **Stars** | 1k+ |
| **许可证** | MIT |

**项目简介**
taskwarrior-tui 是一个 Rust 编写的 Taskwarrior 终端界面。它提供了一个快速的、交互式的界面来管理 Taskwarrior 任务。

**核心功能**
- 与 Taskwarrior 完全集成
- Vim 风格的键绑定
- 快速过滤和搜索
- 任务报告和统计
- 批处理操作

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **Vim 键绑定** | hjkl 导航，适合键盘重度用户 |
| **快速过滤** | 实时过滤任务列表 |
| **批处理** | 多选任务进行批量操作 |

**与 Capture 的关联**
- 参考其键盘导航设计，提供 Vim 风格的快捷键
- 学习其过滤实现，设计快速搜索功能

---

### 2.4 kanbanban (Rust)

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/luizvbo/kanbanban |
| **开发语言** | Rust |
| **Stars** | 300+ |
| **许可证** | MIT |

**项目简介**
kanbanban 是一个 Rust 编写的高性能终端看板应用，采用模态驱动的界面设计，支持 Markdown 和外部编辑器。

**核心功能**
- 模态驱动界面（类似 Vim）
- Markdown 任务描述渲染
- 外部编辑器集成（Vim、VS Code 等）
- YAML 文件存储
- 全局标签注册表

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **模态界面** | Normal/Insert 模式分离，防止误操作 |
| **外部编辑器** | 支持在 Vim/VS Code 中编辑任务详情 |
| **Markdown 渲染** | 在终端中渲染 Markdown 格式 |

**与 Capture 的关联**
- 参考其模态设计，实现防误触的编辑模式
- 学习其外部编辑器集成，支持用户偏好的编辑器

---

### 2.5 综合对比

| 项目 | 语言 | 框架 | 主要特点 |
|------|------|------|----------|
| Textual | Python | 自研 | 现代声明式 TUI 框架 |
| kanban-tui | Python | Textual | 功能完整的看板，多后端 |
| taskwarrior-tui | Rust | Rust TUI | 高性能，Vim 风格 |
| kanbanban | Rust | Ratatui | 模态界面，Markdown |

---

## 3. Web/Desktop UI 看板模块

### 3.1 Wekan

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/wekan/wekan |
| **开发语言** | JavaScript (Meteor) |
| **Stars** | 14k+ |
| **许可证** | MIT |

**项目简介**
Wekan 是一个开源的、自托管的 Kanban 工具，是 Trello 的主要开源替代品。它基于 Meteor 全栈框架构建，提供实时协作功能和丰富的项目管理特性。

**核心功能**
- 实时协作看板，支持多用户同时编辑
- 拖拽式任务卡片管理
- 标签、成员分配、截止日期
- 任务清单（Checklists）和附件
- 过滤和搜索功能
- 多语言支持（50+ 语言）
- REST API 和 Webhook
- 多种认证方式（LDAP、OAuth、SAML）

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **看板交互** | 流畅的拖拽体验和实时同步 |
| **权限模型** | 看板级、列表级、卡片级的权限控制 |
| **通知系统** | 邮件通知、Web 推送、桌面通知 |
| **导入导出** | 支持 JSON/CSV 导入导出 |
| **插件架构** | 基于 Meteor 的插件扩展机制 |

**与 Capture 的关联**
- 参考其看板 UI 布局和交互设计
- 学习其实时同步机制，为未来 Capture 的 Web 版提供参考
- 借鉴其权限模型，设计多用户协作功能

---

### 3.2 Kanboard

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/kanboard/kanboard |
| **开发语言** | PHP |
| **Stars** | 8.5k+ |
| **许可证** | MIT |

**项目简介**
Kanboard 是一个极简主义的 Kanban 工具，专注于简单性和易用性。它采用 Kanban 方法论，限制进行中的工作量（WIP），帮助用户专注于当前任务。

**核心功能**
- 简洁的看板界面，专注核心功能
- WIP 限制（Work In Progress）
- 自动动作（Automated Actions）- 类似 IFTTT 的规则引擎
- 任务泳道（Swimlanes）
- 时间跟踪和燃尽图
- 插件扩展系统（100+ 插件）
- 支持 SQLite、MySQL、PostgreSQL
- Docker 一键部署

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **极简设计** | 去除冗余功能，专注于看板核心体验 |
| **自动化规则** | 基于事件的自动化工作流 |
| **WIP 限制** | 防止同时处理过多任务，提升专注力 |
| **插件生态** | 完善的插件开发文档和生态 |
| **数据库抽象** | 支持多种数据库后端 |

**与 Capture 的关联**
- 学习其极简设计理念，保持 Capture 的核心功能简洁
- 参考其自动化规则设计，实现任务状态自动流转
- 借鉴其插件架构，设计 Capture 的扩展机制

---

### 3.3 Planka

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/plankanban/planka |
| **开发语言** | React/Node.js |
| **Stars** | 4.5k+ |
| **许可证** | AGPL-3.0 |

**项目简介**
Planka 是一个开源的、基于 React 和 Redux 的 Kanban 工具，界面现代美观，功能对标 Trello。它提供实时更新和丰富的项目管理功能。

**核心功能**
- React + Redux 现代前端架构
- 实时协作和即时更新
- 项目、看板、列表、卡片四级结构
- 标签、截止日期、成员分配
- 卡片模板功能
- 项目背景自定义
- OIDC/SAML 单点登录
- Docker 支持

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **前端架构** | React + Redux 的状态管理 |
| **实时同步** | WebSocket 实现实时协作 |
| **卡片模板** | 可复用的任务模板 |
| **项目管理** | 项目和看板的分层组织 |
| **美观 UI** | 现代化的界面设计语言 |

**与 Capture 的关联**
- 参考其 React 组件设计，为未来 Web 版提供架构参考
- 学习其 Redux 状态管理模式
- 借鉴其卡片模板功能，实现 Capture 的任务模板

---

### 3.4 Kanri

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/trobonox/kanri |
| **开发语言** | TypeScript (Tauri + Nuxt.js) |
| **Stars** | 800+ |
| **许可证** | Apache-2.0 |

**项目简介**
Kanri 是一个跨平台的桌面 Kanban 应用，使用 Tauri 和 Nuxt.js 构建。它注重隐私（本地优先）、用户体验和离线工作能力。

**核心功能**
- Tauri 构建的轻量级桌面应用
- 离线优先，数据本地存储
- 多看板管理
- 标签和颜色编码
- 主题和自定义背景
- 数据导出（JSON）
- 键盘快捷键支持
- 自动保存

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **Tauri 架构** | Rust + Web 前端的轻量级桌面应用方案 |
| **离线优先** | 本地数据存储，无需网络 |
| **跨平台** | 一套代码支持 Windows/macOS/Linux |
| **性能** | 比 Electron 更小的包体积和内存占用 |
| **数据可移植** | JSON 格式的数据导出导入 |

**与 Capture 的关联**
- 为未来 Capture Desktop GUI 提供技术选型参考（Tauri vs Electron）
- 学习其离线优先的数据管理策略
- 借鉴其跨平台桌面应用开发模式

---

### 3.5 Focalboard

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/mattermost/focalboard |
| **开发语言** | Go + TypeScript (React) |
| **Stars** | 20k+ |
| **许可证** | Apache-2.0 |

**项目简介**
Focalboard 是 Mattermost 开发的开源项目管理工具，提供自托管的看板、表格和日历视图。它既是独立的桌面应用，也可以作为 Mattermost 的插件。

**核心功能**
- 多种视图：看板、表格、日历、画廊
- 多平台：Web、桌面（Windows/macOS/Linux）、移动
- 与 Mattermost 集成
- 模板库
- 导入 Trello、Asana、Notion 数据
- 自托管或个人本地使用
- 实时协作

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **多视图** | 同一数据的不同展示方式 |
| **多平台** | Web、桌面、移动端统一体验 |
| **导入导出** | 支持多种第三方工具数据导入 |
| **模板系统** | 预设模板和自定义模板 |
| **集成能力** | 与团队协作工具深度集成 |

**与 Capture 的关联**
- 参考其多视图设计，为 Capture 规划看板/表格/日历视图
- 学习其多平台架构，规划 Capture 的 Web/Desktop 版本
- 借鉴其模板系统，设计 Capture 的任务模板

---

### 3.6 综合对比

| 项目 | 类型 | 技术栈 | 主要特点 |
|------|------|--------|----------|
| Wekan | Web | Meteor | 功能完整，实时协作 |
| Kanboard | Web | PHP | 极简主义，WIP 限制 |
| Planka | Web | React/Node.js | 现代 UI，实时更新 |
| Kanri | Desktop | Tauri/Vue | 离线优先，轻量级 |
| Focalboard | Web/Desktop | Go/React | 多视图，多平台 |


---

---

## 4. Vibe Coding 看板模块（AI Agent 专用）

> Vibe Coding 是指使用 AI Agent 进行编程的新范式，这类看板专门为管理和编排 AI 编程 Agent 设计。

### 4.1 Vibe Kanban (BloopAI)

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/BloopAI/vibe-kanban |
| **开发语言** | Rust + TypeScript |
| **Stars** | 22k+ |
| **许可证** | Apache-2.0 |

**项目简介**
Vibe Kanban 是专为 AI 编程 Agent 设计的编排平台，由 BloopAI 开发。它解决了开发者在使用 Claude Code、Gemini CLI 等工具时的 "doomscrolling gap" 问题——等待 Agent 完成时的空闲时间。

**核心功能**
- **并行执行**: 同时运行多个 AI Agent，提高效率
- **Git Worktree 隔离**: 每个 Agent 在独立的 git worktree 中工作，避免冲突
- **看板管理**: 任务状态流转（To Do → In Progress → Review → Done）
- **实时流式输出**: WebSocket 实时显示 Agent 的执行日志和思考过程
- **多 Agent 支持**: Claude Code、Codex、Gemini CLI、OpenCode、Cursor 等 10+ 种 Agent
- **代码审查**: 内置 diff 查看和行内评论
- **一键合并**: 自动 rebase 并合并到 main 分支
- **开发服务器集成**: 一键启动 dev server 预览结果

**架构亮点**
- Rust 后端（高性能，内存安全）
- Git Worktree 实现零开销隔离（无需 Docker）
- SQLite 本地优先存储
- MCP（Model Context Protocol）双向支持

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **并行编排** | 多 Agent 同时工作的调度机制 |
| **隔离方案** | Git Worktree 替代 Docker，轻量且高效 |
| **实时反馈** | WebSocket 流式传输 Agent 输出 |
| **成本追踪** | Token 使用量和成本实时监控 |
| **MCP 集成** | 标准化的 Agent 工具调用协议 |

**与 Capture 的关联**
- Vibe Kanban 是 Capture 的**直接竞品和参考标杆**
- 参考其 Git Worktree 隔离方案设计 Capture 的执行环境
- 学习其并行 Agent 编排机制
- 借鉴其成本追踪和预算控制功能

---

### 4.2 KaibanJS

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/kaiban-ai/kaibanjs |
| **开发语言** | JavaScript |
| **Stars** | 5k+ |
| **许可证** | MIT |

**项目简介**
KaibanJS 是一个 JavaScript 原生框架，用于构建和管理多 Agent 系统。它采用看板方法论，让开发者像管理团队一样管理 AI Agent。

**核心功能**
- **看板式 Agent 管理**: 类似 Trello 的界面管理 Agent 任务
- **角色定义**: 为每个 Agent 定义角色（开发者、产品经理、QA 等）
- **任务传递**: 任务结果自动传递给下游任务
- **工具集成**: 支持 LangChain 兼容的工具
- **状态管理**: Redux 风格的状态管理
- **记忆管理**: 团队级别的记忆和上下文控制
- **可观测性**: 详细的日志和成本追踪

**架构特点**
```javascript
// 定义 Agent
const developer = new Agent({
  name: 'Dave',
  role: 'Developer',
  goal: '编写和审查代码',
  background: '精通 JavaScript 和 React',
});

// 创建任务
const task = new Task({
  description: '开发登录功能',
  agent: developer,
});

// 组建团队
const team = new Team({
  name: '开发团队',
  agents: [developer],
  tasks: [task],
});
```

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **角色抽象** | Agent 角色和能力的声明式定义 |
| **任务流** | 任务结果的自动传递机制 |
| **状态管理** | Redux 风格的统一状态管理 |
| **框架设计** | 类似 React 的声明式 API 设计 |
| **成本控制** | Token 使用量的细粒度追踪 |

**与 Capture 的关联**
- 参考其角色定义设计 Capture 的 Agent 配置
- 学习其任务依赖和结果传递机制
- 借鉴其 JavaScript 原生 API 设计

---

### 4.3 Agent Deck

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/asheshgoplani/agent-deck |
| **开发语言** | Go |
| **Stars** | 800+ |
| **许可证** | MIT |

**项目简介**
Agent Deck 是一个终端会话管理器，用于管理多个 AI 编程 Agent（Claude、Gemini、OpenCode、Codex 等）。它是 "AI Agent 的任务控制中心"。

**核心功能**
- **TUI 界面**: 终端用户界面管理所有 Agent 会话
- **多工具支持**: Claude Code、Gemini CLI、OpenCode、Codex、Cursor 等
- **会话分叉**: 从现有会话创建分支（类似 git fork）
- **MCP 管理**: 附加/分离 MCP 服务器到会话
- **技能管理**: 管理 Claude Code 的技能（Skills）
- **成本追踪**: 实时 Token 使用量和成本统计
- **Web 仪表板**: 浏览器查看成本图表和统计

**使用场景**
```bash
# 启动 TUI
agent-deck

# 添加项目
agent-deck add . -c claude

# 分叉会话
agent-deck session fork my-project

# 附加 MCP
agent-deck mcp attach my-project exa

# 启动 Web 界面
agent-deck web
```

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **TUI 设计** | 终端界面管理多个 Agent 会话 |
| **会话管理** | 会话的创建、分叉、恢复机制 |
| **成本监控** | 实时成本追踪和预算限制 |
| **MCP 集成** | 动态附加/分离 MCP 服务器 |
| **Web+TUI 双模式** | 同时提供终端和 Web 界面 |

**与 Capture 的关联**
- 参考其 TUI 设计，增强 Capture 的终端体验
- 学习其成本追踪实现
- 借鉴其 MCP 动态管理功能

---

### 4.4 Claw-Kanban

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/GreenSheep01201/Claw-Kanban |
| **开发语言** | TypeScript/Node.js |
| **Stars** | 300+ |
| **许可证** | MIT |

**项目简介**
Claw-Kanban 是一个 AI Agent 编排看板，支持将任务路由到 Claude Code、Codex CLI 和 Gemini CLI，具有基于角色的自动分配和实时监控功能。

**核心功能**
- **6 列看板**: Backlog → Todo → In Progress → Review → Done → Failed
- **角色映射**: 根据任务类型自动分配给特定 Agent
- **实时终端查看器**: 实时查看 Agent 的工作过程
- **AGENTS.md 集成**: AI Agent 可以读写看板任务
- **Telegram 集成**: 通过 Telegram 消息创建和跟踪任务
- **自动检测**: 自动检测已安装的 AI CLI 工具

**工作流程**
1. 创建任务卡片并分配 Agent
2. Agent 在独立目录中执行任务
3. 实时查看终端输出
4. 审查代码差异
5. 一键合并或反馈修改

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **角色路由** | 基于角色的任务自动分配 |
| **实时查看** | 终端输出的实时流式显示 |
| **AGENTS.md** | Agent 与看板的协议化交互 |
| **即时通讯集成** | Telegram Bot 集成 |
| **失败处理** | Failed 列和重试机制 |

**与 Capture 的关联**
- 参考其 AGENTS.md 协议设计 Capture 的 Agent 交互接口
- 学习其角色路由机制
- 借鉴其失败处理和重试逻辑

---

### 4.5 4ga Boards

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/RARgames/4gaBoards |
| **开发语言** | JavaScript (Node.js/React) |
| **Stars** | 1.5k+ |
| **许可证** | MIT |

**项目简介**
4ga Boards 是一个实时看板管理工具，具有优雅的深色模式、可折叠的待办事项列表和多任务处理能力。即将支持 GitHub 双向同步。

**核心功能**
- **实时更新**: 无需刷新页面的实时协作
- **深色模式**: 优雅的深色主题
- **高级 Markdown 编辑器**: 支持富文本编辑
- **多级层级**: 项目 → 看板 → 列表 → 卡片 → 任务
- **多任务处理**: 同时编辑/审查卡片和筛选/重新排列看板
- **可折叠列表和侧边栏**: 节省屏幕空间
- **快捷键**: 强大的键盘快捷键支持
- **GitHub 双向同步**: 即将推出

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **实时协作** | WebSocket 实现的实时同步 |
| **多级层级** | 项目-看板-列表-卡片-任务五级结构 |
| **多任务 UI** | 同时处理多个任务不丢失本地更改 |
| **深色模式** | 优雅的深色主题设计 |
| **快捷键** | 完整的键盘操作支持 |

**与 Capture 的关联**
- 参考其多级层级结构设计 Capture 的任务组织
- 学习其实时协作实现
- 借鉴其深色模式 UI 设计

---

### 4.6 综合对比

| 项目 | 类型 | 技术栈 | 主要特点 | 适用场景 |
|------|------|--------|----------|----------|
| Vibe Kanban | Web | Rust+TS | 并行编排、Git Worktree 隔离 | 多 Agent 并行开发 |
| KaibanJS | 框架 | JavaScript | 角色定义、任务流 | 构建多 Agent 应用 |
| Agent Deck | TUI/Web | Go | 会话管理、成本追踪 | 终端用户管理 Agent |
| Claw-Kanban | Web | TypeScript | 角色路由、AGENTS.md | Agent 任务编排 |
| 4ga Boards | Web | Node+React | 实时协作、多级层级 | 通用项目管理 |

**选型建议**：
- **多 Agent 并行开发**: Vibe Kanban（功能最完整）
- **构建自己的 Agent 应用**: KaibanJS（框架级）
- **终端重度用户**: Agent Deck
- **快速原型**: Claw-Kanban
- **通用项目**: 4ga Boards


## 5. AI Agent 执行模块

### 5.1 Claude Code

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/anthropics/claude-code |
| **开发语言** | TypeScript |
| **Stars** | 5k+ |
| **许可证** | 专有软件 |

**项目简介**
Claude Code 是 Anthropic 官方推出的智能编程助手，可以直接在终端中与 Claude 对话并执行代码操作。它是 Capture 计划集成的核心 Agent 之一。

**核心功能**
- 自然语言代码编辑
- 代码库理解和导航
- 终端命令执行
- 文件系统操作
- 多步骤任务规划
- 代码审查和建议

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **上下文管理** | 维护会话上下文，支持多轮对话 |
| **工具调用** | 将 AI 能力封装为可调用的工具 |
| **安全确认** | 执行敏感操作前要求用户确认 |
| **进度反馈** | 长时间任务提供进度更新 |

**与 Capture 的关联**
- 作为 Capture 的首选执行 Agent
- 参考其工具调用设计，抽象 Agent 接口
- 学习其安全模型，设计执行权限控制

---

### 5.2 LangGraph

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/langchain-ai/langgraph |
| **开发语言** | Python |
| **Stars** | 24k+ |
| **许可证** | MIT |

**项目简介**
LangGraph 是 LangChain 生态系统中的 Agent 编排框架，专注于构建可控、有状态的 Agent 工作流。它被 Cisco、Uber、LinkedIn 等公司用于生产环境。

**核心功能**
- 有状态 Agent 编排
- 支持循环、条件分支的工作流
- 长期记忆管理
- 人机协作（Human-in-the-loop）
- 多 Agent 协作
- 与 LangSmith 集成监控

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **状态机设计** | 使用状态机管理 Agent 执行流程 |
| **记忆管理** | 长期和短期记忆的分离和持久化 |
| **人机协作** | 在关键节点暂停等待人类输入 |
| **可观测性** | 完整的执行追踪和监控 |

**与 Capture 的关联**
- 参考其状态机设计，实现任务执行的状态管理
- 学习其记忆系统，设计上下文传递机制
- 借鉴其可观测性设计，实现执行日志记录

---

### 5.3 OpenAI Agents SDK

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/openai/openai-agents-python |
| **开发语言** | Python |
| **Stars** | 19k+ |
| **许可证** | MIT |

**项目简介**
OpenAI Agents SDK 是 OpenAI 官方发布的轻量级 Agent 框架，专注于多 Agent 工作流、追踪和安全护栏。

**核心功能**
- 轻量级多 Agent 工作流
- 内置追踪（Tracing）
- 安全护栏（Guardrails）
- Agent 交接（Handoffs）
- 支持 100+ LLM 提供商

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **护栏设计** | 输入/输出验证，内容过滤 |
| **追踪系统** | 详细的执行追踪和调试信息 |
| **Agent 交接** | 多 Agent 之间的任务委托 |
| **轻量级** | 核心功能精简，易于理解和扩展 |

**与 Capture 的关联**
- 作为 Capture 的可选 Agent 之一（Codex）
- 参考其护栏设计，实现执行安全检查
- 学习其追踪系统，设计执行日志

---

### 5.4 Smolagents

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/huggingface/smolagents |
| **开发语言** | Python |
| **Stars** | 15k+ |
| **许可证** | Apache-2.0 |

**项目简介**
Smolagents 是 Hugging Face 开发的极简 Agent 框架，核心理念是"Agent 直接编写和执行 Python 代码"，无需将意图转换为 JSON。

**核心功能**
- CodeAgent：生成并执行 Python 代码
- ToolCallingAgent：支持 JSON 调用
- 沙箱代码执行
- 轻量级设计（核心代码 < 1000 行）
- 与 Transformers 生态集成

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **代码优先** | Agent 直接生成代码而非结构化数据 |
| **沙箱执行** | 隔离的执行环境确保安全 |
| **极简设计** | 核心代码精简，易于理解和定制 |
| **工具即函数** | Python 函数直接作为工具 |

**与 Capture 的关联**
- 参考其代码优先方法，设计 Agent 执行接口
- 学习其沙箱设计，实现安全的代码执行
- 借鉴其工具定义方式，简化 Agent 集成

---

### 5.5 OpenManus

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/mannaandpoem/OpenManus |
| **开发语言** | Python |
| **Stars** | 30k+ |
| **许可证** | MIT |

**项目简介**
OpenManus 是 MetaGPT 团队在 2025 年 3 月推出的开源 AI Agent 框架，作为 Manus AI 的开源替代品。它在 3 小时内完成开发，迅速获得社区关注，成为当时增长最快的 AI Agent 开源项目。

**核心功能**
- 基于 ReAct 框架的自主任务执行
- 浏览器自动化（Playwright）
- 多工具调用能力（文件操作、搜索、代码执行）
- 多 Agent 协作框架
- 本地 LLM 支持（Ollama）
- 实时任务规划和执行反馈

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **浏览器自动化** | Playwright 集成实现 Web 操作 |
| **多工具集成** | 统一的工具注册和调用机制 |
| **本地 LLM 支持** | 支持私有化部署的大模型 |
| **快速原型** | 简洁的代码结构，易于理解和扩展 |
| **社区驱动** | 开源社区快速迭代和改进 |

**与 Capture 的关联**
- 参考其工具调用设计，实现 Capture 的 Agent 工具集
- 学习其浏览器自动化，支持 Web 相关任务执行
- 借鉴其多 Agent 协作框架，设计复杂任务执行

---

### 5.6 Auto-GPT

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/Significant-Gravitas/AutoGPT |
| **开发语言** | Python |
| **Stars** | 170k+ |
| **许可证** | MIT |

**项目简介**
Auto-GPT 是 2023 年引发 AI Agent 热潮的开源项目，它展示了如何让 GPT 模型递归地执行任务，实现自主规划和执行。虽然经历了多次架构调整，但仍是 Agent 领域最具影响力的项目之一。

**核心功能**
- 递归任务执行
- 持久化记忆管理（向量数据库）
- Web 浏览和信息检索
- 文件系统操作
- 代码执行和调试
- Agent 工作流编排
- Forge 和 Benchmark 框架

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **自主执行** | 任务分解和自主执行的循环机制 |
| **记忆管理** | 长期记忆和上下文管理 |
| **递归模式** | 复杂任务的递归处理模式 |
| **工作流编排** | Agent 工作流的定义和执行 |
| **基准测试** | Agent 能力的评估框架 |

**与 Capture 的关联**
- 参考其任务递归执行模式，处理复杂多步骤任务
- 学习其记忆管理，设计 Capture 的上下文保持机制
- 借鉴其工作流编排，实现可配置的任务执行流程

---

### 5.7 MetaGPT

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/geekan/MetaGPT |
| **开发语言** | Python |
| **Stars** | 50k+ |
| **许可证** | MIT |

**项目简介**
MetaGPT 是一个多智能体协作框架，模拟软件公司的组织结构，通过不同角色的 Agent（CEO、产品经理、架构师、程序员等）协作完成复杂任务。它是 OpenManus 的基础框架。

**核心功能**
- 多角色 Agent 协作
- 标准化操作流程（SOP）
- 自然语言编程
- 代码自动生成和审查
- 需求分析和架构设计
- 可观测性和调试工具

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **角色分配** | 不同 Agent 担任不同角色 |
| **SOP 设计** | 标准化的操作流程 |
| **协作机制** | Agent 间的消息传递和协作 |
| **代码生成** | 从需求到代码的完整流程 |
| **可观测性** | 多 Agent 系统的调试和监控 |

**与 Capture 的关联**
- 参考其多角色设计，实现不同 Agent 执行不同类型任务
- 学习其 SOP 设计，定义任务执行的标准流程
- 借鉴其协作机制，实现多 Agent 协作完成复杂任务

---

### 5.8 CrewAI

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/crewAIInc/crewAI |
| **开发语言** | Python |
| **Stars** | 25k+ |
| **许可证** | MIT |

**项目简介**
CrewAI 是一个用于构建多 Agent 系统的框架，专注于角色扮演和任务委派。它让开发者能够创建扮演不同角色的 Agent 团队，通过协作完成复杂工作流。

**核心功能**
- 基于角色的 Agent 定义
- 任务委派和协作
- 工作流编排
- 工具集成
- 记忆和上下文管理
- 与 LangChain 生态集成

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **角色定义** | 声明式定义 Agent 角色和能力 |
| **任务委派** | 主 Agent 将任务委派给专业 Agent |
| **工作流** | 顺序、并行、条件工作流 |
| **工具绑定** | Agent 与工具的灵活绑定 |
| **企业级** | 生产环境的高可用设计 |

**与 Capture 的关联**
- 参考其角色定义，为 Capture Agent 设置不同执行角色
- 学习其任务委派机制，实现智能任务分配
- 借鉴其工作流设计，支持复杂任务的编排

---

### 5.9 AgentGPT

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/reworkd/AgentGPT |
| **开发语言** | TypeScript/Next.js |
| **Stars** | 30k+ |
| **许可证** | GPL-3.0 |

**项目简介**
AgentGPT 是一个在浏览器中运行的自主 AI Agent 平台。它提供了简洁的 Web UI，让非技术用户也能轻松创建和运行 AI Agent，无需编写代码。

**核心功能**
- 浏览器内运行，无需安装
- 可视化 Agent 配置
- 任务自动分解和执行
- 实时执行日志
- 支持多种 LLM 提供商
- 自托管支持

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **Web UI** | 直观的 Agent 配置和监控界面 |
| **零代码** | 非技术用户也能使用 |
| **实时反馈** | 任务执行过程的实时展示 |
| **可访问性** | 降低 Agent 技术的使用门槛 |
| **部署灵活** | 云服务和自托管双模式 |

**与 Capture 的关联**
- 参考其 Web UI 设计，为未来 Capture Web 版提供界面参考
- 学习其零代码配置方式，简化 Agent 配置
- 借鉴其实时反馈机制，设计任务执行状态展示

### 5.10 综合对比

| 项目 | 类型 | 主要特点 | 适用场景 |
|------|------|----------|----------|
| Claude Code | 闭源产品 | 编程专用，深度代码理解 | 代码生成、重构 |
| LangGraph | 编排框架 | 状态机，复杂工作流 | 多步骤任务编排 |
| OpenAI Agents SDK | Agent 框架 | 轻量，护栏，追踪 | 通用 Agent 应用 |
| Smolagents | Agent 框架 | 极简，代码优先 | 代码生成任务 |
| OpenManus | Agent 框架 | 浏览器自动化，多工具 | Web 操作、自动化 |
| Auto-GPT | 自主 Agent | 递归执行，记忆管理 | 自主任务执行 |
| MetaGPT | 多 Agent 框架 | 角色扮演，SOP | 复杂协作任务 |
| CrewAI | 多 Agent 框架 | 角色定义，任务委派 | 团队工作流 |
| AgentGPT | Web Agent | 零代码，浏览器运行 | 快速原型、演示 |

**选型建议**：
- **编程任务**：Claude Code（首选）、Smolagents
- **复杂工作流**：LangGraph、CrewAI
- **多 Agent 协作**：MetaGPT、CrewAI
- **Web 自动化**：OpenManus
- **快速原型**：AgentGPT、OpenManus

---

## 6. Bot 集成模块

### 6.1 nanobot

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/HKUDS/nanobot |
| **开发语言** | Python |
| **Stars** | 2k+ |
| **许可证** | MIT |

**项目简介**
nanobot 是一个多平台聊天机器人框架，支持 Telegram、Discord、WhatsApp、WeChat、飞书等多个平台。它是 OpenClaw 的轻量级开源替代品。

**核心功能**
- 多平台支持（10+ 聊天应用）
- WebSocket 长连接（飞书等）
- 流式消息响应
- 定时任务（Cron）
- 消息模板系统
- 多 LLM 提供商支持

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **多平台抽象** | 统一的消息接口适配不同平台 |
| **WebSocket 模式** | 飞书 Bot 使用 WebSocket 而非 Webhook |
| **配置管理** | 灵活的 JSON 配置系统 |
| **权限控制** | `allowFrom` 等访问控制机制 |

**与 Capture 的关联**
- 直接参考其飞书 Bot 实现
- 学习其多平台抽象层设计
- 借鉴其配置管理模式

---

### 6.2 Feishu-Webhook-Proxy

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/ConnectAI-E/Feishu-Webhook-Proxy |
| **开发语言** | Python/Go |
| **Stars** | 300+ |
| **许可证** | MIT |

**项目简介**
这是一个飞书 Webhook 代理服务，用于解决飞书 Bot 回调的内网穿透问题。它使用 WebSocket 长连接，无需公网 IP。

**核心功能**
- WebSocket 长连接接收飞书消息
- 支持多 Bot 管理
- 消息转发和响应
- Python SDK（ca-lark-websocket）

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **WebSocket 方案** | 解决 Webhook 需要公网 IP 的问题 |
| **多 Bot 管理** | 一个连接管理多个 Bot |
| **消息解析** | 飞书消息结构的解析和验证 |

**与 Capture 的关联**
- 参考其 WebSocket 方案设计飞书 Bot 连接
- 学习其消息解析和验证逻辑

---

### 6.3 NoneBot2

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/nonebot/nonebot2 |
| **开发语言** | Python |
| **Stars** | 6k+ |
| **许可证** | MIT |

**项目简介**
NoneBot2 是一个现代、跨平台的 Python 聊天机器人框架，主要面向 QQ（OneBot 协议），但也支持其他平台。

**核心功能**
- 插件系统
- 依赖注入
- 事件驱动架构
- 多平台适配器
- 命令解析器

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **插件系统** | 动态加载和卸载插件 |
| **依赖注入** | 使用依赖注入管理组件 |
| **命令解析** | 强大的命令解析和参数提取 |
| **适配器模式** | 统一接口适配不同聊天平台 |

**与 Capture 的关联**
- 参考其插件系统设计 Capture 的扩展机制
- 学习其命令解析器，实现 Bot 消息解析
- 借鉴其适配器模式，支持多 Bot 平台

---

### 6.4 综合对比

| 项目 | 支持平台 | 架构特点 | 主要优势 |
|------|----------|----------|----------|
| nanobot | 10+ | 多 Provider 支持 | 功能完整，配置灵活 |
| Feishu-Webhook-Proxy | 飞书 | WebSocket 代理 | 无需公网 IP |
| NoneBot2 |  primarily QQ | 插件+依赖注入 | 生态丰富，扩展性强 |

---

## 7. 通知系统模块

### 7.1 knockknock

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/huggingface/knockknock |
| **开发语言** | Python |
| **Stars** | 3.5k+ |
| **许可证** | MIT |

**项目简介**
knockknock 是 Hugging Face 开发的通知库，用于在函数执行完成或失败时发送通知。支持 12+ 种通知渠道。

**核心功能**
- 装饰器方式使用
- 支持 Email、Slack、Discord、Telegram 等
- 自动捕获函数执行结果和异常
- 支持异步函数

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **装饰器模式** | 非侵入式地添加通知功能 |
| **多平台支持** | 统一接口，多种通知渠道 |
| **异常捕获** | 自动捕获并通知异常信息 |

**与 Capture 的关联**
- 参考其装饰器设计，实现任务执行通知
- 学习其多平台抽象，支持多种通知渠道

---

### 7.2 apprise

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/caronc/apprise |
| **开发语言** | Python |
| **Stars** | 12k+ |
| **许可证** | MIT |

**项目简介**
Apprise 是一个强大的多平台通知库，支持 100+ 种通知服务。它提供了统一的接口来发送通知到各种平台。

**核心功能**
- 100+ 通知服务支持
- 统一的 URL 格式配置通知服务
- 支持附件发送
- 异步发送支持
- 通知队列和批量发送

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **URL 配置** | `service://token@channel` 格式的服务配置 |
| **批量发送** | 支持向多个渠道同时发送 |
| **异步支持** | 非阻塞的通知发送 |

**与 Capture 的关联**
- 参考其 URL 配置方式，简化通知服务配置
- 学习其批量发送机制，实现多渠道通知

---

### 7.3 综合对比

| 项目 | 支持渠道 | 使用方式 | 特点 |
|------|----------|----------|------|
| knockknock | 12+ | 装饰器 | 简单易用，适合函数通知 |
| apprise | 100+ | 函数调用 | 功能全面，配置灵活 |

---

## 8. 数据同步模块

### 8.1 syncall

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/bergercookie/syncall |
| **开发语言** | Python |
| **Stars** | 800+ |
| **许可证** | MIT |

**项目简介**
syncall 是一个双向同步工具，支持 Taskwarrior、Google Calendar、Notion、Asana 等服务之间的数据同步。

**核心功能**
- 双向同步多种服务
- 冲突检测和解决
- 基于时间戳的增量同步
- 可配置的字段映射
- 冲突时的手动选择

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **双向同步算法** | 检测变更并合并到两端 |
| **冲突解决** | 时间戳对比和手动选择 |
| **字段映射** | 不同服务间的字段映射配置 |
| **增量同步** | 只同步变更的数据 |

**与 Capture 的关联**
- 直接参考其双向同步算法实现飞书同步
- 学习其冲突解决策略
- 借鉴其字段映射配置

---

### 8.2 Syncthing

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/syncthing/syncthing |
| **开发语言** | Go |
| **Stars** | 65k+ |
| **许可证** | MPL-2.0 |

**项目简介**
Syncthing 是一个去中心化的文件同步工具，P2P 架构，无需中心服务器。

**核心功能**
- P2P 去中心化同步
- 块级去重
- 版本控制
- 冲突文件处理
- 端到端加密

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **块级同步** | 只传输文件的变更部分 |
| **冲突处理** | 自动重命名冲突文件 |
| **版本控制** | 保留文件历史版本 |

**与 Capture 的关联**
- 参考其冲突处理机制
- 学习其块级同步思想（应用到记录级别）

---

### 8.3 Electric SQL

| 属性 | 详情 |
|------|------|
| **项目地址** | https://github.com/electric-sql/electric |
| **开发语言** | TypeScript/Elixir |
| **Stars** | 8k+ |
| **许可证** | Apache-2.0 |

**项目简介**
Electric 是一个本地优先的同步引擎，将 Postgres 数据同步到本地应用。

**核心功能**
- Postgres 逻辑复制
- 部分数据同步（Shapes）
- 离线优先架构
- 实时订阅

**可借鉴之处**
| 方面 | 借鉴点 |
|------|--------|
| **Shapes** | 只同步需要的数据子集 |
| **实时订阅** | 变更实时推送到客户端 |
| **离线优先** | 本地数据优先，后台同步 |

**与 Capture 的关联**
- 参考其部分同步思想，实现任务的增量同步
- 学习其实时订阅机制

---

### 8.4 综合对比

| 项目 | 同步模式 | 主要特点 | 适用场景 |
|------|----------|----------|----------|
| syncall | 双向同步 | 多服务集成，冲突解决 | 任务管理同步 |
| Syncthing | P2P 文件同步 | 去中心化，块级去重 | 文件同步 |
| Electric | 数据库同步 | 实时订阅，离线优先 | 应用数据同步 |

---

## 9. 架构设计参考

### 7.1 分层架构参考

```
┌─────────────────────────────────────────────────────────────┐
│                     交互层 (Interface)                       │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐     │
│  │   CLI    │  │   TUI    │  │   Bot    │  │  Web API │     │
│  │  Typer   │  │ Textual  │  │ FastAPI  │  │ FastAPI  │     │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘     │
├───────┴─────────────┴─────────────┴─────────────┴───────────┤
│                     服务层 (Service)                         │
│  ┌──────────────┐ ┌──────────────┐ ┌──────────────┐         │
│  │ TaskService  │ │ExecutionSvc  │ │  SyncService │         │
│  └──────────────┘ └──────────────┘ └──────────────┘         │
├─────────────────────────────────────────────────────────────┤
│                     核心层 (Core)                            │
│  ┌──────────────┐ ┌──────────────┐ ┌──────────────┐         │
│  │    Models    │ │   Storage    │ │    Events    │         │
│  │   Pydantic   │ │SQLite+Files  │ │   Event Bus  │         │
│  └──────────────┘ └──────────────┘ └──────────────┘         │
├─────────────────────────────────────────────────────────────┤
│                     适配层 (Adapter)                         │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐        │
│  │  Agents  │ │   Bots   │ │  Notify  │ │  Feishu  │        │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘        │
└─────────────────────────────────────────────────────────────┘
```

### 7.2 技术栈选型建议

| 模块 | 推荐方案 | 备选方案 | 理由 |
|------|----------|----------|------|
| CLI | Typer | Click, argparse | 类型提示，自动生成文档 |
| TUI | Textual | Blessed, npyscreen | 现代组件化，文档完善 |
| Web | FastAPI | Flask, Django | 异步支持，自动生成 API 文档 |
| ORM | SQLAlchemy 2.0 | Peewee, Tortoise | 类型安全，异步支持 |
| 数据验证 | Pydantic | Marshmallow, Cerberus | 性能优秀，生态丰富 |
| Agent 调用 | 子进程 + API | 直接库调用 | 隔离性，支持多种 Agent |

---

## 10. 总结与建议

### 9.1 核心参考项目

| Capture 模块 | 首推参考项目 | 次要参考项目 |
|--------------|--------------|--------------|
| 任务管理 | dstask | Taskwarrior, todo.txt |
| TUI 看板 | kanban-tui (Zaloog) | kanbanban, taskwarrior-tui |
| Web/Desktop 看板 | Kanri (Tauri) | Planka, Focalboard |
| **Vibe Coding 看板** | **Vibe Kanban** | **KaibanJS, Agent Deck** |
| Agent 集成 | LangGraph | OpenAI Agents SDK, OpenManus |
| Bot 集成 | nanobot | NoneBot2 |
| 通知系统 | apprise | knockknock |
| 数据同步 | syncall | Electric SQL |

### 9.2 关键设计决策建议

1. **存储方案**: 采用 dstask 的 "Git + Markdown/YAML" 方案，兼顾可读性和版本控制
2. **TUI 框架**: 选择 Textual，生态丰富，开发效率高
3. **Vibe Coding 看板**: **Vibe Kanban 是 Capture 的直接参考标杆**，其 Git Worktree 隔离和并行编排机制极具借鉴价值
4. **Web/Desktop 看板**: 如需桌面 GUI，参考 Kanri (Tauri)；如需 Web，参考 Planka (React)
5. **Agent 抽象**: 参考 LangGraph 的状态机设计，支持复杂执行流程
6. **多 Agent 协作**: 参考 MetaGPT/CrewAI 的角色分配和任务委派机制
7. **Bot 连接**: 参考 nanobot 的 WebSocket 方案，避免公网 IP 依赖
8. **通知系统**: 参考 apprise 的 URL 配置方式，简化多渠道配置
9. **同步机制**: 参考 syncall 的双向同步算法和冲突解决策略
10. **执行隔离**: 参考 Vibe Kanban 的 Git Worktree 方案，比 Docker 更轻量
11. **成本追踪**: 参考 Vibe Kanban 和 Agent Deck 的 Token 使用量监控
12. **MCP 集成**: 参考 Vibe Kanban 的 Model Context Protocol 实现

---

**文档版本**: 1.1  
**最后更新**: 2025-04-02
