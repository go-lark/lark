package lark

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	var m EventMessage
	json.NewDecoder(r.Body).Decode(&m)
	w.Write([]byte(m.Event.Text))
}

func TestPostEventPrivateMessage(t *testing.T) {
	message := EventMessage{
		Timestamp: "",
		Token:     "",
		EventType: "event_callback",
		Event: EventBody{
			Type:          "message",
			ChatType:      "private",
			MsgType:       "text",
			OpenID:        testUserOpenID,
			Text:          "private event",
			Title:         "",
			OpenMessageID: "",
			ImageKey:      "",
			ImageURL:      "",
		},
	}
	w := performRequest(httpHandler, "POST", "/", message)
	assert.Equal(t, "private event", string(w.Body.Bytes()))
}

func TestPostEventAtMessage(t *testing.T) {
	message := EventMessage{
		Timestamp: "",
		Token:     "",
		EventType: "event_callback",
		Event: EventBody{
			Type:          "message",
			ChatType:      "group",
			MsgType:       "text",
			OpenID:        testUserOpenID,
			Text:          "public event",
			Title:         "",
			OpenMessageID: "",
			ImageKey:      "",
			ImageURL:      "",
		},
	}
	w := performRequest(httpHandler, "POST", "/", message)
	assert.Equal(t, "public event", string(w.Body.Bytes()))
}
