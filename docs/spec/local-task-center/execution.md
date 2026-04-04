# Execution Module Spec

## 目标

记录任务执行过程，而不是把执行细节散落在外部工具里。

## 当前范围

- execution.agent
- execution.model
- execution.exec_status
- execution.started_at
- execution.completed_at
- execution.result

## 下一步扩展

- execution.session_id
- execution.log_path
- execution.last_heartbeat_at
- execution.exit_code

## 阶段关系

- 任务开始实际运行时，stage 应进入 `execution`
- 执行结果等待人工确认时，stage 应进入 `review`

## 风险

- 不同 agent 的日志格式不同
- terminal 会话标识不统一
- 多仓库并发执行需要额外调度层

## MVP

- 本次不实现统一执行器
- 仅保留字段与阶段接口，为下一阶段实现留接口
