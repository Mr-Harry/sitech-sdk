package oauth_model

import "gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/result_model"

type AccessTokenResp struct {
	result_model.Result
	Data TokenData `json:"data"`
}

type TokenData struct {
	Token string `json:"token"`
}