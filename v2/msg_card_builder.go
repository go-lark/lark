package lark

import (
	"github.com/go-lark/card-builder"
	"github.com/go-lark/card-builder/i18n"
)

type i18nCardBuilder struct{}

// CardBuilder .
type CardBuilder struct {
	I18N *i18nCardBuilder
}

// Card wraps i18n card
func (i18nCardBuilder) Card(blocks ...*i18n.LocalizedBlock) *i18n.Block {
	return i18n.Card(blocks...)
}

func (i18nCardBuilder) WithLocale(locale string, elements ...card.Element) *i18n.LocalizedBlock {
	return i18n.WithLocale(locale, elements...)
}

// Title wraps i18n title block
func (i18nCardBuilder) LocalizedText(locale, s string) *i18n.LocalizedTextBlock {
	return i18n.LocalizedText(locale, s)
}

// NewCardBuilder .
func NewCardBuilder() *CardBuilder {
	return &CardBuilder{
		I18N: &i18nCardBuilder{},
	}
}

// Card assigns elements
func (CardBuilder) Card(elements ...card.Element) *card.Block {
	return card.Card(elements...)
}

// Action elements including Button, SelectMenu, Overflow, DatePicker, TimePicker, DatetimePicker
func (CardBuilder) Action(actions ...card.Element) *card.ActionBlock {
	return card.Action(actions...)
}

// Button .
func (CardBuilder) Button(text *card.TextBlock) *card.ButtonBlock {
	return card.Button(text)
}

// Confirm .
func (CardBuilder) Confirm(title, text string) *card.ConfirmBlock {
	return card.Confirm(title, text)
}

// DatePicker .
func (CardBuilder) DatePicker() *card.DatePickerBlock {
	return card.DatePicker()
}

// TimePicker .
func (CardBuilder) TimePicker() *card.TimePickerBlock {
	return card.TimePicker()
}

// DatetimePicker .
func (CardBuilder) DatetimePicker() *card.DatetimePickerBlock {
	return card.DatetimePicker()
}

// Div .
func (CardBuilder) Div(fields ...*card.FieldBlock) *card.DivBlock {
	return card.Div(fields...)
}

// Field .
func (CardBuilder) Field(text *card.TextBlock) *card.FieldBlock {
	return card.Field(text)
}

// Hr .
func (CardBuilder) Hr() *card.HrBlock {
	return card.Hr()
}

// Img .
func (CardBuilder) Img(key string) *card.ImgBlock {
	return card.Img(key)
}

// Note .
func (CardBuilder) Note() *card.NoteBlock {
	return card.Note()
}

// Option .
func (CardBuilder) Option(value string) *card.OptionBlock {
	return card.Option(value)
}

// Overflow .
func (CardBuilder) Overflow(options ...*card.OptionBlock) *card.OverflowBlock {
	return card.Overflow(options...)
}

// SelectMenu .
func (CardBuilder) SelectMenu(options ...*card.OptionBlock) *card.SelectMenuBlock {
	return card.SelectMenu(options...)
}

// Text .
func (CardBuilder) Text(s string) *card.TextBlock {
	return card.Text(s)
}

// Markdown .
func (CardBuilder) Markdown(s string) *card.MarkdownBlock {
	return card.Markdown(s)
}

// URL .
func (CardBuilder) URL() *card.URLBlock {
	return card.URL()
}

// ColumnSet column set module
func (CardBuilder) ColumnSet(columns ...*card.ColumnBlock) *card.ColumnSetBlock {
	return card.ColumnSet(columns...)
}

// Column column module
func (CardBuilder) Column(elements ...card.Element) *card.ColumnBlock {
	return card.Column(elements...)
}

// ColumnSetAction column action module
func (CardBuilder) ColumnSetAction(url *card.URLBlock) *card.ColumnSetActionBlock {
	return card.ColumnSetAction(url)
}
