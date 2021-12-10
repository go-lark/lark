package lark

import (
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadImage(t *testing.T) {
	resp, err := bot.UploadImage("./fixtures/test.jpg")
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.ImageKey)
		assert.NotEmpty(t, resp.Data.ImageKey)
	}
}

func TestUploadImageObject(t *testing.T) {
	file, _ := os.Open("./fixtures/test.jpg")
	img, _ := jpeg.Decode(file)

	resp, err := bot.UploadImageObject(img)
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.ImageKey)
		assert.NotEmpty(t, resp.Data.ImageKey)
	}
}
