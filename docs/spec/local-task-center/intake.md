# Intake Module Spec

## 目标

把来自不同入口的原始输入统一转换为 Task。

## 输入

- CLI 文本输入
- 飞书 Bot 消息
- 后续 Web 表单
- 后续 Desktop 快捷捕捉

## 输出

- 标准化 Task
- 默认 `status=todo`
- 默认 `stage=inbox`

## 核心字段

- title
- description
- source
- priority
- tags
- context

## 规则

1. 输入必须能落到统一 Task 模型
2. 原始输入尽量保留，不提前过度结构化
3. source 必须标记来源
4. 默认进入 inbox，等待后续分析

## MVP

- CLI `add`
- 飞书 Bot `记录`
