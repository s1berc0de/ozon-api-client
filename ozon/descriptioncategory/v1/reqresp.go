package v1

type TreeRequest struct {
	Language Language `json:"language"`
}

type TreeResult struct {
	DescriptionCategoryID *int64        `json:"description_category_id,omitempty"`
	CategoryName          *string       `json:"category_name,omitempty"`
	TypeName              *string       `json:"type_name,omitempty"`
	TypeID                *int64        `json:"type_id,omitempty"`
	Disabled              bool          `json:"disabled"`
	Children              []*TreeResult `json:"children,omitempty"`
}

type TreeResponse struct {
	Result []*TreeResult `json:"result"`
}

type AttributeRequest struct {
	DescriptionCategoryID int64    `json:"description_category_id"`
	Language              Language `json:"language"`
	TypeID                int64    `json:"type_id"`
}

type AttributeResponseResult struct {
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
	Result []AttributeResponseResult `json:"result"`
}
