package update

type OfferIDRequestUpdateOfferIDItem struct {
	NewOfferID string `json:"new_offer_id"`
	OfferID    string `json:"offer_id"`
}

type OfferIDRequest struct {
	UpdateOfferID []OfferIDRequestUpdateOfferIDItem `json:"update_offer_id"`
}

type OfferIDResponseError struct {
	Message string `json:"message"`
	OfferID string `json:"offer_id"`
}

type OfferIDResponse struct {
	Errors []OfferIDResponseError `json:"errors"`
}
