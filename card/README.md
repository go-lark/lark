# Declarative card builder to Go
## Basic Usage
```go
b := NewCardBuilder()
c := b.Card(
    b.Markdown("some text"),
).
    Title("title"). 
    NoForward()

fmt.Println(c.String())
```

will render like below:

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

Which you can use directly in card content.

We mapped every element of the card (div, text, button, etc.) to declarative calls like above.

All inner elements (e.g. `fields` of `div` blocks) are considered as `arguments`,
while all element properties (e.g. `forward` property of `card` blocks) are considered as `chained calls`.

Refer to `../msg_card_builder_test.go` for more examples. 

> NOTE
>
> i18n cards are currently NOT YET SUPPORTED. Use raw json if necessary.
> 
> `CardBuilder` contains ONLY a group of card-builder-related functions and contains NO card content.
> Thus, you can use the same `CardBuilder` whenever building a card instead of making a new one before build.

## Working with go-lark

```go
b := lark.NewCardBuilder()
c := b.Card(...Elements)
// Use c.String() or c.MarshalJSON() to render card content to string or []byte
msg := lark.NewMsgBuffer(lark.MsgInteractive)
om := msg.BindEmail("youremail@example.com").Card(c.String()).Build()
bot.PostMessage(om)
```
