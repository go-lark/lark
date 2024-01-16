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

// PostBuf .
type PostBuf struct {
	Title   string       `json:"title"`
	Content [][]PostElem `json:"content"`
}

// MsgPostBuilder for build text buf
type MsgPostBuilder struct {
	buf       map[string]*PostBuf
	curLocale string
}

const defaultLocale = LocaleZhCN

// NewPostBuilder creates a text builder
func NewPostBuilder() *MsgPostBuilder {
	return &MsgPostBuilder{
		buf:       make(map[string]*PostBuf),
		curLocale: defaultLocale,
	}
}

// Locale renamed to WithLocale but still available
func (pb *MsgPostBuilder) Locale(locale string) *MsgPostBuilder {
	return pb.WithLocale(locale)
}

// WithLocale switches to locale and returns self
func (pb *MsgPostBuilder) WithLocale(locale string) *MsgPostBuilder {
	if _, ok := pb.buf[locale]; !ok {
		buf := &PostBuf{Content: make([][]PostElem, 1)}
		pb.buf[locale] = buf
	}

	pb.curLocale = locale
	return pb
}

// CurLocale switches to locale and returns the buffer of that locale
func (pb *MsgPostBuilder) CurLocale() *PostBuf {
	return pb.WithLocale(pb.curLocale).buf[pb.curLocale]
}

// Title sets title
func (pb *MsgPostBuilder) Title(title string) *MsgPostBuilder {
	pb.CurLocale().Title = title
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
	pb.addElem(pe)
	return pb
}

// LinkTag creates a link tag
func (pb *MsgPostBuilder) LinkTag(text, href string) *MsgPostBuilder {
	pe := PostElem{
		Tag:  msgPostLink,
		Text: &text,
		Href: &href,
	}
	pb.addElem(pe)
	return pb
}

// AtTag creates an at tag
func (pb *MsgPostBuilder) AtTag(text, userID string) *MsgPostBuilder {
	pe := PostElem{
		Tag:    msgPostAt,
		Text:   &text,
		UserID: &userID,
	}
	pb.addElem(pe)
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
	pb.addElem(pe)
	return pb
}

// NewLine starts a new line
func (pb *MsgPostBuilder) NewLine() *MsgPostBuilder {
	// if pb.Len() == 0 then there's no need to start a new line
	if pb.Len() > 0 {
		buf := pb.CurLocale()
		buf.Content = append(buf.Content, make([]PostElem, 0))
	}
	return pb
}

// addElem adds a PostElem to the latest line
func (pb *MsgPostBuilder) addElem(pe PostElem) {
	buf := pb.CurLocale()
	idx := len(buf.Content) - 1
	buf.Content[idx] = append(buf.Content[idx], pe)
}

// Clear all message
func (pb *MsgPostBuilder) Clear() {
	pb.curLocale = defaultLocale
	pb.buf = make(map[string]*PostBuf)
}

// Render message
func (pb *MsgPostBuilder) Render() *PostContent {
	content := make(PostContent)
	for locale, buf := range pb.buf {
		content[locale] = PostBody{
			Title:   buf.Title,
			Content: buf.Content,
		}
	}
	return &content
}

// Len returns the latest line buf len
func (pb MsgPostBuilder) Len() int {
	buf := pb.CurLocale()
	idx := len(buf.Content) - 1
	return len(buf.Content[idx])
}

// Len returns the lines count
func (pb MsgPostBuilder) Lines() int {
	return len(pb.CurLocale().Content)
}
