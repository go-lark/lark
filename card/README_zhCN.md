# 声明式卡片构造工具

## 使用方式

```go
b := NewCardBuilder()
c := b.Card(
    b.Markdown("some text"),
).
    Title("title").
    NoForward()

fmt.Println(c.String())
```

会渲染出：

```json
{
  "config": {
    "wide_screen_mode": true,
    "enable_forward": false
  },
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "title"
    }
  },
  "elements": [
    {
      "tag": "markdown",
      "content": "some text"
    }
  ]
}
```

我们把每一个元素（`div`、`text`、`button` 等）都映射成这样的声明形式。

我们约定使用「参数」来表示**元素内部**的子元素（例如 `div` 中的 `fields`），用「链式调用」来设置**元素本身**的属性（例如卡片中的 `forward`）。

详细请参见[msg_card_builder_test.go](./msg_card_builder_test.go)。

## 限制

- 暂不支持 i18n 卡片，如有相关需求请暂时使用 json
- `CardBuilder` 只作为卡片构造相关方法的集合，不具备承载卡片的功能，所以你可以在任意地方使用同一个 `CardBuilder` ，无需每次使用前新建。

## 使用 go-lark 发送

```go
b := lark.NewCardBuilder()
c := b.Card(...Elements)
// 使用 c.String() 或者 c.MarshalJSON() 将卡片内容渲染为 string 或 []byte
msg := lark.NewMsgBuffer(lark.MsgInteractive)
om := msg.BindEmail("youremail@example.com").Card(c.String()).Build()
bot.PostMessage(om)
```
