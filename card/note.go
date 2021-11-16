package card

var _ Element = (*NoteBlock)(nil)

type NoteBlock struct {
	elements []Element
}

type noteRenderer struct {
	ElementTag
	Elements []Renderer `json:"elements"`
}

func (n *NoteBlock) Render() Renderer {
	return noteRenderer{
		ElementTag: ElementTag{
			Tag: "note",
		},
		Elements: RenderElements(n.elements),
	}
}

// Note 备注模块
func Note() *NoteBlock {
	return &NoteBlock{}
}

// AddText 添加一个文本模块
func (n *NoteBlock) AddText(t *TextBlock) *NoteBlock {
	n.elements = append(n.elements, t)
	return n
}

// AddImage 添加一个图片模块
func (n *NoteBlock) AddImage(i *ImgBlock) *NoteBlock {
	n.elements = append(n.elements, i)
	return n
}
