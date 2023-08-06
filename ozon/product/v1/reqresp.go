package v1

type ImportBySKURequestItem struct {
	SKU          int64                              `json:"sku"`
	Name         string                             `json:"name"`
	OfferID      string                             `json:"offer_id"`
	CurrencyCode ImportBySKURequestItemCurrencyCode `json:"currency_code"`
	OldPrice     string                             `json:"old_price"`
	Price        string                             `json:"price"`
	PremiumPrice string                             `json:"premium_price"`
	Vat          string                             `json:"vat"`
}

type ImportBySKURequest struct {
	Items []ImportBySKURequestItem `json:"items"`
}

type ImportBySKUResponseResult struct {
	TaskID           int64   `json:"task_id"`
	UnmatchedSKUList []int64 `json:"unmatched_sku_list"`
}

type ImportBySKUResponse struct {
	Result ImportBySKUResponseResult `json:"result"`
}

type RatingBySKURequest struct {
	SKUs []string `json:"skus"`
}

type RatingBySKUResponseProductGroupCondition struct {
	Cost        float32 `json:"cost"`
	Description string  `json:"description"`
	Fulfilled   bool    `json:"fulfilled"`
	Key         string  `json:"key"`
}

type RatingBySKUResponseProductImproveAttribute struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RatingBySKUResponseProductGroup struct {
	Conditions        []RatingBySKUResponseProductGroupCondition   `json:"conditions"`
	ImproveAtLeast    int32                                        `json:"improve_at_least"`
	ImproveAttributes []RatingBySKUResponseProductImproveAttribute `json:"improve_attributes"`
	Key               string                                       `json:"key"`
	Name              string                                       `json:"name"`
	Rating            float32                                      `json:"rating"`
	Weight            float32                                      `json:"weight"`
}

type RatingBySKUResponseProduct struct {
	SKU    int64                             `json:"sku"`
	Rating float32                           `json:"rating"`
	Groups []RatingBySKUResponseProductGroup `json:"groups"`
}

type RatingBySKUResponse struct {
	Products []RatingBySKUResponseProduct `json:"products"`
}

type ArchiveRequest struct {
	ProductID []int64 `json:"product_id"`
}

type ArchiveResponse struct {
	Result bool `json:"result"`
}

type UnArchiveRequest struct {
	ProductID []int64 `json:"product_id"`
}

type UnArchiveResponse struct {
	Result bool `json:"result"`
}

type UploadDigitalCodesRequest struct {
	DigitalCodes []string `json:"digital_codes"`
	ProductID    int64    `json:"product_id"`
}

type UploadDigitalCodesResponseResult struct {
	TaskID int64 `json:"task_id"`
}

type UploadDigitalCodesResponse struct {
	Result UploadDigitalCodesResponseResult `json:"result"`
}
