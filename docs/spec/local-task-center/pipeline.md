# Pipeline Module Spec

## 目标

表达任务从想法到执行完成的流程阶段。

## 阶段定义

- inbox：原始输入
- mindstorm：问题发散
- analysis：可行性分析
- planning：方案规划
- prd：需求整理
- tasks：任务拆解
- dispatch：等待或已经下发
- execution：执行中
- review：结果确认

## 状态定义

- todo
- in_progress
- done
- cancelled
- archived

## 核心原则

1. stage 与 status 独立
2. stage 描述流程位置
3. status 描述执行状态

## 接口

- `capture stage <id> <stage>`
- `capture list --stage <stage>`
- `capture edit <id> --stage <stage>`

## MVP

- 阶段合法性校验
- CLI 可修改阶段
- TUI 展示阶段摘要
