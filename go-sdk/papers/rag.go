package papers

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/rag_model"
	jsoniter "github.com/json-iterator/go"
)

const APIHost = "https://sdk.visualx.art"

// 论文RAG检索
func RAGSearchPaper(query, accessToken string) (rs rag_model.RAGSearchPaperData, err error) {
	formData := url.Values{}
	formData.Set("query", query)
	uri := APIHost + "/openapi/v1/rag/rag_search_paper"
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
	var result rag_model.RAGSearchPaperResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("RAGSearchPaper err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	rs = result.Data
	return
}

func RAGSearchPaperByDate(query, startTime, endTime, accessToken string) (rs rag_model.RAGSearchPaperData, err error) {
	formData := url.Values{}
	formData.Set("query", query)
	formData.Set("start_time", startTime)
	formData.Set("end_time", endTime)
	uri := APIHost + "/openapi/v1/rag/rag_search_paper_by_date"
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
	var result rag_model.RAGSearchPaperResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("RAGSearchPaperByDate err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	rs = result.Data
	return
}

func RAGSearchPaperByDoiOrUrl(doi, paperUrl, accessToken string) (rs rag_model.RAGSearchPaperData, err error) {
	formData := url.Values{}
	formData.Set("doi", doi)
	formData.Set("url", paperUrl)
	uri := APIHost + "/openapi/v1/rag/rag_search_paper_by_doi_or_url"
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
	var result rag_model.RAGSearchPaperResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("RRAGSearchPaperByDoiOrUrl err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	rs = result.Data
	return
}

func RAGSearchPaperByAbstract(abstract, accessToken string) (rs rag_model.RAGSearchPaperData, err error) {
	formData := url.Values{}
	formData.Set("abstract", abstract)
	uri := APIHost + "/openapi/v1/rag/rag_search_paper_by_abstract"
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
	var result rag_model.RAGSearchPaperResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("RAGSearchPaperByAbstract err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	rs = result.Data
	return
}

func RAGSearchPaperByFile(filename string, file io.Reader, accessToken string) (rs rag_model.RAGSearchPaperData, err error) {
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)

	fw1, _ := bw.CreateFormFile("file", filename)
	io.Copy(fw1, file)
	bw.Close()

	uri := APIHost + "/openapi/v1/rag/rag_search_paper_by_file"
	client := &http.Client{Timeout: 600 * time.Second}
	req, err := http.NewRequest("POST", uri, buf)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", bw.FormDataContentType())
	req.Header.Set("Authorization", accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var result rag_model.RAGSearchPaperResp
	if err = jsoniter.Unmarshal(body, &result); err != nil {
		log.Println("RAGSearchPaperByFile err, ", err)
		return
	}
	if result.ErrorCode != 0 {
		err = errors.New(result.ErrorMsg)
		return
	}
	rs = result.Data
	return
}
