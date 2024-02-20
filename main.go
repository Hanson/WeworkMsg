package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/NICEXAI/WeWorkFinanceSDK"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io"
	"log"
	"net/http"
)

type ChatData struct {
	Seq          uint64      `json:"seq,omitempty"`           // 消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	MsgId        string      `json:"msgid,omitempty"`         // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	PublickeyVer uint32      `json:"publickey_ver,omitempty"` // 加密此条消息使用的公钥版本号。
	Message      interface{} `json:"message"`
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	err := loadConfig()
	if err != nil {
		log.Printf("err: %+v", err)
		return
	}

	client, err := WeWorkFinanceSDK.NewClient(Cfg.CorpId, Cfg.CorpSecret, Cfg.RsaPrivateKey)
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}

	http.HandleFunc("/get_chat_data", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		b, err := io.ReadAll(request.Body)
		if err != nil {
			responseError(writer, err)
			return
		}

		seq := gjson.GetBytes(b, "seq").Uint()
		limit := gjson.GetBytes(b, "limit").Uint()
		proxy := gjson.GetBytes(b, "proxy").String()
		passwd := gjson.GetBytes(b, "passwd").String()
		timeout := gjson.GetBytes(b, "timeout").Int()

		//同步消息
		chatDataList, err := client.GetChatData(seq, limit, proxy, passwd, int(timeout))
		if err != nil {
			responseError(writer, err)
			return
		}

		var list []ChatData

		for _, chatData := range chatDataList {
			//消息解密
			chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
			if err != nil {
				responseError(writer, err)
				return
			}

			var cd ChatData

			cd.Seq = chatData.Seq
			cd.MsgId = chatData.MsgId
			cd.PublickeyVer = chatData.PublickeyVer

			switch chatInfo.Type {
			case "text":
				cd.Message = chatInfo.GetTextMessage()
			case "image":
				cd.Message = chatInfo.GetImageMessage()
			case "revoke":
				cd.Message = chatInfo.GetRevokeMessage()
			case "agree":
				cd.Message = chatInfo.GetAgreeMessage()
			case "voice":
				cd.Message = chatInfo.GetVoiceMessage()
			case "video":
				cd.Message = chatInfo.GetVideoMessage()
			case "card":
				cd.Message = chatInfo.GetCardMessage()
			}

			list = append(list, cd)
		}

		responseOk(writer, list)
	})
	http.HandleFunc("/get_media_data", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		b, err := io.ReadAll(request.Body)
		if err != nil {
			responseError(writer, err)
			return
		}

		sdkfileid := gjson.GetBytes(b, "sdk_file_id").String()
		proxy := gjson.GetBytes(b, "proxy").String()
		passwd := gjson.GetBytes(b, "passwd").String()
		timeout := gjson.GetBytes(b, "timeout").Int()

		isFinish := false
		buffer := bytes.Buffer{}
		indexBuf := ""
		for !isFinish {
			//获取媒体数据
			mediaData, err := client.GetMediaData(indexBuf, sdkfileid, proxy, passwd, int(timeout))
			if err != nil {
				responseError(writer, err)
				return
			}
			buffer.Write(mediaData.Data)
			if mediaData.IsFinish {
				isFinish = mediaData.IsFinish
			}
			indexBuf = mediaData.OutIndexBuf
		}

		responseOk(writer, base64.StdEncoding.EncodeToString(buffer.Bytes()))
	})

	http.ListenAndServe(":"+Cfg.Port, nil)
}

func responseError(w http.ResponseWriter, err error) {
	response(w, 1, err.Error())
}

func responseOk(w http.ResponseWriter, data interface{}) {
	response(w, 0, data)
}

func response(w http.ResponseWriter, errCode int, data interface{}) {
	resp, _ := sjson.SetBytes([]byte{}, "err_code", errCode)
	resp, _ = sjson.SetBytes(resp, "data", data)
	_, _ = w.Write(resp)
}
