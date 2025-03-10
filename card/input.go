package card

var _ Element = (*InputBlock)(nil)

// InputBlock 输入元素
type InputBlock struct {
	name          string
	placeholder   string
	inputType     string
	label         string
	labelPosition string
	maxLength     int
	rows          int
	maxRows       int
	autoResize    bool
	disabled      bool
	defaultValue  string
	width         string
}

type InputBlockRenderer struct {
	ElementTag
	Name          string   `json:"name"`
	Placeholder   Renderer `json:"placeholder,omitempty"`
	LabelPosition string   `json:"label_position"`
	Label         Renderer `json:"label,omitempty"`
	MaxLength     int      `json:"max_length"`
	InputType     string   `json:"input_type,omitempty"`
	Rows          int      `json:"rows,omitempty"`
	MaxRows       int      `json:"max_rows,omitempty"`
	AutoResize    bool     `json:"auto_resize,omitempty"`
	Disabled      bool     `json:"disabled,omitempty"`
	DefaultValue  string   `json:"default_value,omitempty"`
	Width         string   `json:"width,omitempty"`
}

// Render 渲染为 Renderer
func (s *InputBlock) Render() Renderer {
	ret := InputBlockRenderer{
		ElementTag: ElementTag{
			Tag: "input",
		},
		Name:          s.name,
		LabelPosition: s.labelPosition,
		MaxLength:     s.maxLength,
		Rows:          s.rows,
		Label:         Text(s.label).Render(),
		Placeholder:   Text(s.placeholder).Render(),
		InputType:     s.inputType,
		MaxRows:       s.maxRows,
		AutoResize:    s.autoResize,
		Disabled:      s.disabled,
		DefaultValue:  s.defaultValue,
		Width:         s.width,
	}
	return ret
}

// Input 输入组件
func Input(name string) *InputBlock {
	return &InputBlock{
		name:          name,
		label:         name,
		labelPosition: "top",
		placeholder:   name,
	}
}

// Placeholder 默认展示内容
func (s *InputBlock) Placeholder(str string) *InputBlock {
	s.placeholder = str
	return s
}

// Label 输入框前的标签
func (s *InputBlock) Label(str string) *InputBlock {
	s.label = str
	return s
}

// LabelPosition 标签位置
func (s *InputBlock) LabelPosition(p string) *InputBlock {
	s.labelPosition = p
	return s
}

// MaxLength 最大长度
func (s *InputBlock) MaxLength(l int) *InputBlock {
	s.maxLength = l
	return s
}

// Rows 输入框行数
func (s *InputBlock) Rows(r int) *InputBlock {
	s.rows = r
	if s.maxRows < r {
		s.maxRows = r
	}
	return s
}

// MaxRows 最大行数
func (s *InputBlock) MaxRows(r int) *InputBlock {
	s.maxRows = r
	if s.rows > r {
		s.rows = r
	}
	return s
}

// InputType 输入框类型
func (s *InputBlock) InputType(t string) *InputBlock {
	s.inputType = t
	return s
}

// Width 宽度
func (s *InputBlock) Width(w string) *InputBlock {
	s.width = w
	return s
}

// AutoResize 自动调整高度
func (s *InputBlock) AutoResize() *InputBlock {
	s.autoResize = true
	return s
}

// Disabled 是否禁用
func (s *InputBlock) Disabled() *InputBlock {
	s.disabled = true
	return s
}

// DefaultValue 默认值
func (s *InputBlock) DefaultValue(v string) *InputBlock {
	s.defaultValue = v
	return s
}
