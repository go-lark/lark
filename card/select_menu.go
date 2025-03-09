package card

var _ Element = (*SelectMenuBlock)(nil)

// SelectMenuBlock 菜单元素
type SelectMenuBlock struct {
	name          string
	tag           string
	placeholder   string
	initialOption string
	options       []Element
	value         map[string]interface{}
	required      bool
	width         string
	confirm       *ConfirmBlock
}

type selectMenuRenderer struct {
	ElementTag
	Name          string                 `json:"name,omitempty"`
	Placeholder   Renderer               `json:"placeholder,omitempty"`
	InitialOption string                 `json:"initial_option,omitempty"`
	Options       []Renderer             `json:"options,omitempty"`
	Value         map[string]interface{} `json:"value,omitempty"`
	Width         string                 `json:"width,omitempty"`
	Confirm       Renderer               `json:"confirm,omitempty"`
	Required      bool                   `json:"required,omitempty"`
}

// Render 渲染为 Renderer
func (s *SelectMenuBlock) Render() Renderer {
	ret := selectMenuRenderer{
		ElementTag: ElementTag{
			Tag: s.tag,
		},
		InitialOption: s.initialOption,
		Options:       renderElements(s.options),
		Value:         s.value,
		Placeholder:   Text(s.placeholder).Render(),
		Name:          s.name,
		Required:      s.required,
	}
	if s.confirm != nil {
		ret.Confirm = s.confirm.Render()
	}
	return ret
}

// SelectMenu 菜单组件
func SelectMenu(opt ...*OptionBlock) *SelectMenuBlock {
	ret := &SelectMenuBlock{
		tag:     "select_static",
		options: make([]Element, len(opt)),
	}
	for i, v := range opt {
		ret.options[i] = v
	}
	return ret
}

// Name 菜单的标识
func (s *SelectMenuBlock) Name(n string) *SelectMenuBlock {
	s.name = n
	return s
}

// SelectPerson 选人模式，value应设置为人员的open_id，options 为空则候选人员为当前群组
func (s *SelectMenuBlock) SelectPerson() *SelectMenuBlock {
	s.tag = "select_person"
	return s
}

// InitialOption 默认选项的 value 字段值
func (s *SelectMenuBlock) InitialOption(o string) *SelectMenuBlock {
	s.initialOption = o
	return s
}

// Placeholder 未选中时展示的内容，无默认选项时必须设置
func (s *SelectMenuBlock) Placeholder(p string) *SelectMenuBlock {
	s.placeholder = p
	return s
}

// Value 选中后发送给业务方的数据
func (s *SelectMenuBlock) Value(v map[string]interface{}) *SelectMenuBlock {
	s.value = v
	return s
}

// Width 宽度
func (s *SelectMenuBlock) Width(w string) *SelectMenuBlock {
	s.width = w
	return s
}

// Confirm 选中后二次确认的弹框
func (s *SelectMenuBlock) Confirm(title, text string) *SelectMenuBlock {
	s.confirm = Confirm(title, text)
	return s
}

// Required 是否必填
func (s *SelectMenuBlock) Required(r bool) *SelectMenuBlock {
	s.required = r
	return s
}
