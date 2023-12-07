package info

type StocksRequestFilter struct {
	OfferID    []string                      `json:"offer_id"`
	ProductID  []string                      `json:"product_id"`
	Visibility StocksRequestFilterVisibility `json:"visibility"`
}

type StocksRequest struct {
	Filter StocksRequestFilter `json:"filter"`
	LastID string              `json:"last_id"`
	Limit  int64               `json:"limit"`
}

type StocksResponseResultItemStock struct {
	Present  int32  `json:"present"`
	Reserved int32  `json:"reserved"`
	Type     string `json:"type"`
}

type StocksResponseResultItem struct {
	OfferID   string                          `json:"offer_id"`
	ProductID int64                           `json:"product_id"`
	Stocks    []StocksResponseResultItemStock `json:"stocks"`
}

type StocksResponseResult struct {
	Items  []StocksResponseResultItem `json:"items"`
	LastID string                     `json:"last_id"`
	Total  int32                      `json:"total"`
}

type StocksResponse struct {
	Result StocksResponseResult `json:"result"`
}
