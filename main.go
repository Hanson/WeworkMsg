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

		var chatInfoList []WeWorkFinanceSDK.ChatMessage

		for _, chatData := range chatDataList {
			//消息解密
			chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
			if err != nil {
				responseError(writer, err)
				return
			}

			chatInfoList = append(chatInfoList, chatInfo)
		}

		responseOk(writer, chatInfoList)
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
