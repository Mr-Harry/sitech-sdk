package chat_model

import "gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/result_model"

type SendMessageResp struct {
	result_model.Result
	Data SendMessageData `json:"data"`
}

type SendMessageData struct {
	Answer string `json:"answer"`
}