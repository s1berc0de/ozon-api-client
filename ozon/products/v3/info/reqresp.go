package info

type AttributesRequestFilter struct {
	OfferID    []string                          `json:"offer_id,omitempty"`
	ProductID  []string                          `json:"product_id"`
	Visibility AttributesRequestFilterVisibility `json:"visibility"`
}

type AttributesRequest struct {
	Filter  AttributesRequestFilter `json:"filter"`
	Limit   int64                   `json:"limit"`
	LastID  string                  `json:"last_id"`
	SortBy  string                  `json:"sort_by,omitempty"`
	SortDir string                  `json:"sort_dir"`
}

type AttributesResponseResultAttributeValue struct {
	DictionaryValueID int64  `json:"dictionary_value_id"`
	Value             string `json:"value"`
}

type AttributesResponseResultAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID   int64 `json:"complex_id"`
	Values      []AttributesResponseResultAttributeValue
}

type AttributesResponseResultComplexAttributeAttributeValue struct {
	DictionaryValueID int64  `json:"dictionary_value_id"`
	Value             string `json:"value"`
}

type AttributesResponseResultComplexAttributeAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID   int64 `json:"complex_id"`
	Values      []AttributesResponseResultComplexAttributeAttributeValue
}

type AttributesResponseResultComplexAttribute struct {
	Attributes []AttributesResponseResultComplexAttributeAttribute
}

type AttributesResponseResultImage struct {
	Default  bool   `json:"default"`
	FileName string `json:"file_name"`
	Index    int64  `json:"index"`
}

type AttributesResponseResultImage360 struct {
	FileName string `json:"file_name"`
	Index    int64  `json:"index"`
}

type AttributesResponseResultPDFItem struct {
	FileName string `json:"file_name"`
	Index    int64  `json:"index"`
	Name     string `json:"name"`
}

type AttributesResponseResult struct {
	Attributes        []AttributesResponseResultAttribute        `json:"attributes"`
	Barcode           string                                     `json:"barcode"`
	CategoryID        int64                                      `json:"category_id"`
	ColorImage        string                                     `json:"color_image"`
	ComplexAttributes []AttributesResponseResultComplexAttribute `json:"complex_attributes"`
	Depth             int32                                      `json:"depth"`
	DimensionUnit     AttributesResponseResultDimensionUnit      `json:"dimension_unit"`
	Height            int32                                      `json:"height"`
	ID                int64                                      `json:"id"`
	ImageGroupID      string                                     `json:"image_group_id"`
	Images            []AttributesResponseResultImage            `json:"images"`
	Images360         []AttributesResponseResultImage360         `json:"images360"`
	Name              string                                     `json:"name"`
	OfferID           string                                     `json:"offer_id"`
	PDFList           []AttributesResponseResultPDFItem          `json:"pdf_list"`
	Weight            int32                                      `json:"weight"`
	WeightUnit        string                                     `json:"weight_unit"`
	Width             int32                                      `json:"width"`
}

type AttributesResponse struct {
	Result []AttributesResponseResult `json:"result"`
	LastID string                     `json:"last_id"`
	Total  int32                      `json:"total"`
}
