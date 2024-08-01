package papers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/paper_model"
	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/rag_model"
	jsoniter "github.com/json-iterator/go"
)

func SearchPubsByKwd(kwds []string, withKnowledgeGraph bool, accessToken string) (rs []rag_model.PubItem, err error) {
	kwdsStr, err := jsoniter.MarshalToString(kwds)
	formData := url.Values{}
	formData.Set("kwds", kwdsStr)
	formData.Set("with_knowledge_graph", strconv.FormatBool(withKnowledgeGraph))
	uri := APIHost + "/openapi/v1/paper/search_pubs_by_kwds"
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
	var result paper_model.SearchPubsByKwdResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("SearchPubsByKwd err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	rs = result.Data.Pubs
	return
}

func GetPubsCount(accessToken string) (total int, err error) {
	uri := APIHost + "/openapi/v1/paper/get_pubs_count"
	client := http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
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
	var result paper_model.PubsCountResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("SearchPubsByKwd err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	total = result.Data.Total
	return
}

func TextSummary(abstracts []string, accessToken string) (answer string, err error) {
	kwdsStr, err := jsoniter.MarshalToString(abstracts)
	formData := url.Values{}
	formData.Set("abstracts", kwdsStr)
	uri := APIHost + "/openapi/v1/paper/text_summary"
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
