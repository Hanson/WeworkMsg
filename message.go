package main

import (
	"encoding/json"

	"github.com/NICEXAI/WeWorkFinanceSDK"
)

// WeAppMessage represents an enterprise WeChat mini-program message.
//
// The SDK's WeAppMessage type predates appid and pagepath in the enterprise
// WeChat payload, so keep the complete payload here instead of dropping those
// fields during JSON unmarshalling.
type WeAppMessage struct {
	WeWorkFinanceSDK.BaseMessage
	WeApp struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Username    string `json:"username,omitempty"`
		DisplayName string `json:"displayname,omitempty"`
		AppID       string `json:"appid,omitempty"`
		PagePath    string `json:"pagepath,omitempty"`
	} `json:"weapp,omitempty"`
}

func decodeWeAppMessage(raw json.RawMessage) (WeAppMessage, error) {
	var message WeAppMessage
	if err := json.Unmarshal(raw, &message); err != nil {
		return WeAppMessage{}, err
	}
	return message, nil
}

// decodeChatMessage converts every message type currently exposed by the
// Finance SDK to its typed representation. Unknown or newer message types are
// returned as raw JSON so the API remains forward-compatible instead of
// returning a null message.
func decodeChatMessage(chatInfo WeWorkFinanceSDK.ChatMessage) (interface{}, error) {
	switch chatInfo.Type {
	case "text":
		return chatInfo.GetTextMessage(), nil
	case "image":
		return chatInfo.GetImageMessage(), nil
	case "revoke":
		return chatInfo.GetRevokeMessage(), nil
	case "agree", "disagree":
		return chatInfo.GetAgreeMessage(), nil
	case "voice":
		return chatInfo.GetVoiceMessage(), nil
	case "video":
		return chatInfo.GetVideoMessage(), nil
	case "card":
		return chatInfo.GetCardMessage(), nil
	case "location":
		return chatInfo.GetLocationMessage(), nil
	case "emotion":
		return chatInfo.GetEmotionMessage(), nil
	case "file":
		return chatInfo.GetFileMessage(), nil
	case "link":
		return chatInfo.GetLinkMessage(), nil
	case "weapp":
		return decodeWeAppMessage(chatInfo.GetRawChatMessage())
	case "chatrecord":
		return chatInfo.GetChatRecordMessage(), nil
	case "todo":
		return chatInfo.GetTodoMessage(), nil
	case "vote":
		return chatInfo.GetVoteMessage(), nil
	case "collect":
		return chatInfo.GetCollectMessage(), nil
	case "redpacket":
		return chatInfo.GetRedpacketMessage(), nil
	case "meeting":
		return chatInfo.GetMeetingMessage(), nil
	case "docmsg":
		return chatInfo.GetDocMessage(), nil
	case "markdown":
		return chatInfo.GetMarkdownMessage(), nil
	case "news":
		return chatInfo.GetNewsMessage(), nil
	case "calendar":
		return chatInfo.GetCalendarMessage(), nil
	case "mixed":
		return chatInfo.GetMixedMessage(), nil
	case "meeting_voice_call":
		return chatInfo.GetMeetingVoiceCallMessage(), nil
	case "voip_doc_share":
		return chatInfo.GetVoipDocShareMessage(), nil
	case "external_redpacket":
		return chatInfo.GetExternalRedPacketMessage(), nil
	case "sphfeed":
		return chatInfo.GetSphFeedMessage(), nil
	case "switch":
		return chatInfo.GetSwitchMessage(), nil
	case "voiptext":
		return chatInfo.GetVoiptextMessage(), nil
	default:
		return chatInfo.GetRawChatMessage(), nil
	}
}
