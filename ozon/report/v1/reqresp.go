package v1

import "time"

type InfoRequest struct {
	Code string `json:"code"`
}

type InfoResponse struct {
	Result InfoResponseResult `json:"result"`
}

type InfoResponseResult struct {
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	Error      string    `json:"error"`
	File       string    `json:"file"`
	ReportType string    `json:"report_type"`
	CreatedAt  time.Time `json:"created_at"`
}
