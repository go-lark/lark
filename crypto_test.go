package lark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAESDecrypt(t *testing.T) {
	key := EncryptKey("test key")
	data, _ := Decrypt(key, "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk=")
	assert.Equal(t, "hello world", string(data))
	assert.Equal(t, 11, len(data))
}

func TestGenSign(t *testing.T) {
	sign, err := GenSign("xxx", time.Now().Unix())
	if assert.NoError(t, err) {
		assert.Equal(t, "pmFOXqGI9z7QepNkBUWbdAZs7TQQ9yx1uqAPK44BPG4=", sign)
	}
}
