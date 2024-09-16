package v2

import "time"

type ImportRequestItemAttributeValue struct {
	DictionaryValueID int64  `json:"dictionary_value_id,omitempty"`
	Value             string `json:"value"`
}

type ImportRequestItemAttribute struct {
	ComplexID int64                             `json:"complex_id"`
	ID        int64                             `json:"id"`
	Values    []ImportRequestItemAttributeValue `json:"values"`
}

type ImportRequestItemComplexAttributeAttributeValue struct {
	DictionaryValueID int64  `json:"dictionary_value_id"`
	Value             string `json:"value"`
}

type ImportRequestItemComplexAttributeAttribute struct {
	ComplexID int64                                             `json:"complex_id"`
	ID        int64                                             `json:"id"`
	Values    []ImportRequestItemComplexAttributeAttributeValue `json:"values"`
}

type ImportRequestItemComplexAttribute struct {
	Attributes []ImportRequestItemComplexAttributeAttribute `json:"attributes"`
}

type ImportRequestItemPDFItem struct {
	Index  int64  `json:"index"`
	Name   string `json:"name"`
	SrcURL string `json:"src_url"`
}

type ImportRequestItem struct {
	Attributes        []ImportRequestItemAttribute        `json:"attributes"`
	Barcode           string                              `json:"barcode"`
	CategoryID        int64                               `json:"category_id"`
	ColorImage        string                              `json:"color_image"`
	ComplexAttributes []ImportRequestItemComplexAttribute `json:"complex_attributes"`
	CurrencyCode      ImportRequestItemCurrencyCode       `json:"currency_code"`
	Depth             int32                               `json:"depth"`
	DimensionUnit     ImportRequestItemDimensionUnit      `json:"dimension_unit"`
	GeoNames          []string                            `json:"geo_names,omitempty"`
	Height            int32                               `json:"height"`
	Images            []string                            `json:"images"`
	ServiceType       ImportRequestItemServiceType        `json:"service_type,omitempty"`
	Images360         []string                            `json:"images360"`
	Name              string                              `json:"name"`
	OfferID           string                              `json:"offer_id"`
	OldPrice          string                              `json:"old_price"`
	PDFList           []ImportRequestItemPDFItem          `json:"pdf_list"`
	PremiumPrice      string                              `json:"premium_price"`
	Price             string                              `json:"price"`
	PrimaryImage      string                              `json:"primary_image"`
	Vat               string                              `json:"vat"`
	Weight            int32                               `json:"weight"`
	WeightUnit        ImportRequestItemWeightUnit         `json:"weight_unit"`
	Width             int32                               `json:"width"`
}

type ImportRequest struct {
	Items []ImportRequestItem `json:"items"`
}

type ImportResponseResult struct {
	TaskID int64 `json:"task_id"`
}

type ImportResponse struct {
	Result ImportResponseResult `json:"result"`
}

type ListRequestFilter struct {
	OfferID    []string                    `json:"offer_id"`
	ProductID  []string                    `json:"product_id"`
	Visibility ListRequestFilterVisibility `json:"visibility"`
}

type ListRequest struct {
	Filter ListRequestFilter `json:"filter"`
	LastID string            `json:"last_id"`
	Limit  int64             `json:"limit"`
}

type ListResponseResultItem struct {
	OfferID   string `json:"offer_id"`
	ProductID int64  `json:"product_id"`
}

type ListResponseResult struct {
	Items  []ListResponseResultItem `json:"items"`
	LastID string                   `json:"last_id"`
	Total  int32                    `json:"total"`
}

type ListResponse struct {
	Result ListResponseResult `json:"result"`
}

type InfoRequest struct {
	OfferID   string `json:"offer_id"`
	ProductID int64  `json:"product_id"`
	SKU       int64  `json:"sku"`
}

type InfoResponseResultCommission struct {
	DeliveryAmount float64 `json:"delivery_amount"`
	MinValue       float64 `json:"min_value"`
	Percent        float64 `json:"percent"`
	ReturnAmount   float64 `json:"return_amount"`
	SaleSchema     string  `json:"sale_schema"`
	Value          float64 `json:"value"`
}

type InfoResponseResultDiscountedStocks struct {
	Coming   int32 `json:"coming"`
	Present  int32 `json:"present"`
	Reserved int32 `json:"reserved"`
}

type InfoResponseResultPriceIndexesExternalIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type InfoResponseResultPriceIndexesOzonIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type InfoResponseResultPriceIndexesSelfMarketplacesIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type InfoResponseResultPriceIndexes struct {
	ExternalIndexData         InfoResponseResultPriceIndexesExternalIndexData         `json:"external_index_data"`
	OzonIndexData             InfoResponseResultPriceIndexesOzonIndexData             `json:"ozon_index_data"`
	PriceIndex                InfoResponseResultPriceIndexesPriceIndex                `json:"price_index"`
	SelfMarketplacesIndexData InfoResponseResultPriceIndexesSelfMarketplacesIndexData `json:"self_marketplaces_index_data"`
}

