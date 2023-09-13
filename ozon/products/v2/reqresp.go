package v2

type DeleteRequestProduct struct {
	OfferID string `json:"offer_id"`
}

type DeleteRequest struct {
	Products []DeleteRequestProduct `json:"products"`
}

type DeleteResponseStatus struct {
	Error     string `json:"error"`
	IsDeleted bool   `json:"is_deleted"`
	OfferID   string `json:"offer_id"`
}

type DeleteResponse struct {
	Status []DeleteResponseStatus `json:"status"`
}

type StocksRequestStock struct {
	OfferID     string `json:"offer_id"`
	ProductID   int64  `json:"product_id"`
	Stock       int64  `json:"stock"`
	WarehouseID int64  `json:"warehouse_id"`
}

type StocksRequest struct {
	Stocks []StocksRequestStock `json:"stocks"`
}

type StocksResponseResultError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type StocksResponseResult struct {
	Errors      []StocksResponseResultError `json:"errors"`
	OfferID     string                      `json:"offer_id"`
	ProductID   int64                       `json:"product_id"`
	Updated     bool                        `json:"updated"`
	WarehouseID int64                       `json:"warehouse_id"`
}

type StocksResponse struct {
	Result []StocksResponseResult `json:"result"`
}
