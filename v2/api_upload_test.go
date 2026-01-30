package lark

import (
	"image/jpeg"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadImage(t *testing.T) {
	resp, err := bot.UploadImage(t.Context(), "../fixtures/test.jpg")
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.ImageKey)
		assert.NotEmpty(t, resp.Data.ImageKey)
	}
}

func TestUploadImageObject(t *testing.T) {
	file, _ := os.Open("../fixtures/test.jpg")
	img, _ := jpeg.Decode(file)

	resp, err := bot.UploadImageObject(t.Context(), img)
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.ImageKey)
		assert.NotEmpty(t, resp.Data.ImageKey)
	}
}

func TestUploadFile(t *testing.T) {
	resp, err := bot.UploadFile(
		t.Context(),
		UploadFileRequest{
			FileType: "pdf",
			FileName: "hello.pdf",
			Path:     "../fixtures/test.pdf",
		})
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.FileKey)
		assert.NotEmpty(t, resp.Data.FileKey)
	}
}

func TestUploadFile_Binary(t *testing.T) {
	resp, err := bot.UploadFile(
		t.Context(),
		UploadFileRequest{
			FileType: "stream",
			FileName: "test-data.csv",
			Reader: strings.NewReader(`Name,Age,Location
Foo,25,Sleman
Bar,23,Sidoarjo
Baz,27,Bantul`),
		})
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.FileKey)
		assert.NotEmpty(t, resp.Data.FileKey)
	}
}