type InfoResponseResultStatusItemErrorOptionalDescriptionElements struct {
	PropertyName string `json:"property name*"`
}

type InfoResponseResultStatusItemError struct {
	Code                        string                                                       `json:"code"`
	State                       string                                                       `json:"state"`
	Level                       string                                                       `json:"level"`
	Description                 string                                                       `json:"description"`
	Field                       string                                                       `json:"field"`
	AttributeID                 int64                                                        `json:"attribute_id"`
	AttributeName               string                                                       `json:"attribute_name"`
	OptionalDescriptionElements InfoResponseResultStatusItemErrorOptionalDescriptionElements `json:"optional_description_elements"`
}

type InfoResponseResultStatus struct {
	State            string                              `json:"state"`
	StateFailed      string                              `json:"state_failed"`
	ModerateStatus   string                              `json:"moderate_status"`
	DeclineReasons   []string                            `json:"decline_reasons"`
	ValidationState  string                              `json:"validation_state"`
	StateName        string                              `json:"state_name"`
	StateDescription string                              `json:"state_description"`
	IsFailed         bool                                `json:"is_failed"`
	IsCreated        bool                                `json:"is_created"`
	StateTooltip     string                              `json:"state_tooltip"`
	ItemErrors       []InfoResponseResultStatusItemError `json:"item_errors"`
	StateUpdatedAt   time.Time                           `json:"state_updated_at"`
}

type InfoResponseResultSource struct {
	IsEnabled bool   `json:"is_enabled"`
	SKU       int64  `json:"sku"`
	Source    string `json:"source"`
}

type InfoResponseResultStocks struct {
	Coming   int32 `json:"coming"`
	Present  int32 `json:"present"`
	Reserved int32 `json:"reserved"`
}

type InfoResponseResultVisibilityDetails struct {
	ActiveProduct bool `json:"active_product"`
	HasPrice      bool `json:"has_price"`
	HasStock      bool `json:"has_stock"`
}

type InfoResponseResult struct {
	Barcode             string                              `json:"barcode"`
	Barcodes            []string                            `json:"barcodes"`
	BuyboxPrice         string                              `json:"buybox_price"`
	CategoryID          int64                               `json:"category_id"`
	ColorImage          string                              `json:"color_image"`
	Commissions         []InfoResponseResultCommission      `json:"commissions"`
	CreatedAt           time.Time                           `json:"created_at"`
	SKU                 int64                               `json:"sku"`
	FBOSKU              int64                               `json:"fbo_sku"`
	FBSSKU              int64                               `json:"fbs_sku"`
	ID                  int64                               `json:"id"`
	Images              []string                            `json:"images"`
	PrimaryImage        string                              `json:"primary_image"`
	Images360           []string                            `json:"images360"`
	HasDiscountedItem   bool                                `json:"has_discounted_item"`
	IsDiscounted        bool                                `json:"is_discounted"`
	DiscountedStocks    InfoResponseResultDiscountedStocks  `json:"discounted_stocks"`
	IsKGT               bool                                `json:"is_kgt"`
	IsPrepayment        bool                                `json:"is_prepayment"`
	IsPrepaymentAllowed bool                                `json:"is_prepayment_allowed"`
	CurrencyCode        InfoResponseResultCurrencyCode      `json:"currency_code"`
	MarketingPrice      string                              `json:"marketing_price"`
	MinOzonPrice        string                              `json:"min_ozon_price"`
	MinPrice            string                              `json:"min_price"`
	Name                string                              `json:"name"`
	OfferID             string                              `json:"offer_id"`
	OldPrice            string                              `json:"old_price"`
	PremiumPrice        string                              `json:"premium_price"`
	Price               string                              `json:"price"`
	PriceIndex          string                              `json:"price_index"`
	PriceIndexes        InfoResponseResultPriceIndexes      `json:"price_indexes"`
	RecommendedPrice    string                              `json:"recommended_price"`
	Status              InfoResponseResultStatus            `json:"status"`
	Sources             []InfoResponseResultSource          `json:"sources"`
	Stocks              InfoResponseResultStocks            `json:"stocks"`
	UpdatedAt           time.Time                           `json:"updated_at"`
	Vat                 string                              `json:"vat"`
	VisibilityDetails   InfoResponseResultVisibilityDetails `json:"visibility_details"`
	Visible             bool                                `json:"visible"`
	VolumeWeight        float64                             `json:"volume_weight"`
}

type InfoResponse struct {
	Result InfoResponseResult `json:"result"`
}
