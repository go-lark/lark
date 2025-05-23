package lark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

// ExpandURL expands url path to full url
func (bot Bot) ExpandURL(urlPath string) string {
	url := fmt.Sprintf("%s%s", bot.domain, urlPath)
	return url
}

func (bot Bot) httpErrorLog(ctx context.Context, prefix, text string, err error) {
	bot.logger.Log(ctx, LogLevelError, fmt.Sprintf("[%s] %s: %+v\n", prefix, text, err))
}

// PerformAPIRequest performs API request
func (bot Bot) PerformAPIRequest(
	ctx context.Context,
	method string,
	prefix, urlPath string,
	header http.Header, auth bool,
	body io.Reader,
	output interface{},
) error {
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
		respBody, err = bot.customClient.Do(ctx, method, url, header, body)
		if err != nil {
			bot.httpErrorLog(ctx, prefix, "call failed", err)
			return err
		}
	} else {
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			bot.httpErrorLog(ctx, prefix, "init request failed", err)
			return err
		}
		req.Header = header
		resp, err := bot.client.Do(req)
		if err != nil {
			bot.httpErrorLog(ctx, prefix, "call failed", err)
			return err
		}
		if bot.debug {
			b, _ := httputil.DumpResponse(resp, true)
			bot.logger.Log(ctx, LogLevelDebug, string(b))
		}
		respBody = resp.Body
	}
	defer respBody.Close()
	err = json.NewDecoder(respBody).Decode(&output)
	if err != nil {
		bot.httpErrorLog(ctx, prefix, "decode body failed", err)
		return err
	}
	return err
}

func (bot Bot) wrapAPIRequest(ctx context.Context, method, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		bot.httpErrorLog(ctx, prefix, "encode JSON failed", err)
		return err
	}

	header := make(http.Header)
	header.Set("Content-Type", "application/json; charset=utf-8")
	err = bot.PerformAPIRequest(ctx, method, prefix, urlPath, header, auth, buf, output)
	if err != nil {
		return err
	}
	return nil
}

// PostAPIRequest call Lark API
func (bot Bot) PostAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodPost, prefix, urlPath, auth, params, output)
}

// GetAPIRequest call Lark API
func (bot Bot) GetAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodGet, prefix, urlPath, auth, params, output)
}

// DeleteAPIRequest call Lark API
func (bot Bot) DeleteAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodDelete, prefix, urlPath, auth, params, output)
}

// PutAPIRequest call Lark API
func (bot Bot) PutAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodPut, prefix, urlPath, auth, params, output)
}

// PatchAPIRequest call Lark API
func (bot Bot) PatchAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodPatch, prefix, urlPath, auth, params, output)
}
