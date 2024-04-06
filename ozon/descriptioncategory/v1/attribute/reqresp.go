package attribute

type ValuesRequest struct {
	AttributeID int64    `json:"attribute_id"`
	CategoryID  int64    `json:"category_id"`
	Language    Language `json:"language"`
	LastValueID int64    `json:"last_value_id"`
	Limit       int64    `json:"limit"`
	TypeID      uint64   `json:"type_id"`
}

type ValuesResponseResult struct {
	ID      int64  `json:"id"`
	Info    string `json:"info"`
	Picture string `json:"picture"`
	Value   string `json:"value"`
}

type ValuesResponse struct {
	Result  []ValuesResponseResult `json:"result"`
	HasNext bool                   `json:"has_next"`
}
