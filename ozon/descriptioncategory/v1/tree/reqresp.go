package tree

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
	Result []TreeResult `json:"result"`
}
