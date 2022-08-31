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

func TestGenSign(t *testing.T) {
	sign, err := GenSign("xxx", 1661860880)
	if assert.NoError(t, err) {
		assert.Equal(t, "QnWVTSBe6FmQDE0bG6X0mURbI+DnvVyu1h+j5dHOjrU=", sign)
	}
}
