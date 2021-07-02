package dto

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type SuccessResponses struct {
	Total  int64       `json:"total"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
