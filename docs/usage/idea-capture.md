# Capture 使用指南

## 概述

Capture 是一个命令行工具，帮助你快速捕捉工作中的想法和灵感，管理任务生命周期。

## 安装

```bash
go build -o capture .
# 或
go install .
```

## 初始化

首次使用需要初始化数据目录：

```bash
capture init
```

这会在 `~/.capture` 创建以下结构：
- `tasks/` — 任务 Markdown 文件
- `config.yaml` — 配置文件
- `capture.db` — SQLite 索引数据库
- `logs/` — 日志目录

## 基本使用

### 创建任务

```bash
# 快速创建
capture add "优化项目构建脚本"

# 带描述和标签
capture add "学习 Go 语言" -d "完成官方教程和练习" -t "学习,Go" -p high

# 参数说明
#   -d, --description   任务描述
#   -t, --tags          标签（逗号分隔）
#   -p, --priority      优先级: high, medium, low
```

### 查看任务列表

```bash
# 查看所有任务
capture list

# 按状态筛选
capture list -s todo
capture list -s in_progress
capture list -s done
```

### 查看任务详情

```bash
capture show TASK-00001
```

### 编辑任务

```bash
# 修改标题
capture edit TASK-00001 --title "新的标题"

# 修改优先级
capture edit TASK-00001 -p high

# 修改描述和标签
capture edit TASK-00001 -d "新的描述" -t "新标签1,新标签2"
```

### 修改任务状态

```bash
# 有效状态: todo, in_progress, done, cancelled, archived
capture status TASK-00001 in_progress
capture status TASK-00001 done
```

状态流转规则：
- `todo` → `in_progress`, `done`, `cancelled`
- `in_progress` → `done`, `cancelled`, `todo`
- `done` → `archived`
- `cancelled` → `todo`, `archived`

### 删除任务

```bash
# 需要确认
capture delete TASK-00001

# 强制删除（不需要确认）
capture delete TASK-00001 -f
```

## TUI 看板

```bash
capture kanban
```

启动交互式终端看板界面：

- 三个列：TODO | IN PROGRESS | DONE
- 键盘操作：
  - `↑/k` `↓/j` — 上下移动
  - `←/h` `→/l` — 左右切换列
  - `Enter` — 查看详情
  - `a` — 新建任务
  - `d` — 删除任务
  - `?` — 显示帮助
  - `q/Esc` — 退出

## 飞书 Bot

参见 [飞书配置指南](./feishu-setup.md)

## 配置管理

```bash
# 查看配置
capture config get app.data_dir

# 修改配置
capture config set defaults.editor vim
capture config set defaults.priority high
```

## 数据同步

```bash
# 推送到飞书多维表格
capture sync -d push

# 从飞书多维表格拉取
capture sync -d pull

# 双向同步
capture sync -d bidirectional
```

## 任务文件格式

任务以 Markdown + YAML frontmatter 存储在 `~/.capture/tasks/YYYY/MM/TASK-NNNNN.md`。
你可以直接用编辑器打开和编辑这些文件，Capture 会自动识别修改。
