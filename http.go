package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ExpandURL expands url path to full url
func (bot Bot) ExpandURL(urlPath string) string {
	url := fmt.Sprintf("%s%s", bot.domain, urlPath)
	return url
}

// httpPostWithAuth send http posts with bearer token
func (bot Bot) httpPost(urlPath string, auth bool, params interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		bot.logger.Printf("Encode json failed: %+v\n", err)
		return nil, err
	}
	url := bot.ExpandURL(urlPath)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		bot.logger.Printf("Init request failed: %+v\n", err)
		return nil, err
	}

	if auth {
		bearer := "Bearer " + bot.tenantAccessToken
		req.Header.Set("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	resp, err := bot.client.Do(req)
	return resp, err
}

// PostAPIRequest call Lark API without auth tokens
func (bot Bot) PostAPIRequest(prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	resp, err := bot.httpPost(urlPath, auth, params)
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
