package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// httpPost send http posts
func (bot Bot) httpPost(urlPath string, params interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		bot.logger.Printf("Encode json failed: %+v\n", err)
		return nil, err
	}
	url := fmt.Sprintf("%s%s", bot.domain, urlPath)
	resp, err := bot.client.Post(url, "application/json; charset=utf-8", buf)
	return resp, err
}

// httpPostWithAuth send http posts with bearer token
func (bot Bot) httpPostWithAuth(urlPath string, token string, params interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		bot.logger.Printf("Encode json failed: %+v\n", err)
		return nil, err
	}
	url := fmt.Sprintf("%s%s", bot.domain, urlPath)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		bot.logger.Printf("Init request failed: %+v\n", err)
		return nil, err
	}

	var bearer = "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := bot.client.Do(req)
	return resp, err
}

// PostAPIRequest call Lark API without auth tokens
func (bot Bot) PostAPIRequest(prefix, urlPath string, params interface{}, output interface{}) error {
	resp, err := bot.httpPost(urlPath, params)
	if err != nil {
		bot.logger.Printf("[%s] call failed: %+v\n", prefix, err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		bot.logger.Printf("[%s] decode body failed: %+v\n", prefix, err)
		return err
	}
	return nil
}

// PostAPIRequestWithAuth call Lark API with auth tokens
func (bot Bot) PostAPIRequestWithAuth(prefix, urlPath string, params interface{}, output interface{}) error {
	resp, err := bot.httpPostWithAuth(urlPath, bot.tenantAccessToken, params)
	if err != nil {
		bot.logger.Printf("[%s] call failed: %+v\n", prefix, err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		bot.logger.Printf("[%s] decode body failed: %+v\n", prefix, err)
		return err
	}
	return nil
}
