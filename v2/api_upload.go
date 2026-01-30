package lark

// Import from Lark API Go demo
// with adaption to go-lark frame

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	uploadImageURL = "/open-apis/im/v1/images"
	uploadFileURL  = "/open-apis/im/v1/files"
)

// UploadImageResponse .
type UploadImageResponse struct {
	BaseResponse
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data"`
}

// UploadFileRequest .
type UploadFileRequest struct {
	FileType string    `json:"-"`
	FileName string    `json:"-"`
	Duration int       `json:"-"`
	Path     string    `json:"-"`
	Reader   io.Reader `json:"-"`
}

// UploadFileResponse .
type UploadFileResponse struct {
	BaseResponse
	Data struct {
		FileKey string `json:"file_key"`
	} `json:"data"`
}

// UploadImage uploads image file
func (bot *Bot) UploadImage(ctx context.Context, path string) (*UploadImageResponse, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("image_type", "message")
	part, err := writer.CreateFormFile("image", path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	var respData UploadImageResponse
	header := make(http.Header)
	header.Set("Content-Type", writer.FormDataContentType())
	err = bot.PerformAPIRequest(ctx, "POST", "UploadImage", uploadImageURL, header, true, body, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, err
}

// UploadImageObject uploads image object
func (bot *Bot) UploadImageObject(ctx context.Context, img image.Image) (*UploadImageResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("image_type", "message")
	part, err := writer.CreateFormFile("image", "temp_image")
	if err != nil {
		return nil, err
	}
	err = jpeg.Encode(part, img, nil)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	var respData UploadImageResponse
	header := make(http.Header)
	header.Set("Content-Type", writer.FormDataContentType())
	err = bot.PerformAPIRequest(ctx, "POST", "UploadImage", uploadImageURL, header, true, body, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, err
}

// UploadFile uploads file to Lark server
func (bot *Bot) UploadFile(ctx context.Context, req UploadFileRequest) (*UploadFileResponse, error) {
	var content io.Reader
	if req.Reader == nil {
		file, err := os.Open(req.Path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		content = file
	} else {
		content = req.Reader
		req.Path = req.FileName
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("file_type", req.FileType)
	writer.WriteField("file_name", req.FileName)
	if req.FileType == "mp4" && req.Duration > 0 {
		writer.WriteField("duration", fmt.Sprintf("%d", req.Duration))
	}
	part, err := writer.CreateFormFile("file", req.Path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, content)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	var respData UploadFileResponse
	header := make(http.Header)
	header.Set("Content-Type", writer.FormDataContentType())
	err = bot.PerformAPIRequest(ctx, "POST", "UploadFile", uploadFileURL, header, true, body, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, err
}
