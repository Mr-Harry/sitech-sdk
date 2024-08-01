package oauth

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/oauth_model"
	jsoniter "github.com/json-iterator/go"
)

const OAuthHost = "https://sdk.visualx.art"

func GetAccessToken(appKey, appSecret string) (token string, err error) {
	url := OAuthHost + fmt.Sprintf("/openapi/v1/oauth/token?client_id=%s&client_secret=%s", appKey, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var tokenResp oauth_model.AccessTokenResp 
	if err = jsoniter.Unmarshal(body, &tokenResp); err != nil {
		return
	}
	if len(tokenResp.Data.Token) <= 0 {
		err = errors.New("system error, access_token get failed")
		return
	}
	token = tokenResp.Data.Token
	return
}
