package card

var _ Element = (*ConfirmBlock)(nil)

type ConfirmBlock struct {
	title, text string
}

type confirmRenderer struct {
	Title Renderer `json:"title"`
	Text  Renderer `json:"text"`
}

func (c *ConfirmBlock) Render() Renderer {
	return confirmRenderer{
		Title: Text(c.title).Render(),
		Text:  Text(c.text).Render(),
	}
}

// Confirm 用于交互元素的二次确认
func Confirm(title, text string) *ConfirmBlock {
	return &ConfirmBlock{
		title: title,
		text:  text,
	}
}
