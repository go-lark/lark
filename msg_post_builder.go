package lark

// PostContent .
type PostContent map[string]PostBody

// PostBody .
type PostBody struct {
	Title   string       `json:"title"`
	Content [][]PostElem `json:"content"`
}

// PostElem .
type PostElem struct {
	Tag string `json:"tag"`
	// For Text
	UnEscape *bool   `json:"un_escape,omitempty"`
	Text     *string `json:"text,omitempty"`
	Lines    *int    `json:"lines,omitempty"`
	// For Link
	Href *string `json:"href,omitempty"`
	// For At
	UserID *string `json:"user_id,omitempty"`
	// For Image
	ImageKey    *string `json:"image_key,omitempty"`
	ImageWidth  *int    `json:"width,omitempty"`
	ImageHeight *int    `json:"height,omitempty"`
}

const (
	msgPostText  = "text"
	msgPostLink  = "a"
	msgPostAt    = "at"
	msgPostImage = "img"
)

// MsgPostBuilder for build text buf
type MsgPostBuilder struct {
	buf    []PostElem
	locale string
	title  string
}

const defaultLocale = "zh_cn"

// NewPostBuilder creates a text builder
func NewPostBuilder() *MsgPostBuilder {
	return &MsgPostBuilder{
		buf:    make([]PostElem, 0),
		locale: defaultLocale,
	}
}

// Locale sets locale
func (pb *MsgPostBuilder) Locale(locale string) *MsgPostBuilder {
	pb.locale = locale
	return pb
}

// Title sets title
func (pb *MsgPostBuilder) Title(title string) *MsgPostBuilder {
	pb.title = title
	return pb
}

// TextTag creates a text tag
func (pb *MsgPostBuilder) TextTag(text string, lines int, unescape bool) *MsgPostBuilder {
	pe := PostElem{
		Tag:      msgPostText,
		Text:     &text,
		Lines:    &lines,
		UnEscape: &unescape,
	}
	pb.buf = append(pb.buf, pe)
	return pb
}

// LinkTag creates a link tag
func (pb *MsgPostBuilder) LinkTag(text, href string) *MsgPostBuilder {
	pe := PostElem{
		Tag:  msgPostLink,
		Text: &text,
		Href: &href,
	}
	pb.buf = append(pb.buf, pe)
	return pb
}

// AtTag creates an at tag
func (pb *MsgPostBuilder) AtTag(text, userID string) *MsgPostBuilder {
	pe := PostElem{
		Tag:    msgPostAt,
		Text:   &text,
		UserID: &userID,
	}
	pb.buf = append(pb.buf, pe)
	return pb
}

// ImageTag creates an image tag
func (pb *MsgPostBuilder) ImageTag(imageKey string, imageWidth, imageHeight int) *MsgPostBuilder {
	pe := PostElem{
		Tag:         msgPostImage,
		ImageKey:    &imageKey,
		ImageWidth:  &imageWidth,
		ImageHeight: &imageHeight,
	}
	pb.buf = append(pb.buf, pe)
	return pb
}

// Clear all message
func (pb *MsgPostBuilder) Clear() {
	pb.title = ""
	pb.locale = defaultLocale
	pb.buf = make([]PostElem, 0)
}

// Render message
func (pb *MsgPostBuilder) Render() *PostContent {
	content := make(PostContent)
	content[pb.locale] = PostBody{
		Title:   pb.title,
		Content: [][]PostElem{pb.buf},
	}
	return &content
}

// Len returns buf len
func (pb MsgPostBuilder) Len() int {
	return len(pb.buf)
}
