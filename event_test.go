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

func TestPostEventMessage(t *testing.T) {
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

func httpHandlerV2(w http.ResponseWriter, r *http.Request) {
	var m EventV2
	json.NewDecoder(r.Body).Decode(&m)
	w.Write([]byte(m.Schema))
}

func TestPostEventV2(t *testing.T) {
	message := EventV2{
		Schema: "2.0",
		Header: EventV2Header{
			AppID: "666",
		},
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
	w := performRequest(httpHandlerV2, "POST", "/", message)
	assert.Equal(t, "2.0", string(w.Body.Bytes()))
}
