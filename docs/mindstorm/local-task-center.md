# Local Task Center

目前期望做的事情是：
1. 一个个人日常任务的收口中心
2. 任务的来源有哪些:
    1. 随机的idea输入，通过terminal，飞书机器人收集
    2. idea筛选之后编程想要实现的东西
    3. 想要实现的东西开始做分析，也就是：
      - mindstorm
      - analysis
      - planning
      - prd
      - tasks
    这些肯定也是计划过程。但是这些都是给AI Agent要做的事情的来源
3. 任务分配机制：
   1. 选择好任务之后分配飞对应的AI Agent执行
   2. AI Agent 再一个代码仓库中执行任务，任务通过收口中心下发
   3. AI Agent 最好是cli可以有一个可视化的terminal 看到
   4. AI Agent 完成任务之后反馈给任务中心
4. 从收集任务，下发任务，分析任务，任务确认这个是有一个大的看板也好，dashboard也好需要一个可视化页面
5. 任务执行阶段也需要这个可视化过程，那个地方在做那个任务，哪里可以可以看到，点击可以进入到对应的terminal，或者对应的机器

总体这个任务Task Center是做这些事情，结合当前项目，请帮我分析，进行可行性分析，架构设计，同时产出计划，
可以和planning 里面的东西一起分析，参考，原先的内容只是一部分可能。

实现想法：
1. cli-tui应用
2. web端
3. desktop端都需要

目标是轻量级给个人使用。所有的内容都需要写入到新的文档，可以多个文档。比如分析，架构设计，计划这些不同的维度的文档。


## UI实现的问题

实现想法：
1. cli-tui应用
2. web端
3. desktop端都需要

UI 相关结论：

1. CLI/TUI 是 P0 主入口
   - 适合个人本地优先场景
   - 最适合承担快速录入、任务流转、执行观察、终端跳转
   - 应作为第一阶段必须落地的主界面
2. Web 是 P1 观察与管理界面
   - 适合提供 dashboard、筛选、聚合、只读或轻编辑
   - 不适合承载深度 terminal 交互
   - 适合作为第二阶段补充界面
3. Desktop 是 P2 本地增强界面
   - 适合做系统托盘、快捷捕捉、terminal 容器、通知聚合
   - 复杂度最高，不适合作为第一阶段主战场
   - 更适合在本地 API 稳定后再实现

UI 需要实现的模块：

- Capture Entry：快速录入与收集
- Task Inbox：输入任务池
- Workflow Board：按 stage 和 status 的流程看板
- Dispatch Panel：任务分派到 agent / repo / terminal
- Execution Monitor：执行状态、日志、心跳、活动时间
- Review Queue：执行结果确认与回流
- Dashboard Views：按 agent / repo / stage 的聚合视图
- Notification Surface：重要状态变化提醒

文档输出位置：

- 严格评估文档：`docs/analysis/local-task-center-ui-evaluation.md`
- UI 架构文档：`docs/architecture/local-task-center-ui.md`
- UI Spec 索引：`docs/spec/local-task-center/ui-overview.md`
- CLI/TUI Spec：`docs/spec/local-task-center/ui-cli-tui.md`
- Web Spec：`docs/spec/local-task-center/ui-web.md`
- Desktop Spec：`docs/spec/local-task-center/ui-desktop.md`

请对这些内容进行分析哪些需要实现的UI功能和模块。写入到UI相关的Spec中，最后给出一份经过严格评估的文档。
