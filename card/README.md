# Card Builder

Interactive card is rich in formats. However, it takes much efforts to build one. Thus, we provide a declarative card builder to make it easier.

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

## Sending Message

```go
b := lark.NewCardBuilder()
card := b.Card(
    b.Div(
        b.Field(b.Text("Content")).Short(),
    ),
).
    Wathet().
    Title("Card Title")
msg := lark.NewMsgBuffer(lark.MsgInteractive)
om := msg.BindEmail("youremail@example.com").Card(card.String()).Build()
resp, err := bot.PostMessage(om)
```
