package gosdk

import (
	"io"

	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/chat"
	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/rag_model"
	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/oauth"
	"gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/papers"
)

type SiTechClient struct {
	AppKey    string
	AppSecret string
}

func NewSiTechClient(appKey, appSecret string) (c *SiTechClient) {
	return &SiTechClient{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (c *SiTechClient) GetAccessToken() (token string, err error) {
	return oauth.GetAccessToken(c.AppKey, c.AppSecret)
}

func (c *SiTechClient) RAGSearchPaper(query string) (result rag_model.RAGSearchPaperData, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.RAGSearchPaper(query, token)
	return
}

func (c *SiTechClient) RAGSearchPaperByDate(query, startTime, endTime string) (result rag_model.RAGSearchPaperData, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.RAGSearchPaperByDate(query, startTime, endTime, token)
	return
}

func (c *SiTechClient) RAGSearchPaperByDoiOrUrl(doi, uri string) (result rag_model.RAGSearchPaperData, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.RAGSearchPaperByDoiOrUrl(doi, uri, token)
	return
}

func (c *SiTechClient) RAGSearchPaperByAbstract(abstract string) (result rag_model.RAGSearchPaperData, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.RAGSearchPaperByAbstract(abstract, token)
	return
}

func (c *SiTechClient) RAGSearchPaperByFile(filename string, file io.Reader) (result rag_model.RAGSearchPaperData, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.RAGSearchPaperByFile(filename, file, token)
	return
}

func (c *SiTechClient) SearchPubsByKwd(kwds []string, withKnowledgeGraph bool) (result []rag_model.PubItem, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.SearchPubsByKwd(kwds, withKnowledgeGraph, token)
	return
}

func (c *SiTechClient) GetPubsCount() (result int, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.GetPubsCount(token)
	return
}

func (c *SiTechClient) TextSummary(abstracts []string) (result string, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = papers.TextSummary(abstracts, token)
	return
}

func (c *SiTechClient) SendMessage(prompt string) (result string, err error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return
	}
	result, err = chat.SendMessage(prompt, token)
	return
}
