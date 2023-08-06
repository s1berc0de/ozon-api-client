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
	Status DeleteResponseStatus `json:"status"`
}
