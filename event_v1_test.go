package lark

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	w := performRequest(func(w http.ResponseWriter, r *http.Request) {
		var m EventMessage
		json.NewDecoder(r.Body).Decode(&m)
		w.Write([]byte(m.Event.Text))
	}, "POST", "/", message)
	assert.Equal(t, "private event", string(w.Body.Bytes()))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m, _ := json.Marshal(message)
		w.Write([]byte(m))
	}))
	defer ts.Close()
	resp, err := PostEvent(http.DefaultClient, ts.URL, message)
	if assert.NoError(t, err) {
		var event EventMessage
		body, err := ioutil.ReadAll(resp.Body)
		if assert.NoError(t, err) {
			defer resp.Body.Close()
			_ = json.Unmarshal(body, &event)
			assert.Equal(t, "event_callback", event.EventType)
			assert.Equal(t, "message", event.Event.Type)
		}
	}
}
