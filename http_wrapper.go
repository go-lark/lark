package lark

import (
	"io"
	"net/http"
)

// HTTPWrapper is a wrapper. The Default implementation is for net/http.
// And you may implement your own by yourself.
type HTTPWrapper interface {
	Do(method, prefix, url string, header http.Header, body io.Reader) (io.ReadCloser, error)
}
