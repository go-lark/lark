package card

type Element interface {
	Render() Renderer
}

type Renderer interface{}

type ElementTag struct {
	Tag string `json:"tag"`
}
