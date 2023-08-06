package attribute

type ValuesRequest struct {
	AttributeId int64    `json:"attribute_id"`
	CategoryId  int64    `json:"category_id"`
	Language    Language `json:"language" default:"DEFAULT"`
	LastValueId int64    `json:"last_value_id"`
	Limit       int64    `json:"limit"`
}

type ValuesResponseResult struct {
	ID      int64  `json:"id"`
	Info    string `json:"info"`
	Picture string `json:"picture"`
	Value   string `json:"value"`
}

type ValuesResponse struct {
	HasNext bool                   `json:"has_next"`
	Result  []ValuesResponseResult `json:"result"`
}
