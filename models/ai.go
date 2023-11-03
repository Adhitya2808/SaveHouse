package models

type GudangRequest struct {
	Text string `json:"text" form:"text"`
}

type AIResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
