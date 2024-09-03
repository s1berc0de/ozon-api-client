package values

type SearchRequest struct {
	AttributeId           int64  `json:"attribute_id"`
	DescriptionCategoryID int64  `json:"description_category_id"`
	Limit                 int64  `json:"limit"`
	TypeID                int64  `json:"type_id"`
	Value                 string `json:"value"`
}

type SearchResponseResult struct {
	ID      int64  `json:"id"`
	Info    string `json:"info"`
	Picture string `json:"picture"`
	Value   string `json:"value"`
}

type SearchResponse struct {
	Result []SearchResponseResult `json:"result"`
}
