package card

var _ Element = (*HrBlock)(nil)

type HrBlock struct{}

// Hr 分割线模块
func Hr() *HrBlock {
	return &HrBlock{}
}

func (h *HrBlock) Render() Renderer {
	return ElementTag{
		Tag: "hr",
	}
}
