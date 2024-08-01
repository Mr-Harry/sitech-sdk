package paper_model

import (
	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/rag_model"
	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/result_model"
)

type SearchPubsByKwdResp struct {
	result_model.Result
	Data SearchPubsByKwdData `json:"data"`
}

type SearchPubsByKwdData struct {
	Pubs []rag_model.PubItem `json:"pubs"`
}

type PubsCountResp struct {
	result_model.Result
	Data PubsCountData `json:"data"`
}

type PubsCountData struct {
	Total int `json:"total"`
}

type TextSummaryResp struct {
	result_model.Result
	Data TextSummaryData `json:"data"`
}

type TextSummaryData struct {
	Answer string `json:"answer"`
}