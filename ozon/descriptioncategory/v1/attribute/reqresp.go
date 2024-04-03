package attribute

type AttributeRequest struct {
	DescriptionCategoryID int64    `json:"description_category_id"`
	Language              Language `json:"language"`
	TypeID                int64    `json:"type_id"`
}

type AttributeResult struct {
	CategoryDependent  bool   `json:"category_dependent"`
	Description        string `json:"description"`
	DictionaryID       int64  `json:"dictionary_id"`
	GroupID            int64  `json:"group_id"`
	GroupName          string `json:"group_name"`
	ID                 int64  `json:"id"`
	IsAspect           bool   `json:"is_aspect"`
	IsCollection       bool   `json:"is_collection"`
	IsRequired         bool   `json:"is_required"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	AttributeComplexID int64  `json:"attribute_complex_id"`
	MaxValueCount      int64  `json:"max_value_count"`
}

type AttributeResponse struct {
	Result []AttributeResult `json:"result"`
}

type ValuesRequest struct {
	AttributeID           int64    `json:"attribute_id"`
	DescriptionCategoryID int64    `json:"description_category_id"`
	Language              Language `json:"language"`
	LastValueID           int64    `json:"last_value_id"`
	Limit                 int64    `json:"limit"`
	TypeID                uint64   `json:"type_id"`
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
