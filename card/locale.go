package card

var _ Element = (*I18NBlock)(nil)

// I18NBlock .
type I18NBlock struct {
	locale   string
	elements []Element
}

type localeBlockRenderer struct {
	Elements []Renderer `json:"elements"`
}

// BlockWithLocale creates a block with locale
func BlockWithLocale(locale string, elements ...Element) *I18NBlock {
	return &I18NBlock{
		locale:   locale,
		elements: elements,
	}
}

// Render .
func (l *I18NBlock) Render() Renderer {
	return renderElements(l.elements)
}

// I18NTitleBlock .
type I18NTitleBlock struct {
	titles []*I18NTextBlock
}

// TitleWithLocale .
func TitleWithLocale(titles ...*I18NTextBlock) *I18NTitleBlock {
	return &I18NTitleBlock{
		titles: titles,
	}
}

type i18nTitleBlockRenderer struct {
	ElementTag
	I18N map[string]Renderer `json:"i18n"`
}

// Render .
func (t *I18NTitleBlock) Render() Renderer {
	ret := i18nTitleBlockRenderer{
		ElementTag: ElementTag{
			Tag: "plain_text",
		},
		I18N: make(map[string]Renderer),
	}
	for _, tt := range t.titles {
		ret.I18N[tt.locale] = tt.Render()
	}
	return ret
}

// I18NTextBlock .
type I18NTextBlock struct {
	locale string
	text   string
}

// TextWithLocale .
func TextWithLocale(locale, text string) *I18NTextBlock {
	return &I18NTextBlock{
		locale: locale,
		text:   text,
	}
}

// Render .
func (t *I18NTextBlock) Render() Renderer {
	return t.text
}
