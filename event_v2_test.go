package lark

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestEventTypes(t *testing.T) {
	event := EventV2{
		Header: EventV2Header{
			EventType: EventTypeChatDisbanded,
		},
		EventRaw: json.RawMessage(`{ "message": { "chat_id": "oc_ae7f3952a9b28588aeac46c9853d25d3", "chat_type": "p2p", "content": "{\"text\":\"333\"}", "create_time": "1641385820771", "message_id": "om_6ff2cff41a3e9248bbb19bf0e4762e6e", "message_type": "text" }, "sender": { "sender_id": { "open_id": "ou_4f75b532aff410181e93552ad0532072", "union_id": "on_2312aab89ab7c87beb9a443b2f3b1342", "user_id": "4gbb63af" }, "sender_type": "user", "tenant_key": "736588c9260f175d" } }`),
	}
	m, e := event.GetMessageReceived()
	assert.Error(t, e)
	event.Header.EventType = EventTypeMessageReceived
	m, e = event.GetMessageReceived()
	assert.NoError(t, e)
	assert.Equal(t, "p2p", m.Message.ChatType)
}

func TestGetEvent(t *testing.T) {
	event := EventV2{
		Header: EventV2Header{
			EventType: EventTypeMessageReceived,
		},
		EventRaw: json.RawMessage(`{ "message": { "chat_id": "oc_ae7f3952a9b28588aeac46c9853d25d3", "chat_type": "p2p", "content": "{\"text\":\"333\"}", "create_time": "1641385820771", "message_id": "om_6ff2cff41a3e9248bbb19bf0e4762e6e", "message_type": "text" }, "sender": { "sender_id": { "open_id": "ou_4f75b532aff410181e93552ad0532072", "union_id": "on_2312aab89ab7c87beb9a443b2f3b1342", "user_id": "4gbb63af" }, "sender_type": "user", "tenant_key": "736588c9260f175d" } }`),
	}
	var ev EventV2MessageReceived
	err := event.GetEvent(EventTypeMessageReceived, &ev)
	if assert.NoError(t, err) {
		assert.Equal(t, "oc_ae7f3952a9b28588aeac46c9853d25d3", ev.Message.ChatID)
	}
}
