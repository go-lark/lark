package lark

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	w := performRequest(func(w http.ResponseWriter, r *http.Request) {
		var m EventV2
		json.NewDecoder(r.Body).Decode(&m)
		w.Write([]byte(m.Schema))
	}, "POST", "/", message)
	assert.Equal(t, "2.0", string(w.Body.Bytes()))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m, _ := json.Marshal(message)
		w.Write([]byte(m))
	}))
	defer ts.Close()
	resp, err := message.PostEvent(http.DefaultClient, ts.URL)
	if assert.NoError(t, err) {
		var event EventV2
		body, err := ioutil.ReadAll(resp.Body)
		if assert.NoError(t, err) {
			defer resp.Body.Close()
			_ = json.Unmarshal(body, &event)
			assert.Equal(t, "2.0", event.Schema)
			assert.Equal(t, "666", event.Header.AppID)
		}
	}
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

	event = EventV2{
		Header: EventV2Header{
			EventType: EventTypeCardV2Callback,
		},
		EventRaw: json.RawMessage(`{"operator":{"tenant_key":"73658260f175d","user_id":"2a91f9","union_id":"on_fcc3539c08b49244f3f4d4106d2e","open_id":"ou_6bfe63b8001cb8a2cae0555b79ecd"},"token":"c-7f23aeca4f5cd726c2b70bfa616da2","action":{"value":"confirm","tag":"button","timezone":"","name":"","form_value":"","input_value":"","option":"","options":null,"checked":false},"host":"im_message","delivery_type":"","context":{"url":"","preview_token":"","open_message_id":"om_227c08473975fa87dfcb777e39322","open_chat_id":"oc_5d2f53edc7eadc731f4a984b171"}}`),
	}
	callbackReq, err := event.GetCardV2Callback()
	assert.NoError(t, err)
	val, ok := callbackReq.Action.Value.(string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "confirm", val)

	event = EventV2{
		Header: EventV2Header{
			EventType: EventTypeCardV2Callback,
		},
		EventRaw: json.RawMessage(`{"operator":{"tenant_key":"73658260f175d","user_id":"2a91f9","union_id":"on_fcc3539c08b49244f3f4d4106d2e","open_id":"ou_6bfe63b8001cb8a2cae0555b79ecd"},"token":"c-7f23aeca4f5cd726c2b70bfa616da2","action":{"value":{"confirm":"true"},"tag":"button","timezone":"","name":"","form_value":"","input_value":"","option":"","options":null,"checked":false},"host":"im_message","delivery_type":"","context":{"url":"","preview_token":"","open_message_id":"om_227c08473975fa87dfcb777e39322","open_chat_id":"oc_5d2f53edc7eadc731f4a984b171"}}`),
	}
	callbackReq, err = event.GetCardV2Callback()
	assert.NoError(t, err)

	valObj, ok := callbackReq.Action.Value.(map[string]interface{})
	assert.Equal(t, true, ok)
	assert.Equal(t, "true", valObj["confirm"])

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
