package lark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

// HTTPClient is an interface handling http requests
type HTTPClient interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

// ExpandURL expands url path to full url
func (bot Bot) ExpandURL(urlPath string) string {
	url := fmt.Sprintf("%s%s", bot.domain, urlPath)
	return url
}

func (bot Bot) httpErrorLog(ctx context.Context, prefix, text string, err error) {
	bot.logger.Log(ctx, LogLevelError, fmt.Sprintf("[%s] %s: %+v\n", prefix, text, err))
}

func (bot *Bot) loadAndRenewToken(ctx context.Context) (string, error) {
	now := time.Now()
	// check token
	token, ok := bot.tenantAccessToken.Load().(TenantAccessToken)
	tenantAccessToken := token.TenantAccessToken
	if !ok || token.TenantAccessToken == "" || (token.EstimatedExpireAt != nil && now.After(*token.EstimatedExpireAt)) {
		// renew token
		if bot.autoRenew {
			tacResp, err := bot.GetTenantAccessTokenInternal(ctx)
			if err != nil {
				return "", err
			}
			now := time.Now()
			expire := time.Duration(tacResp.Expire - 10)
			eta := now.Add(expire)
			token := TenantAccessToken{
				TenantAccessToken: tacResp.TenantAccessToken,
				Expire:            expire,
				LastUpdatedAt:     &now,
				EstimatedExpireAt: &eta,
			}
			bot.tenantAccessToken.Store(token)
			tenantAccessToken = tacResp.TenantAccessToken
		}
	}
	return tenantAccessToken, nil
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
		tenantAccessToken, err := bot.loadAndRenewToken(ctx)
		if err != nil {
			return err
		}
		header.Add("Authorization", fmt.Sprintf("Bearer %s", tenantAccessToken))
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		bot.httpErrorLog(ctx, prefix, "init request failed", err)
		return err
	}
	req.Header = header
	resp, err := bot.client.Do(ctx, req)
	if err != nil {
		bot.httpErrorLog(ctx, prefix, "call failed", err)
		return err
	}
	if bot.debug {
		b, _ := httputil.DumpResponse(resp, true)
		bot.logger.Log(ctx, LogLevelDebug, string(b))
	}
	respBody = resp.Body
	defer respBody.Close()
	buffer, err := io.ReadAll(respBody)
	if err != nil {
		bot.httpErrorLog(ctx, prefix, "read body failed", err)
		return err
	}
	// read response content
	err = json.Unmarshal(buffer, &output)
	if err != nil {
		bot.httpErrorLog(ctx, prefix, "decode body failed", err)
		return err
	}
	// read error code
	var dummyOutput DummyResponse
	err = json.Unmarshal(buffer, &dummyOutput)
	if err == nil && dummyOutput.Code != 0 {
		apiError := APIError(url, dummyOutput.BaseResponse)
		bot.logger.Log(ctx, LogLevelError, apiError.Error())
		return apiError
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
	return err
}

// PostAPIRequest POSTs Lark API
func (bot Bot) PostAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodPost, prefix, urlPath, auth, params, output)
}

// GetAPIRequest GETs Lark API
func (bot Bot) GetAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodGet, prefix, urlPath, auth, params, output)
}

// DeleteAPIRequest DELETEs Lark API
func (bot Bot) DeleteAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodDelete, prefix, urlPath, auth, params, output)
}

// PutAPIRequest PUTs Lark API
func (bot Bot) PutAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodPut, prefix, urlPath, auth, params, output)
}

// PatchAPIRequest PATCHes Lark API
func (bot Bot) PatchAPIRequest(ctx context.Context, prefix, urlPath string, auth bool, params interface{}, output interface{}) error {
	return bot.wrapAPIRequest(ctx, http.MethodPatch, prefix, urlPath, auth, params, output)
}
