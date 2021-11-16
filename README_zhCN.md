# go-lark

[![build](https://github.com/go-lark/lark/actions/workflows/ci.yml/badge.svg)](https://github.com/go-lark/lark/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/go-lark/lark/branch/main/graph/badge.svg)](https://codecov.io/gh/go-lark/lark)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-lark/lark.svg)](https://pkg.go.dev/github.com/go-lark/lark)

一个简单、开发者友好的 Lark 开放平台机器人 SDK。

## 介绍

go-lark 主要实现了消息类 API，提供完整的聊天机器人和通知机器人支持。在字节跳动公司内部得到广泛应用，有大约 450 开发者和超过 1500 个 Go 仓库使用。

## 功能

- 聊天机器人和通知机器人
- 发送各类消息（群发、私聊、富文本、卡片消息）
- 快速消息体构造 `MsgBuffer`
- 一站式解决服务器 Challenge 和聊天消息响应
- 支持加密和校验
- 支持 [Gin](https://github.com/go-lark/lark-gin) 框架中间件
- 文档、测试覆盖

## 安装

```shell
go get github.com/go-lark/lark
```

## 快速入门

### 前置准备

我们支持两种类型的机器人，需要分别用以下方式创建：

聊天机器人：

- 飞书: 通过[飞书开放平台](https://open.feishu.cn/)创建。
- Lark: 通过 [Lark Developer](https://open.larksuite.com/) 创建。
- 需要 App ID 和 App Secret 来初始化 `ChatBot`。

通知机器人：

- 通过群聊创建-群机器人创建。
- 需要使用 WebHook URL。

### 消息发送

聊天机器人：

```go
import "github.com/go-lark/lark"

func main() {
    bot := lark.NewChatBot("<App ID>", "<App Secret>")
    bot.StartHeartbeat()
    bot.PostText("hello, world", lark.WithEmail("someone@example.com"))
}
```

通知机器人：

```go
import "github.com/go-lark/lark"

func main() {
    bot := lark.NewNotificationBot("WEB HOOK URL")
    bot.PostNotification("go-lark", "example")
}
```

## 限制

- go-lark 基于飞书域名进行测试，理论上可以完全兼容 Lark 平台（API 定义一致）。但我们不保证在 Lark 下完全可用，因为账户限于，没有专门测试过。
- go-lark 仅支持企业自建应用，不支持应用商店应用（ISV）。
- go-lark 仅实现了机器人和消息 API，对于飞书文档、日历等功能，并不支持。
- go-lark 目前实现的是 API v3/v4 版本*（官方文档通常还会出现 im/v1 版本）以及事件 Schema 1.0 版本。

### 切换到 Lark 域名

go-lark 默认使用飞书 API 域名，我们需要调用`SetDomain`来切换到 Lark：

```go
bot := lark.NewChatBot("<App ID>", "<App Secret>")
bot.SetDomain(lark.DomainLark)
```

## 用法

### 鉴权

自动更新授权：

```go
// initialize a chat bot with appID and appSecret
bot := lark.NewChatBot(appID, appSecret)
// Renew access token periodically
bot.StartHeartbeat()
// Stop renewal
bot.StopHeartbeat()
```

单次授权：

```go
bot := lark.NewChatBot(appID, appSecret)
resp, err := bot.GetTenantAccessTokenInternal(true)
// and we can now access the token value with `bot.tenantAccessToken()`
```

参考实例：[鉴权](https://github.com/go-lark/examples/tree/main/auth)

### 消息

简单消息可以以下接口直接通过：

- `PostText`
- `PostTextMention`
- `PostTextMentionAll`
- `PostImage`
- `PostShareChatCard`

参考实例：[基本消息](https://github.com/go-lark/examples/tree/main/basic-message)。

对于复杂消息，可以使用 [Message Buffer](#message-buffer) 进行链式构造。

### 参考实例

- [鉴权](https://github.com/go-lark/examples/tree/main/auth)
- [基本消息](https://github.com/go-lark/examples/tree/main/basic-message)
- [图片消息](https://github.com/go-lark/examples/tree/main/image-message)
- [富文本消息](https://github.com/go-lark/examples/tree/main/rich-text-message)
- [分享群卡片](https://github.com/go-lark/examples/tree/main/share-chat)
- [交互卡片](https://github.com/go-lark/examples/tree/main/interactive-message)
- [群操作](https://github.com/go-lark/examples/tree/main/group)

### Message Buffer

发送消息需要先通过 MsgBuffer 构造消息体，然后调用 `PostMessage` 进行发送。

MsgBuffer 支持多种类型的消息：

- `MsgText`：文本
- `MsgPost`：富文本
- `MsgImage`：图片
- `MsgShareCard`：群名片
- `MsgInteractive`：交互式卡片

MsgBuffer 主要有两类函数，Bind 函数和内容函数。

Bind 函数：

| 函数       | 作用        | 备注                                  |
| ---------- | ----------- | ------------------------------------- |
| BindChatID | 绑定 ChatID | OpenID/UserID/Email/ChatID 四选一即可 |
| BindOpenID | 绑定 OpenID |                                       |
| BindUserID | 绑定 UserID |                                       |
| BindEmail  | 绑定邮箱    |                                       |
| BindReply  | 绑定回复    | 回复他人时需要                        |

内容函数大多跟消息类型是强关联的，类型错误不会生效。内容函数：

| 函数      | 适用范围         | 作用           | 备注                                                                                         |
| --------- | ---------------- | -------------- | -------------------------------------------------------------------------------------------- |
| Text      | `MsgText`        | 添加文本内容   | 可使用 `TextBuilder` 构造                                                                    |
| Post      | `MsgPost`        | 添加富文本内容 | 可使用 `PostBuilder` 构造                                                                    |
| Image     | `MsgImage`       | 添加图片       | 图片需要先上传到 飞书服务器                                                                  |
| ShareChat | `MsgShareCard`   | 添加分享群卡片 |                                                                                              |
| Card      | `MsgInteractive` | 添加交互式卡片 | 非国际化卡片可使用 `CardBuilder` 构造，详见[声明式卡片搭建工具 to Go](card/README_zhCN.md) |

### 异常处理

每个 API 都会返回 `response` 和 `error`。`error` 是 HTTP 客户端返回，`response` 是开放平台接口返回。一般来说，每个接口的 `response` 都会有 `code` 字段，如果非 0 则表示有错误。具体错误码含义，请查看[官方文档](https://open.feishu.cn/document/ukTMukTMukTM/ugjM14COyUjL4ITN)。

## 事件处理

事件是飞书机器人用于实现机器人交互的机制，创建聊天机器人后我们并不具有和机器人交互的能力，需要通过开放平台的挑战和消息相应完成交互。

飞书开放平台提供多种事件，并且有两种版本的格式（1.0 和 2.0）。go-lark 只实现了 1.0 中的两种事件。

在开发交互机器人过程中，我们主要需要用到这两类事件：

- URL 挑战
- 接收消息

我们推荐使用 Gin 中间件处理事件。

### [Gin Middleware](https://github.com/go-lark/lark-gin)

实例：[examples/gin-middleware](https://github.com/go-lark/examples/tree/main/gin-middleware)

#### URL 挑战

```go
r := gin.Default()
middleware := larkgin.NewLarkMiddleware()
middleware.BindURLPrefix("/handle") // 假设 URL 是 http://your.domain.com/handle
r.Use(middleware.LarkChallengeHandler())
```

#### 接收消息

```go
r := gin.Default()
middleware := larkgin.NewLarkMiddleware()
middleware.BindURLPrefix("/handle") // supposed URL is http://your.domain.com/handle
r.POST("/handle", func(c *gin.Context) {
    if msg, ok := middleware.GetMessage(c); ok && msg != nil {
        text := msg.Event.Text
        // 你的业务逻辑
    }
})
```

### 加密安全

飞书开放平台目前有两种加密安全策略（可以同时启用），分别是 AES 加密和 Token 校验。

- AES 加密：需要在验证 Challenge 时就开启，此后所有收到的消息都会走 AES 加密。
- Token 校验：验证消息来自 Lark 开放平台。

我们建议开启 Token 校验。如果没有使用 HTTPS 协议，则开启 AES。

```go
middleware.WithTokenVerfication("<verification-token>")
middleware.WithEncryption("<encryption-key>")
```

### 调试

飞书官方没有提供发消息工具，如果测试消息交互的话不得不在飞书上发消息，直接在“线上” URL 调试，很不方便。

我们加入了线下模拟消息事件的 `PostEvent`，通过它可以在任何地方进行调试。当然，模拟消息的包体需要自己构造。同时，我们也可以使用 `PostEvent` 对原消息进行转发，对消息进行反向代理。

参考实例：[examples/event-forward](https://github.com/go-lark/examples/tree/main/event-forward)

> `PostEvent`目前不支持 AES 加密。

## FAQ

- 调用接口发消息报错，错误码 99991401
  - 在开发者后台“安全”中取消“IP 白名单”
- 机器人发消息失败了
  - 常见原因：1，忘了开启授权；2，没进群发群消息；3，其它权限类问题
- go-lark 可以发消息卡片吗？怎么发？
  - 可以，但需要自己构造拼接消息 JSON body，然后使用`PostMessage`发送。

## 贡献

- 如果在使用 go-lark 时遇到 Bug，请提交 Issue。
- 欢迎通过 Pull Request 提交功能或 Bug 修复。

## 协议

Copyright (c) David Zhang, 2018-2021. Licensed under MIT License.
