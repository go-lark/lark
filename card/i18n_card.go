package card

import (
	"encoding/json"
)

var _ Element = (*I18NBlock)(nil)

type I18NBlock struct {
	cnElements []Element
	enElements []Element
	jpElements []Element

	disableForward bool
	updateMulti    bool
	title          I18NCardTitle
	template       string
	links          *URLBlock
}

type I18NCardTitle struct {
	ZhName string
	EnName string
	JpName string
}

// MarshalJSON implements json.Marshaller
func (b *I18NBlock) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(b.Render(), "", "  ")
}

// String implements fmt.Stringer
func (b *I18NBlock) String() string {
	bytes, _ := b.MarshalJSON()
	return string(bytes)
}

type i18NCardRenderer struct {
	Config       cardConfigRenderer     `json:"config,omitempty"`
	Header       i18NCardHeaderRenderer `json:"header,omitempty"`
	CardLink     Renderer               `json:"card_link,omitempty"`
	I18NElements map[string][]Renderer  `json:"i18n_elements,omitempty"`
}

type i18NCardHeaderRenderer struct {
	Title    i18nTitleRenderer `json:"title"`
	Template string            `json:"template,omitempty"` // 卡片标题颜色
}

type i18nTitleRenderer struct {
	Tag      string `json:"tag"` // 固定为plain_text
	I18NName struct {
		ZhName string `json:"zh_cn,omitempty"`
		EnName string `json:"en_us,omitempty"`
		JpName string `json:"ja_jp,omitempty"`
	} `json:"i18n"`
}

func (b *I18NBlock) Render() Renderer {
	ret := i18NCardRenderer{
		Config: cardConfigRenderer{
			WideScreenMode: true,
			EnableForward:  !b.disableForward,
			UpdateMulti:    b.updateMulti,
		},
		Header: i18NCardHeaderRenderer{
			Template: b.template,
			Title: i18nTitleRenderer{
				Tag: "plain_text",
				I18NName: struct {
					ZhName string `json:"zh_cn,omitempty"`
					EnName string `json:"en_us,omitempty"`
					JpName string `json:"ja_jp,omitempty"`
				}{
					b.title.ZhName, b.title.EnName, b.title.JpName,
				},
			},
		},
		I18NElements: map[string][]Renderer{},
	}

	// elements
	if len(b.cnElements) > 0 {
		ret.I18NElements["zh_cn"] = renderElements(b.cnElements)
	}

	if len(b.enElements) > 0 {
		ret.I18NElements["en_us"] = renderElements(b.enElements)
	}

	if len(b.jpElements) > 0 {
		ret.I18NElements["ja_jp"] = renderElements(b.enElements)
	}

	if b.links != nil {
		ret.CardLink = b.links.Render()
	}

	return ret
}

// NewI18NCard 新建一个I18NCard对象
func NewI18NCard() *I18NBlock {
	return &I18NBlock{}
}

// AddCnContent 向卡片中追加中文内容
func (b *I18NBlock) AddCnContent(el ...Element) *I18NBlock {
	b.cnElements = append(b.cnElements, el...)
	return b
}

// AddEnContent 向卡片中追加英文内容
func (b *I18NBlock) AddEnContent(el ...Element) *I18NBlock {
	b.enElements = append(b.enElements, el...)
	return b
}

// AddJpContent 向卡片中追加日文内容
func (b *I18NBlock) AddJpContent(el ...Element) *I18NBlock {
	b.jpElements = append(b.jpElements, el...)
	return b
}

// NoForward 设置后，卡片将不可转发
func (b *I18NBlock) NoForward() *I18NBlock {
	b.disableForward = true
	return b
}

// UpdateMulti set card can be updated
func (b *I18NBlock) UpdateMulti(updateMulti bool) *I18NBlock {
	b.updateMulti = updateMulti
	return b
}

// CnTitle 中文卡片标题
func (b *I18NBlock) CnTitle(title string) *I18NBlock {
	b.title.ZhName = title
	return b
}

// EnTitle 英文卡片标题
func (b *I18NBlock) EnTitle(title string) *I18NBlock {
	b.title.EnName = title
	return b
}

// JpTitle 日文卡片标题
func (b *I18NBlock) JpTitle(title string) *I18NBlock {
	b.title.JpName = title
	return b
}

// Link 设置卡片跳转链接
func (b *I18NBlock) Link(href *URLBlock) *I18NBlock {
	b.links = href
	return b
}

// Blue 设置卡片标题栏颜色（蓝色）
func (b *I18NBlock) Blue() *I18NBlock {
	b.template = "blue"
	return b
}

// Wathet 设置卡片标题栏颜色（浅蓝色）
func (b *I18NBlock) Wathet() *I18NBlock {
	b.template = "wathet"
	return b
}

// Turquoise 设置卡片标题栏颜色（松石绿）
func (b *I18NBlock) Turquoise() *I18NBlock {
	b.template = "turquoise"
	return b
}

// Green 设置卡片标题栏颜色（绿色）
func (b *I18NBlock) Green() *I18NBlock {
	b.template = "green"
	return b
}

// Yellow 设置卡片标题栏颜色（黄色）
func (b *I18NBlock) Yellow() *I18NBlock {
	b.template = "yellow"
	return b
}

// Orange 设置卡片标题栏颜色（橙色）
func (b *I18NBlock) Orange() *I18NBlock {
	b.template = "orange"
	return b
}

// Red 设置卡片标题栏颜色（红色）
func (b *I18NBlock) Red() *I18NBlock {
	b.template = "red"
	return b
}

// Carmine 设置卡片标题栏颜色（洋红色）
func (b *I18NBlock) Carmine() *I18NBlock {
	b.template = "carmine"
	return b
}

// Violet 设置卡片标题栏颜色（紫红色）
func (b *I18NBlock) Violet() *I18NBlock {
	b.template = "violet"
	return b
}

// Purple 设置卡片标题栏颜色（紫色）
func (b *I18NBlock) Purple() *I18NBlock {
	b.template = "purple"
	return b
}

// Indigo 设置卡片标题栏颜色（靛青色）
func (b *I18NBlock) Indigo() *I18NBlock {
	b.template = "indigo"
	return b
}

// Grey 设置卡片标题栏颜色（灰色）
func (b *I18NBlock) Grey() *I18NBlock {
	b.template = "grey"
	return b
}
