package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESDecrypt(t *testing.T) {
	key := EncryptKey("test key")
	data, _ := Decrypt(key, "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk=")
	assert.Equal(t, "hello world", string(data))
	assert.Equal(t, 11, len(data))
}
