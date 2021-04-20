# go-lark

[![build](https://github.com/go-lark/lark/actions/workflows/ci.yml/badge.svg)](https://github.com/go-lark/lark/actions/workflows/ci.yml)

go-lark is an easy-to-use unofficial SDK for Feishu and Lark Open Platform.

go-lark implements messaging APIs, with full-fledged supports on building Chat Bot and Notification Bot.

It is widely used and tested by in-house ~450 developers with over 1.5k Go packages.

## Features

- Notification bot & chat bot supported
- Send messages (Group, Private, Rich Text, and Card)
- Quick to build message with `MsgBuffer`
- Easy to create incoming message hook
- Encryption and Token Verification supported
- Middleware support for Gin web framework
- Documentation & tests

## Installation

```shell
go get github.com/go-lark/lark
```

## Quick Start

### Prerequisite

There are two types of bot that is supported by go-lark. We need to create a bot manually.

Chat Bot:

- Feishu: create from [Feishu Open Platform](https://open.feishu.cn/).
- Lark: create from [Lark Developer](https://open.larksuite.com/).
- App ID and App Secret are required to init a `ChatBot`.

Notification Bot:

- Create from group chat.
- Web Hook URL is required.

### Sending Message

Chat Bot:

```go
import "github.com/go-lark/lark"

func main() {
    bot := lark.NewChatBot("<App ID>", "<App Secret>")
    bot.StartHeartbeat()
    bot.PostText("hello, world", lark.WithEmail("someone@example.com"))
}
```

Notification Bot:

```go
import "github.com/go-lark/lark"

func main() {
    bot := lark.NewNotificationBot("WEB HOOK URL")
    bot.PostNotification("go-lark", "example")
}
```

Feishu/Lark API offers more features, please refers to [Usage](#usage) for further documentation.

## Limits

- go-lark is tested on Feishu endpoints, which literally compats Lark endpoints,
  because Feishu and Lark basically shares the same API specification.
  We do not guarantee all of the APIs work well with Lark, until we have tested it on Lark.
- go-lark only supports Custom App. Marketplace App is not supported yet.
- go-lark implements bot and messaging API, other APIs such as Lark Doc, Calendar and so so are not supported.

### Switch to Lark Endpoints

The default API endpoints are for Feishu, in order to switch to Lark, we should use `SetDomain`:

```go
bot := lark.NewChatBot("<App ID>", "<App Secret>")
bot.SetDomain("https://open.larksuite.com")
```

## Usage

### Auth

Auto-renewable authentication:

```go
// initialize a chat bot with appID and appSecret
bot := lark.NewChatBot(appID, appSecret)
// Renew access token periodically
bot.StartHeartbeat()
// Stop renewal
bot.StopHeartbeat()
```

Single-pass authentication:

```go
bot := lark.NewChatBot(appID, appSecret)
resp, err := bot.GetTenantAccessTokenInternal(true)
// and we can now access the token value with `bot.tenantAccessToken()`
```

Example: [examples/auth](/examples/auth)

### Messaging

For Chat Bot, we can send simple messages with the following method:

- `PostText`
- `PostTextMention`
- `PostTextMentionAll`
- `PostImage`
- `PostShareChatCard`

Basic message examples: [examples/basic-message](/examples/basic-message)

To build rich messages, we may use [Message Buffer](#message-buffer) (or simply `MsgBuffer`),
which builds message conveniently with chaining methods.

### Examples

Apart from the general auth and messaging chapter, there are comprehensive examples for almost all APIs.
Here is a collection of ready-to-run examples for each part of `go-lark`:

- [examples/auth](/examples/auth)
- [examples/basic-message](/examples/basic-message)
- [examples/image-message](/examples/image-message)
- [examples/rich-text-message](/examples/rich-text-message)
- [examples/share-chat](/examples/share-chat)
- [examples/interactive-message](/examples/interactive-message)
- [examples/group](/examples/group)
- [examples/user](/examples/user)

### Message Buffer

We can build message body with `MsgBuffer` and send with `PostMessage`, which supports the following message types:

- `MsgText`: Text
- `MsgPost`: Rich Text
- `MsgImage`: Image
- `MsgShareCard`: Group Share Card
- `MsgInteractive`: Interactive Card

`MsgBuffer` provides binding functions and content functions.

Binding functions:

| Function   | Usage               | Comment                                                          |
| ---------- | ------------------- | ---------------------------------------------------------------- |
| BindChatID | Bind a chat ID      | Either `OpenID`, `UserID`, `Email` or `ChatID` should be present |
| BindOpenID | Bind a user open ID |                                                                  |
| BindUserID | Bind a user ID      |                                                                  |
| BindEmail  | Bind a user email   |                                                                  |
| BindReply  | Bind a reply ID     | Required when reply other message                                |

Content functions pair with message content types. If it mismatched, it would not have sent successfully.
Content functions:

| Function  | Message Type     | Usage                   | Comment                                                                                          |
| --------- | ---------------- | ----------------------- | ------------------------------------------------------------------------------------------------ |
| Text      | `MsgText`        | Append plain text       | May build with `TextBuilder`                                                                     |
| Post      | `MsgPost`        | Append rich text        | May build with `PostBuilder`                                                                     |
| Image     | `MsgImage`       | Append image            | Need to upload to Lark server in advance                                                         |
| ShareChat | `MsgShareCard`   | Append group share card |                                                                                                  |
| Card      | `MsgInteractive` | Append interactive card | No `InteractiveBuilder` yet. You may try [Cardbuilder](https://open.feishu.cn/tool/cardbuilder). |

### Error Handling

Each `go-lark` API function returns `response` and `err`.
`err` is the error from HTTP client, when it was not `nil`, HTTP might have gone wrong.

While `response` is HTTP response from Lark API server, in which `Code` and `Ok` represent whether it succeeds.
The meaning of `Code` is defined [here](https://open.feishu.cn/document/ukTMukTMukTM/ugjM14COyUjL4ITN).

### Event

Lark provides a number of [events](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM).
go-lark now only implements two of them, which are needed for interacting between bot and Lark server:

- URL Challenge
- Receiving Messages

We recommend Gin middleware to handle these events.

### Gin Middleware

Example: [examples/gin-middleware](/examples/gin-middleware)

#### URL Challenge

```go
r := gin.Default()
middleware := larkgin.NewLarkMiddleware()
middleware.BindURLPrefix("/handle") // suppose URL is http://your.domain.com/handle
r.Use(middleware.LarkChallengeHandler())

```

#### Receiving Message

```go
r := gin.Default()
middleware := larkgin.NewLarkMiddleware()
r.POST("/handle", func(c *gin.Context) {
    if msg, ok := middleware.GetMessage(c); ok && msg != nil {
        text := msg.Event.Text
        // your awesome logic
    }
})
```

### Security & Encryption

Lark Open Platform offers AES encryption and token verification to ensure security for events.

- AES Encryption: when switch on, all traffic will be encrypted with AES.
- Token Verification: simple token verification for incoming messages.

We also recommend Gin middleware to handle encryption.

### Debugging

Lark does not provide messaging API debugger officially. Thus, we have to debug with real Lark conversation.
We add `PostEvent` to simulate message sending to make it easier.
`PostEvent` can also be used to redirect message, which acts like a reverse proxy.

Example: [examples/event-forward](/examples/event-forward)

> Notice: `PostEvent` does not support AES encryption at the moment.

## Compatibility

go-lark is based on Lark API v3 and v4. The compatibility of each API depends on actual implementation.
Compatibility table:

| API scope        | Compatibility   | Migration              |
| ---------------- | --------------- | ---------------------- |
| auth             | N/A             | Lark does not offer v4 |
| messaging        | v4 only         |                        |
| interactive card | v4 only         |                        |
| bot              | N/A             | Lark does not offer v4 |
| user             | v4 only         |                        |
| group            | partial v3 & v4 | WIP                    |

## FAQ

- I got `99991401` when sending messages
  - remove IP Whitelist from dashboard
- My bot failed sending messages
  1. check authentication.
  2. not invite to the group.
  3. API permission not applied.
- Does go-lark support interactive message card?
  - It can send but no card builder provided. You may need to build it manually and send with go-lark.

## Contributing

- If you think you've found a bug with go-lark, please file an issue.
- Pull Request is welcomed.

## License

Copyright (c) David Zhang, 2018-2021. Licensed under MIT License.
