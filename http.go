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

// RawAPIRequest builds http request
func (bot Bot) RawAPIRequest(method, prefix, urlPath string, auth bool, body *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, bot.ExpandURL(urlPath), body)
	if err != nil {
		bot.logger.Printf("[%s] init request failed: %+v\n", prefix, err)
		return nil, err
	}

	if auth {
		bearer := "Bearer " + bot.tenantAccessToken
		req.Header.Set("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, err
}

// PostAPIRequest call Lark API without auth tokens
func (bot Bot) PostAPIRequest(prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		bot.logger.Printf("[%s] encode json failed: %+v\n", prefix, err)
		return err
	}

	req, err := bot.RawAPIRequest("POST", prefix, urlPath, auth, buf)
	if err != nil {
		return err
	}
	resp, err := bot.client.Do(req)
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
