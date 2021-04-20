package lark

// Import from Lark API Go demo
// with adaption to go-lark frame

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	uploadImageURL = "/open-apis/image/v4/put/"
)

// UploadImageResponse .
type UploadImageResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data"`
}

// UploadImage uploads image to Lark server
func (bot *Bot) UploadImage(path string) (*UploadImageResponse, error) {
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
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s%s", bot.domain, uploadImageURL), body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	var bearer = "Bearer " + bot.tenantAccessToken
	request.Header.Set("Authorization", bearer)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respData UploadImageResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		bot.logger.Printf("UploadImage decode body failed: %+v\n", err)
		return nil, err
	}
	return &respData, err
}

// UploadImageObject uploads image to Lark server
func (bot *Bot) UploadImageObject(img image.Image) (*UploadImageResponse, error) {
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
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s%s", bot.domain, uploadImageURL), body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	var bearer = "Bearer " + bot.tenantAccessToken
	request.Header.Set("Authorization", bearer)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respData UploadImageResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		bot.logger.Printf("UploadImage decode body failed: %+v\n", err)
		return nil, err
	}
	return &respData, err
}
