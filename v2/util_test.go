package lark

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileDownload(t *testing.T) {
	ts := time.Now().Unix()
	filename := fmt.Sprintf("/tmp/go-lark-ci-%d", ts)
	err := DownloadFile(filename, "https://s1-fs.pstatp.com/static-resource/v1/363e0009ef09d43d5a96~?image_size=72x72&cut_type=&quality=&format=png&sticker_format=.webp")
	if assert.NoError(t, err) {
		assert.FileExists(t, filename)
	}
}
