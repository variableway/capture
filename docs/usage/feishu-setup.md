# 飞书 Bot 配置指南

## 概述

Capture 支持通过飞书 Bot 接收消息，自动将飞书中的想法转化为任务。

## 前置条件

- 飞书企业版账号
- 飞书开放平台开发者权限

## 步骤 1: 创建飞书应用

1. 访问 [飞书开放平台](https://open.feishu.cn/app)
2. 点击「创建企业自建应用」
3. 填写应用名称（如 "Capture"）和描述
4. 记录 **App ID** 和 **App Secret**

## 步骤 2: 启用机器人功能

1. 进入应用管理页面
2. 左侧菜单 → 「应用功能」→「机器人」
3. 开启机器人功能
4. 设置机器人名称和头像

## 步骤 3: 配置权限

在「权限管理」页面，添加以下权限并申请发布：

| 权限 | 权限标识 | 用途 |
|------|----------|------|
| 获取与发送单聊、群组消息 | `im:message` | 接收和发送消息 |
| 接收消息事件 | `im:message.receive_v1` | 实时接收用户消息 |
| 读取多维表格 | `bitable:app:readonly` | 读取任务数据 |
| 读写多维表格 | `bitable:app` | 同步任务到多维表格 |

## 步骤 4: 配置事件订阅

### WebSocket 模式（推荐，适合本地开发）

1. 进入「事件订阅」页面
2. 选择「使用长连接接收事件」
3. 添加事件：`im.message.receive_v1`（接收消息）
4. 点击保存

**优点**：无需公网 IP，无需域名，无需 HTTPS

### Webhook 模式（适合生产部署）

1. 进入「事件订阅」页面
2. 选择「将事件发送至请求地址」
3. 填写请求 URL：`https://your-domain.com/webhook/feishu`
4. 验证通过后，添加事件：`im.message.receive_v1`
5. 记录 **Verification Token** 和 **Encrypt Key**

**要求**：
- 公网可访问的 URL
- HTTPS 协议
- 可用 ngrok/frp 进行本地开发调试

## 步骤 5: 配置多维表格（可选）

如果需要同步任务到飞书多维表格：

1. 在飞书中创建一个新的多维表格
2. 添加以下列：

| 列名 | 类型 | 说明 |
|------|------|------|
| task_id | 文本 | 任务 ID |
| title | 文本 | 任务标题 |
| status | 单选 | 状态值: todo, in_progress, done, cancelled, archived |
| priority | 单选 | 优先级: high, medium, low |
| source | 文本 | 来源 |
| created_at | 日期 | 创建时间 |
| updated_at | 日期 | 更新时间 |

3. 从多维表格 URL 中获取：
   - **App Token**：URL 中 `/base/` 后面的字符串
   - **Table ID**：表格标签页的 ID

## 步骤 6: 发布应用

1. 在「版本管理与发布」页面创建新版本
2. 提交审核（企业内部应用通常自动通过）
3. 审核通过后发布

## 步骤 7: 启动 Bot 服务

### 设置环境变量

```bash
# 必需
export FEISHU_APP_ID="cli_xxxxxxxxxx"
export FEISHU_APP_SECRET="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

# Webhook 模式需要
export FEISHU_VERIFICATION_TOKEN="xxxxxxxxxxxxxxxxxxxxxxxx"
export FEISHU_ENCRYPT_KEY="xxxxxxxxxxxxxxxxxxxxxxxx"

# 多维表格同步需要
export FEISHU_BITABLE_APP_TOKEN="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
export FEISHU_BITABLE_TABLE_ID="tblxxxxxxxxxxxxxx"
```

### 启动 Bot

```bash
# WebSocket 模式（推荐本地开发）
capture bot serve --mode websocket

# Webhook 模式（生产部署）
capture bot serve --mode webhook --port 8080
```

### 使用 systemd 管理（Linux）

```ini
# /etc/systemd/system/capture-bot.service
[Unit]
Description=Capture Feishu Bot
After=network.target

[Service]
Type=simple
User=capture
WorkingDirectory=/opt/capture
ExecStart=/opt/capture/capture bot serve --mode websocket
Restart=always
RestartSec=5
Environment=FEISHU_APP_ID=cli_xxx
Environment=FEISHU_APP_SECRET=xxx

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl enable capture-bot
sudo systemctl start capture-bot
```

## 使用飞书 Bot

在飞书中给机器人发送消息：

| 命令 | 说明 | 示例 |
|------|------|------|
| 记录/添加/新建 | 创建任务 | `记录 优化构建脚本 #优化` |
| 列出/查看 | 查看任务列表 | `列出` |
| 删除 | 删除任务 | `删除 TASK-00001` |
| 帮助 | 显示帮助 | `帮助` |

### 示例

```
用户: 记录 学习 Go 语言并发编程 #学习 #Go 优先级：高
Bot:  已创建: TASK-00005 - 学习 Go 语言并发编程

用户: 列出
Bot:  共 3 个任务:
        TASK-00005 [todo] 学习 Go 语言并发编程
        TASK-00004 [in_progress] 优化项目构建脚本
        TASK-00001 [done] 测试第一个想法

用户: 删除 TASK-00001
Bot:  已删除: TASK-00001

用户: 帮助
Bot:  Capture Bot 命令：
      - 记录 <内容>  — 创建新任务
      - 列出 — 查看所有任务
      - 删除 <TASK-ID> — 删除任务
      - 帮助 — 显示此帮助
```

## 故障排除

### Bot 无法接收消息

1. 检查应用是否已发布
2. 检查事件订阅是否正确配置
3. 检查权限是否已申请
4. WebSocket 模式检查网络连接
5. Webhook 模式检查 URL 是否可访问

### Bot 无法发送消息

1. 检查 `im:message` 权限
2. 检查 App ID 和 Secret 是否正确

### 多维表格同步失败

1. 检查 `bitable:app` 权限
2. 检查 App Token 和 Table ID 是否正确
3. 检查多维表格列名是否与要求一致
4. 检查应用是否有多维表格的访问权限

### 本地开发 Webhook 调试

使用 ngrok 暴露本地端口：

```bash
# 安装 ngrok
brew install ngrok

# 暴露本地端口
ngrok http 8080

# 将 ngrok 提供的 URL 填入飞书事件订阅
# 例如: https://xxxx.ngrok.io/webhook/feishu
```
