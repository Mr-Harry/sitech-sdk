package result_model

type Result struct {
	ErrorCode int `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
}