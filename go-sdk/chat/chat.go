package chat

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/paper_model"
	jsoniter "github.com/json-iterator/go"
)

const APIHost = "https://sdk.visualx.art"

func SendMessage(prompt string, accessToken string) (answer string, err error) {
	formData := url.Values{}
	formData.Set("prompt", prompt)
	uri := APIHost + "/openapi/v1/chat/send_message"
	client := http.Client{}
	req, err := http.NewRequest("POST", uri, strings.NewReader(formData.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var result paper_model.TextSummaryResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	answer = result.Data.Answer
	return
}
