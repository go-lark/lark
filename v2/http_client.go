package lark

import (
	"context"
	"net/http"
	"time"
)

// defaultClient .
type defaultClient struct {
	c *http.Client
}

// newDefaultClient .
func newDefaultClient() *defaultClient {
	return &defaultClient{
		c: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Do .
func (dc defaultClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return dc.c.Do(req)
}
