package lark

import "github.com/nlimpid/lark/card"

// CardBlock 卡片元素
type CardBlock = card.Block

// CardBuilder 卡片构造方法
type CardBuilder interface {
	Card(elements ...card.Element) *CardBlock
	Action(actions ...card.Element) *card.ActionBlock
	Button(text *card.TextBlock) *card.ButtonBlock
	Confirm(title, text string) *card.ConfirmBlock
	DatePicker() *card.DatePickerBlock
	TimePicker() *card.TimePickerBlock
	DatetimePicker() *card.DatetimePickerBlock
	Div(fields ...*card.FieldBlock) *card.DivBlock
	Field(text *card.TextBlock) *card.FieldBlock
	Hr() *card.HrBlock
	Img(key string) *card.ImgBlock
	Note() *card.NoteBlock
	Option(value string) *card.OptionBlock
	Overflow(options ...*card.OptionBlock) *card.OverflowBlock
	SelectMenu(options ...*card.OptionBlock) *card.SelectMenuBlock
	Text(s string) *card.TextBlock
	Markdown(s string) *card.MarkdownBlock
	URL() *card.URLBlock
}

type cardBuilder struct{}

// NewCardBuilder 新建卡片构造器
func NewCardBuilder() CardBuilder {
	return cardBuilder{}
}

// Card 包裹了最外层的卡片结构
func (cardBuilder) Card(elements ...card.Element) *CardBlock {
	return card.Card(elements...)
}

// Action 交互元素，可添加 Button, SelectMenu, Overflow, DatePicker, TimePicker, DatetimePicker
func (cardBuilder) Action(actions ...card.Element) *card.ActionBlock {
	return card.Action(actions...)
}

// Button 按钮交互元素
func (cardBuilder) Button(text *card.TextBlock) *card.ButtonBlock {
	return card.Button(text)
}

// Confirm 用于交互元素的二次确认
func (cardBuilder) Confirm(title, text string) *card.ConfirmBlock {
	return card.Confirm(title, text)
}

// DatePicker 日期选择器
func (cardBuilder) DatePicker() *card.DatePickerBlock {
	return card.DatePicker()
}

// TimePicker 时间选择器
func (cardBuilder) TimePicker() *card.TimePickerBlock {
	return card.TimePicker()
}

// DatetimePicker 日期时间选择器
func (cardBuilder) DatetimePicker() *card.DatetimePickerBlock {
	return card.DatetimePicker()
}

// Div 内容模块
func (cardBuilder) Div(fields ...*card.FieldBlock) *card.DivBlock {
	return card.Div(fields...)
}

// Field 内容模块的排版元素
func (cardBuilder) Field(text *card.TextBlock) *card.FieldBlock {
	return card.Field(text)
}

// Hr 分割线模块
func (cardBuilder) Hr() *card.HrBlock {
	return card.Hr()
}

// Img 图片展示模块
func (cardBuilder) Img(key string) *card.ImgBlock {
	return card.Img(key)
}

// Note 备注模块
func (cardBuilder) Note() *card.NoteBlock {
	return card.Note()
}

// Option 选项模块，可用于 SelectMenu 和 Overflow
func (cardBuilder) Option(value string) *card.OptionBlock {
	return card.Option(value)
}

// Overflow 折叠按钮菜单组件
func (cardBuilder) Overflow(options ...*card.OptionBlock) *card.OverflowBlock {
	return card.Overflow(options...)
}

// SelectMenu 菜单组件
func (cardBuilder) SelectMenu(options ...*card.OptionBlock) *card.SelectMenuBlock {
	return card.SelectMenu(options...)
}

// Text 文本模块
func (cardBuilder) Text(s string) *card.TextBlock {
	return card.Text(s)
}

// Markdown 单独使用的 Markdown 文本模块
func (cardBuilder) Markdown(s string) *card.MarkdownBlock {
	return card.Markdown(s)
}

// URL 链接模块
func (cardBuilder) URL() *card.URLBlock {
	return card.URL()
}
