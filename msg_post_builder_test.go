package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostLocale(t *testing.T) {
	pb := NewPostBuilder()
	assert.Equal(t, defaultLocale, pb.locale)
	pb.Locale("en")
	assert.Equal(t, "en", pb.locale)
}

func TestPostTitle(t *testing.T) {
	pb := NewPostBuilder()
	pb.Title("title")
	assert.Equal(t, "title", pb.title)
}

func TestPostTextTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.TextTag("hello, world", 1, true)
	assert.Equal(t, "text", pb.buf[0].Tag)
	assert.Equal(t, "hello, world", *pb.buf[0].Text)
	assert.Equal(t, 1, *pb.buf[0].Lines)
	assert.Equal(t, true, *pb.buf[0].UnEscape)
}
func TestPostLinkTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.LinkTag("hello, world", "https://www.toutiao.com/")
	assert.Equal(t, "a", pb.buf[0].Tag)
	assert.Equal(t, "hello, world", *pb.buf[0].Text)
	assert.Equal(t, "https://www.toutiao.com/", *pb.buf[0].Href)
}
func TestPostAtTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.AtTag("www", "123456")
	assert.Equal(t, "at", pb.buf[0].Tag)
	assert.Equal(t, "www", *pb.buf[0].Text)
	assert.Equal(t, "123456", *pb.buf[0].UserID)
}
func TestPostImgTag(t *testing.T) {
	pb := NewPostBuilder()
	pb.ImageTag("d9f7d37e-c47c-411b-8ec6-9861132e6986", 320, 240)
	assert.Equal(t, "img", pb.buf[0].Tag)
	assert.Equal(t, "d9f7d37e-c47c-411b-8ec6-9861132e6986", *pb.buf[0].ImageKey)
	assert.Equal(t, 240, *pb.buf[0].ImageHeight)
	assert.Equal(t, 320, *pb.buf[0].ImageWidth)
}

func TestPostClearAndLen(t *testing.T) {
	pb := NewPostBuilder()
	pb.TextTag("hello, world", 1, true).LinkTag("link", "https://www.toutiao.com/")
	assert.Equal(t, 2, pb.Len())
	pb.Clear()
	assert.Empty(t, pb.buf)
	assert.Equal(t, 0, pb.Len())
}
