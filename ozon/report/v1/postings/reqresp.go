package postings

import "time"

type CreateRequestFilter struct {
	ProcessedAtFrom time.Time `json:"processed_at_from"`
	ProcessedAtTo   time.Time `json:"processed_at_to"`
	DeliverySchema  []string  `json:"delivery_schema"`
	Sku             []int64   `json:"sku"`
	CancelReasonId  []int64   `json:"cancel_reason_id"`
	OfferId         string    `json:"offer_id"`
	StatusAlias     []string  `json:"status_alias"`
	Statuses        []int64   `json:"statuses"`
	Title           string    `json:"title"`
}

type CreateRequest struct {
	Filter   CreateRequestFilter `json:"filter"`
	Language string              `json:"language"`
}

type CreateResponseResult struct {
	Code string `json:"code"`
}

type CreateResponse struct {
	Result CreateResponseResult `json:"result"`
}
