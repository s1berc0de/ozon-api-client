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

type DiscountedRequest struct {
	DiscountedSKUs []string `json:"discounted_skus"`
}

type DiscountedResponseItem struct {
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	Condition            string `json:"condition"`
	ConditionEstimation  string `json:"condition_estimation"`
	Defects              string `json:"defects"`
	DiscountedSKU        int64  `json:"discounted_sku"`
	MechanicalDamage     string `json:"mechanical_damage"`
	PackageDamage        string `json:"package_damage"`
	PackagingViolation   string `json:"packaging_violation"`
	ReasonDamaged        string `json:"reason_damaged"`
	Repair               string `json:"repair"`
	Shortage             string `json:"shortage"`
	SKU                  int64  `json:"sku"`
	WarrantyType         string `json:"warranty_type"`
}

type DiscountedResponse struct {
	Items []DiscountedResponseItem `json:"items"`
}
