package info

import "time"

type ListRequest struct {
	OfferID   []string `json:"offer_id"`
	ProductID []int64  `json:"product_id"`
	SKU       []int64  `json:"sku"`
}

type ListResponseResultItemDiscountedStock struct {
	Coming   int32 `json:"coming"`
	Present  int32 `json:"present"`
	Reserved int32 `json:"reserved"`
}

type ListResponseResultItemPriceIndexesExternalIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type ListResponseResultItemPriceIndexesOzonIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type ListResponseResultItemPriceIndexesSelfMarketplacesIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type ListResponseResultItemPriceIndexes struct {
	ExternalIndexData         ListResponseResultItemPriceIndexesExternalIndexData         `json:"external_index_data"`
	OzonIndexData             ListResponseResultItemPriceIndexesOzonIndexData             `json:"ozon_index_data"`
	PriceIndex                ListResponseResultItemPriceIndexesPriceIndex                `json:"price_index"`
	SelfMarketplacesIndexData ListResponseResultItemPriceIndexesSelfMarketplacesIndexData `json:"self_marketplaces_index_data"`
}

type ListResponseResultItemStatusItemErrorOptionalDescriptionElements struct {
	PropertyName string `json:"property name*"`
}

type ListResponseResultItemStatusItemError struct {
	Code                        string                                                           `json:"code"`
	State                       string                                                           `json:"state"`
	Level                       string                                                           `json:"level"`
	Description                 string                                                           `json:"description"`
	Field                       string                                                           `json:"field"`
	AttributeID                 int64                                                            `json:"attribute_id"`
	AttributeName               string                                                           `json:"attribute_name"`
	OptionalDescriptionElements ListResponseResultItemStatusItemErrorOptionalDescriptionElements `json:"optional_description_elements"`
}

type ListResponseResultItemStatus struct {
	State            string                                  `json:"state"`
	StateFailed      string                                  `json:"state_failed"`
	ModerateStatus   string                                  `json:"moderate_status"`
	DeclineReasons   []string                                `json:"decline_reasons"`
	ValidationState  string                                  `json:"validation_state"`
	StateName        string                                  `json:"state_name"`
	StateDescription string                                  `json:"state_description"`
	IsFailed         bool                                    `json:"is_failed"`
	IsCreated        bool                                    `json:"is_created"`
	StateTooltip     string                                  `json:"state_tooltip"`
	ItemErrors       []ListResponseResultItemStatusItemError `json:"item_errors"`
	StateUpdatedAt   time.Time                               `json:"state_updated_at"`
}

type ListResponseResultItemSource struct {
	IsEnabled bool   `json:"is_enabled"`
	SKU       int64  `json:"sku"`
	Source    string `json:"source"`
}

type ListResponseResultItemStocks struct {
	Coming   int32 `json:"coming"`
	Present  int32 `json:"present"`
	Reserved int32 `json:"reserved"`
}

type ListResponseResultItemVisibilityDetailsReason struct {
	Description string `json:"description"`
	ID          int64  `json:"id"`
}

type ListResponseResultItemVisibilityDetails struct {
	ActiveProduct bool                                            `json:"active_product"`
	HasPrice      bool                                            `json:"has_price"`
	HasStock      bool                                            `json:"has_stock"`
	Reasons       []ListResponseResultItemVisibilityDetailsReason `json:"reasons"`
}

type ListResponseResultItem struct {
	Barcode           string                                  `json:"barcode"`
	Barcodes          []string                                `json:"barcodes"`
	BuyboxPrice       string                                  `json:"buybox_price"`
	CategoryID        int64                                   `json:"category_id"`
	ColorImage        string                                  `json:"color_image"`
	CreatedAt         time.Time                               `json:"created_at"`
	SKU               int64                                   `json:"sku"`
	FBOSKU            int64                                   `json:"fbo_sku"`
	FBSSKU            int64                                   `json:"fbs_sku"`
	ID                int64                                   `json:"id"`
	Images            []string                                `json:"images"`
	PrimaryImage      string                                  `json:"primary_image"`
	Images360         []string                                `json:"images360"`
	HasDiscountedItem bool                                    `json:"has_discounted_item"`
	IsDiscounted      bool                                    `json:"is_discounted"`
	DiscountedStocks  ListResponseResultItemDiscountedStock   `json:"discounted_stocks"`
	IsKGT             bool                                    `json:"is_kgt"`
	CurrencyCode      ListResponseResultItemCurrencyCode      `json:"currency_code"`
	MarketingPrice    string                                  `json:"marketing_price"`
	MinOzonPrice      string                                  `json:"min_ozon_price"`
	MinPrice          string                                  `json:"min_price"`
	Name              string                                  `json:"name"`
	OfferID           string                                  `json:"offer_id"`
	OldPrice          string                                  `json:"old_price"`
	PremiumPrice      string                                  `json:"premium_price"`
	Price             string                                  `json:"price"`
	PriceIndex        string                                  `json:"price_index"`
	PriceIndexes      ListResponseResultItemPriceIndexes      `json:"price_indexes"`
	RecommendedPrice  string                                  `json:"recommended_price"`
	Status            ListResponseResultItemStatus            `json:"status"`
	Sources           []ListResponseResultItemSource          `json:"sources"`
	Stocks            ListResponseResultItemStocks            `json:"stocks"`
	UpdatedAt         time.Time                               `json:"updated_at"`
	Vat               string                                  `json:"vat"`
	VisibilityDetails ListResponseResultItemVisibilityDetails `json:"visibility_details"`
	Visible           bool                                    `json:"visible"`
}

type ListResponseResult struct {
	Items []ListResponseResultItem `json:"items"`
}

type ListResponse struct {
	Result ListResponseResult `json:"result"`
}
