package card

var _ Element = (*OverflowBlock)(nil)

type OverflowBlock struct {
	options []Element
	value   map[string]interface{}
	confirm *ConfirmBlock
}

type overflowRenderer struct {
	ElementTag
	Options []Renderer             `json:"options"`
	Value   map[string]interface{} `json:"value,omitempty"`
	Confirm Renderer               `json:"confirm,omitempty"`
}

func (o *OverflowBlock) Render() Renderer {
	ret := overflowRenderer{
		ElementTag: ElementTag{
			Tag: "overflow",
		},
		Options: RenderElements(o.options),
		Value:   o.value,
	}
	if o.confirm != nil {
		ret.Confirm = o.confirm.Render()
	}
	return ret
}

// Overflow 折叠按钮菜单组件
func Overflow(opt ...*OptionBlock) *OverflowBlock {
	ret := &OverflowBlock{options: make([]Element, len(opt))}
	for i, v := range opt {
		ret.options[i] = v
	}
	return ret
}

// Value 选定后发送给业务方的数据
func (o *OverflowBlock) Value(v map[string]interface{}) *OverflowBlock {
	o.value = v
	return o
}

// Confirm 选定后二次确认的弹框
func (o *OverflowBlock) Confirm(title, text string) *OverflowBlock {
	o.confirm = Confirm(title, text)
	return o
}
