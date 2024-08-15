package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindTemplate(t *testing.T) {
	b := NewTemplateBuilder()
	assert.Empty(t, b.id)
	assert.Empty(t, b.versionName)
	assert.Nil(t, b.data)

	_ = b.BindTemplate("AAqCYI07MQWh1", "1.0.0", map[string]interface{}{
		"name": "志田千陽",
	})
	assert.Equal(t, "AAqCYI07MQWh1", b.id)
	assert.Equal(t, "1.0.0", b.versionName)
	assert.NotEmpty(t, b.data)
}
