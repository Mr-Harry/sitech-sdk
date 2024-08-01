package rag_model

import "gitee.com/beijing-turing-tuxian/sitech-sdk/go-sdk/models/result_model"

type RAGSearchPaperResp struct {
	result_model.Result
	Data RAGSearchPaperData `json:"data"`
}

type RAGSearchPaperData struct {
	Title               string    `json:"title"`
	Abstract            string    `json:"abstract"`
	ReferenceTotal      int       `json:"reference_total"`
	RelatedTotal        int       `json:"related_total"`
	KnowledgeGraphTotal int       `json:"knowledge_graph_total"`
	ReferenceList       []PubItem `json:"reference_list"`
	RelatedList         []PubItem `json:"related_list"`
	KnowledgeGraph      []PubItem `json:"knowledge_graph"`
}

type PubItem struct {
	ID     string       `json:"id"`
	Detail PublishField `json:"detail"`
}

type PublishField struct {
	Abstract    string   `json:"abstract"`
	AuthorName1 string   `json:"author_name1"`
	AuthorName2 string   `json:"author_name2"`
	AuthorName3 string   `json:"author_name3"`
	Pdf         string   `json:"pdf"`
	Url         string   `json:"url"`
	AuthorNames []string `json:"author_names"`
}
