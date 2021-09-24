package lark

import (
	"io"
	"net/http"
)

// HTTPWrapper is a wrapper interface, which enables extension on HTTP part.
// Typicall, we do not need this because default client is sufficient.
type HTTPWrapper interface {
	Do(method, prefix, url string, header http.Header, body io.Reader) (io.ReadCloser, error)
}
