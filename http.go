package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ExpandURL expands url path to full url
func (bot Bot) ExpandURL(urlPath string) string {
	url := fmt.Sprintf("%s%s", bot.domain, urlPath)
	return url
}

func (bot Bot) httpErrorLog(prefix, text string, err error) {
	bot.logger.Log(bot.ctx, LogLevelError, fmt.Sprintf("[%s] %s: %+v\n", prefix, text, err))
}

// DoAPIRequest builds http request
func (bot Bot) DoAPIRequest(
	method string,
	prefix, urlPath string,
	header http.Header, auth bool,
	body io.Reader,
	output interface{}) error {
	var (
		err      error
		respBody io.ReadCloser
		url      = bot.ExpandURL(urlPath)
	)
	if header == nil {
		header = make(http.Header)
	}
	if auth {
		header.Add("Authorization", fmt.Sprintf("Bearer %s", bot.TenantAccessToken()))
	}
	if bot.useCustomClient {
		if bot.customClient == nil {
			return ErrCustomHTTPClientNotSet
		}
		respBody, err = bot.customClient.Do(bot.ctx, method, url, header, body)
		if err != nil {
			bot.httpErrorLog(prefix, "call failed", err)
			return err
		}
	} else {
		req, err := http.NewRequestWithContext(bot.ctx, method, url, body)
		if err != nil {
			bot.httpErrorLog(prefix, "init request failed", err)
			return err
		}
		req.Header = header
		resp, err := bot.client.Do(req)
		if err != nil {
			bot.httpErrorLog(prefix, "call failed", err)
			return err
		}
		respBody = resp.Body
	}
	defer respBody.Close()
	err = json.NewDecoder(respBody).Decode(&output)
	if err != nil {
		bot.httpErrorLog(prefix, "decode body failed", err)
		return err
	}
	return err
}

// PostAPIRequest call Lark API
func (bot Bot) PostAPIRequest(prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		bot.httpErrorLog(prefix, "encode JSON failed", err)
		return err
	}

	header := make(http.Header)
	header.Set("Content-Type", "application/json; charset=utf-8")
	err = bot.DoAPIRequest("POST", prefix, urlPath, header, auth, buf, output)
	if err != nil {
		return err
	}
	return nil
}
