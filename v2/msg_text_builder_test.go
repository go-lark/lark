package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseText(t *testing.T) {
	tb := NewTextBuilder()
	msg := tb.Text("hello, ", "world", 123).Render()
	assert.Equal(t, "hello, world123", tb.buf[0].content)
	assert.Equal(t, "hello, world123", msg)
}

func TestTextTextf(t *testing.T) {
	tb := NewTextBuilder()
	msg := tb.Textf("hello, %s: %d", "world", 1).Render()
	assert.Equal(t, "hello, world: 1", tb.buf[0].content)
	assert.Equal(t, "hello, world: 1", msg)
}

func TestTextTextln(t *testing.T) {
	tb := NewTextBuilder()
	msg := tb.Textln("hello", "world").Render()
	assert.Equal(t, "hello world\n", tb.buf[0].content)
	assert.Equal(t, "hello world\n", msg)
}

func TestTextMention(t *testing.T) {
	tb := NewTextBuilder()
	msg := tb.Text("hello, world").Mention("6454030812462448910").Render()
	assert.Equal(t, "hello, world<at user_id=\"6454030812462448910\">@user</at>", msg)
}

func TestTextMentionAll(t *testing.T) {
	tb := NewTextBuilder()
	msg := tb.Text("hello, world").MentionAll().Render()
	assert.Equal(t, "hello, world<at user_id=\"all\">@all</at>", msg)
}

func TestTextClearAndLen(t *testing.T) {
	tb := NewTextBuilder()
	tb.Text("hello, world").MentionAll()
	assert.Equal(t, 2, tb.Len())
	tb.Clear()
	assert.Empty(t, tb.buf)
	assert.Equal(t, 0, tb.Len())
}
