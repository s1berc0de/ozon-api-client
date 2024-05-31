package v3

type ImportItemAttributeValue struct {
	DictionaryValueID *int64 `json:"dictionary_value_id,omitempty"`
	Value             string `json:"value"`
}

type ImportItemAttribute struct {
	ComplexID int64                      `json:"complex_id"`
	ID        int64                      `json:"id"`
	Values    []ImportItemAttributeValue `json:"values"`
}

type ComplexAttributeValue struct {
	Value string `json:"value"`
}

type ComplexAttributeItem struct {
	ComplexID int64                   `json:"complex_id"`
	ID        int64                   `json:"id"`
	Values    []ComplexAttributeValue `json:"values"`
}

type ComplexAttribute struct {
	Attributes []ComplexAttributeItem `json:"attributes"`
}

type ImportItem struct {
	Attributes            []ImportItemAttribute `json:"attributes"`
	Barcode               string                `json:"barcode"`
	DescriptionCategoryID int64                 `json:"description_category_id"`
	ColorImage            string                `json:"color_image"`
	ComplexAttributes     []ComplexAttribute    `json:"complex_attributes"`
	CurrencyCode          string                `json:"currency_code"`
	Depth                 int64                 `json:"depth"`
	DimensionUnit         string                `json:"dimension_unit"`
	Height                int64                 `json:"height"`
	Images                []string              `json:"images"`
	Images360             []string              `json:"images360"`
	Name                  string                `json:"name"`
	OfferID               string                `json:"offer_id"`
	OldPrice              string                `json:"old_price"`
	PdfList               []string              `json:"pdf_list"`
	PremiumPrice          string                `json:"premium_price"`
	Price                 string                `json:"price"`
	PrimaryImage          string                `json:"primary_image"`
	Vat                   string                `json:"vat"`
	Weight                int64                 `json:"weight"`
	WeightUnit            string                `json:"weight_unit"`
	Width                 int64                 `json:"width"`
}

type ImportResult struct {
	TaskID int64 `json:"task_id"`
}

type ImportRequest struct {
	Items []ImportItem `json:"items"`
}

type ImportResponse struct {
	Result ImportResult `json:"result"`
}
