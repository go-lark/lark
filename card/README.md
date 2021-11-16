# Declarative Card Builder

## Getting Started

```go
b := NewCardBuilder()
c := b.Card(
    b.Markdown("some text"),
).
    Title("title").
    NoForward()

fmt.Println(c.String())
```

which will render as following:

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

We map every element of card (`div`, `text`, `button`, etc.) to declarative calls as shown in the demo.

All inner elements (e.g. `fields` of `div` blocks) are considered as arguments,
while all element properties (e.g. `forward` property of `card` blocks) are considered as chained calls.

Refer to [msg_card_builder_test.go](./msg_card_builder_test.go) for details.

## Limits

- i18n cards are currently NOT YET SUPPORTED. Use raw json if necessary.
- `CardBuilder` contains ONLY a group of card-builder-related functions and contains NO card content. Thus, you can use the same `CardBuilder` whenever building a card instead of making a new one before build.

## Works with go-lark

```go
b := lark.NewCardBuilder()
c := b.Card(...Elements)
// Use c.String() or c.MarshalJSON() to render card content to string or []byte
msg := lark.NewMsgBuffer(lark.MsgInteractive)
om := msg.BindEmail("youremail@example.com").Card(c.String()).Build()
bot.PostMessage(om)
```
