package main

import (
	"encoding/json"
	"testing"
)

func TestDecodeWeAppMessage(t *testing.T) {
	raw := json.RawMessage(`{
		"msgid":"11930598857592605935_1603875608",
		"action":"send",
		"from":"kens",
		"tolist":["wmGAgeDQAAsgQetTQGqRbMxrkodpM3fA"],
		"roomid":"",
		"msgtime":1603875608691,
		"msgtype":"weapp",
		"weapp":{
			"title":"开始聊天前请仔细阅读服务须知事项",
			"description":"客户需同意存档聊天记录",
			"username":"xxx@app",
			"displayname":"服务须知",
			"appid":"wx062f7a5507909000",
			"pagepath":"pages/home/index"
		}
	}`)

	message, err := decodeWeAppMessage(raw)
	if err != nil {
		t.Fatalf("decode weapp message: %v", err)
	}
	if message.MsgType != "weapp" || message.WeApp.AppID != "wx062f7a5507909000" {
		t.Fatalf("unexpected message: %+v", message)
	}
	if message.WeApp.PagePath != "pages/home/index" {
		t.Fatalf("unexpected pagepath: %q", message.WeApp.PagePath)
	}
}
