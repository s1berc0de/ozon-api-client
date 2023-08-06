package v3

type AttributeRequest struct {
	AttributeType AttributeType `json:"attribute_type" default:"ALL"`
	CategoryId    []int64       `json:"category_id"`
	Language      Language      `json:"language" default:"DEFAULT"`
}

type AttributeResponseResultAttribute struct {
	CategoryDependent bool   `json:"category_dependent"`
	Description       string `json:"description"`
	DictionaryId      int64  `json:"dictionary_id"`
	GroupId           int64  `json:"group_id"`
	GroupName         string `json:"group_name"`
	ID                int64  `json:"id"`
	IsAspect          bool   `json:"is_aspect"`
	IsCollection      bool   `json:"is_collection"`
	IsRequired        bool   `json:"is_required"`
	Name              string `json:"name"`
	Type              string `json:"type"`
}

type AttributeResponseResult struct {
	Attributes []AttributeResponseResultAttribute `json:"attributes"`
	CategoryId int64                              `json:"category_id"`
}

type AttributeResponse struct {
	Result []AttributeResponseResult `json:"result"`
}
