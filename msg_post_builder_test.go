package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostLocale(t *testing.T) {
	pb := NewPostBuilder()
	assert.Equal(t, defaultLocale, pb.curLocale)
	pb.Locale("en_us")
	assert.Equal(t, "en_us", pb.curLocale)
	pb.WithLocale("ja_jp")
	assert.Equal(t, "ja_jp", pb.curLocale)
}

func TestPostTitle(t *testing.T) {
	pb := NewPostBuilder()
	pb.Title("title")
	assert.Equal(t, "title", pb.CurLocale().Title)
}

func TestPostTextTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.TextTag("hello, world", 1, true)
	buf := pb.CurLocale().Content
	assert.Equal(t, "text", buf[0].Tag)
	assert.Equal(t, "hello, world", *(buf[0].Text))
	assert.Equal(t, 1, *(buf[0].Lines))
	assert.Equal(t, true, *(buf[0].UnEscape))
}

func TestPostLinkTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.LinkTag("hello, world", "https://www.toutiao.com/")
	buf := pb.CurLocale().Content
	assert.Equal(t, "a", buf[0].Tag)
	assert.Equal(t, "hello, world", *(buf[0].Text))
	assert.Equal(t, "https://www.toutiao.com/", *(buf[0].Href))
}

func TestPostAtTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.AtTag("www", "123456")
	buf := pb.CurLocale().Content
	assert.Equal(t, "at", buf[0].Tag)
	assert.Equal(t, "www", *(buf[0].Text))
	assert.Equal(t, "123456", *(buf[0].UserID))
}

func TestPostImgTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.ImageTag("d9f7d37e-c47c-411b-8ec6-9861132e6986", 320, 240)
	buf := pb.CurLocale().Content
	assert.Equal(t, "img", buf[0].Tag)
	assert.Equal(t, "d9f7d37e-c47c-411b-8ec6-9861132e6986", *(buf[0].ImageKey))
	assert.Equal(t, 240, *(buf[0].ImageHeight))
	assert.Equal(t, 320, *(buf[0].ImageWidth))
}

func TestPostClearAndLen(t *testing.T) {
	pb := NewPostBuilder()
	pb.TextTag("hello, world", 1, true).LinkTag("link", "https://www.toutiao.com/")
	assert.Equal(t, 2, pb.Len())
	pb.Clear()
	assert.Empty(t, pb.buf)
	assert.Equal(t, 0, pb.Len())
}

func TestPostMultiLocaleContent(t *testing.T) {
	pb := NewPostBuilder()
	pb.Title("中文标题")
	assert.Equal(t, "中文标题", pb.CurLocale().Title)
	pb.TextTag("你好世界", 1, true).TextTag("其他内容", 1, true)
	assert.Equal(t, 2, pb.Len())

	pb.WithLocale("en_us").Title("en title")
	pb.TextTag("hello, world", 1, true).LinkTag("link", "https://www.toutiao.com/")
	assert.Equal(t, 2, pb.Len())
	assert.Equal(t, "en title", pb.CurLocale().Title)

	content := pb.Render()
	t.Log(content)
	assert.Equal(t, "中文标题", (*content)["zh_cn"].Title)
	assert.Equal(t, "en title", (*content)["en_us"].Title)
}
