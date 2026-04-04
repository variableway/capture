# Desktop UI Spec

## 定位

Desktop 是 Local Task Center 的本地增强壳层，目标是提升个人使用时的便捷性，而不是替代核心系统。

## 核心模块

### 1. Quick Capture

- 全局快捷键唤起
- 快速写入 idea
- 支持最小输入创建任务

### 2. Notification Center

- 本地通知
- 任务完成提醒
- review 待确认提醒

### 3. Local Dashboard Shell

- 承载 Web dashboard
- 作为本地任务中心入口

### 4. Embedded Terminal Entry

- 快速打开关联 repo
- 快速进入目标 terminal 或工作目录

## MVP 条件

- 只有在 Web 与 Local API 稳定后才有意义
- 第一版不应自建太多独占业务逻辑

## 风险

- 平台兼容性
- 打包与更新复杂度
- 与 Web 重复建设

## 结论

- Desktop 适合做“增强层”
- Desktop 不适合做“起步层”
