package info

type DescriptionRequest struct {
	OfferID   string `json:"offer_id"`
	ProductID int64  `json:"product_id"`
}

type DescriptionResponseResult struct {
	Description string `json:"description"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	OfferID     string `json:"offer_id"`
}

type DescriptionResponse struct {
	Result DescriptionResponseResult `json:"result"`
}

type SubscriptionRequest struct {
	SKUs []string `json:"skus"`
}

type SubscriptionResponseResult struct {
	Count int64 `json:"count"`
	SKU   int64 `json:"sku"`
}

type SubscriptionResponse struct {
	Result []SubscriptionResponseResult `json:"result"`
}
