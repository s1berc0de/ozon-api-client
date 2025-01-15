package info

import "time"

type LimitRequest struct{}

type LimitResponseDailyCreate struct {
	Limit   int64     `json:"limit"`
	ResetAt time.Time `json:"reset_at"`
	Usage   int64     `json:"usage"`
}

type LimitResponseDailyUpdate struct {
	Limit   int64     `json:"limit"`
	ResetAt time.Time `json:"reset_at"`
	Usage   int64     `json:"usage"`
}

type LimitResponseTotal struct {
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
}

type LimitResponse struct {
	DailyCreate LimitResponseDailyCreate `json:"daily_create"`
	DailyUpdate LimitResponseDailyUpdate `json:"daily_update"`
	Total       LimitResponseTotal       `json:"total"`
}

type PricesRequestFilter struct {
	OfferID    []string                      `json:"offer_id"`
	ProductID  []string                      `json:"product_id"`
	Visibility PricesRequestFilterVisibility `json:"visibility"`
}

type PricesRequest struct {
	Filter PricesRequestFilter `json:"filter"`
	LastID string              `json:"last_id"`
	Limit  int32               `json:"limit"`
}

type PricesResponseResultItemCommissions struct {
	FBODelivToCustomerAmount    float64 `json:"fbo_deliv_to_customer_amount"`
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FBOFulfillmentAmount        float64 `json:"fbo_fulfillment_amount"`
	FBOReturnFlowAmount         float64 `json:"fbo_return_flow_amount"`
	FBOReturnFlowTransMinAmount float64 `json:"fbo_return_flow_trans_min_amount"`
	FBOReturnFlowTransMaxAmount float64 `json:"fbo_return_flow_trans_max_amount"`
	FBSDelivToCustomerAmount    float64 `json:"fbs_deliv_to_customer_amount"`
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FBSFirstMileMinAmount       float64 `json:"fbs_first_mile_min_amount"`
	FBSFirstMileMaxAmount       float64 `json:"fbs_first_mile_max_amount"`
	FBSReturnFlowAmount         float64 `json:"fbs_return_flow_amount"`
	FBSReturnFlowTransMaxAmount float64 `json:"fbs_return_flow_trans_max_amount"`
	FBSReturnFlowTransMinAmount float64 `json:"fbs_return_flow_trans_min_amount"`
	SalesPercentFBO             float64 `json:"sales_percent_fbo"`
	SalesPercentFBS             float64 `json:"sales_percent_fbs"`
	SalesPercent                float64 `json:"sales_percent"`
}

type PricesResponseResultItemMarketingActionsAction struct {
	DateFrom      time.Time `json:"date_from"`
	DateTo        time.Time `json:"date_to"`
	DiscountValue string    `json:"discount_value"`
	Title         string    `json:"title"`
}

type PricesResponseResultItemMarketingActions struct {
	Actions           []PricesResponseResultItemMarketingActionsAction
	CurrentPeriodFrom time.Time `json:"current_period_from"`
	CurrentPeriodTo   time.Time `json:"current_period_to"`
	OzonActionsExist  bool      `json:"ozon_actions_exist"`
}

type PricesResponseResultItemPrice struct {
	AutoActionEnabled    bool                                      `json:"auto_action_enabled"`
	CurrencyCode         PricesResponseResultItemPriceCurrencyCode `json:"currency_code"`
	MarketingPrice       string                                    `json:"marketing_price"`
	MarketingSellerPrice string                                    `json:"marketing_seller_price"`
	MinOzonPrice         string                                    `json:"min_ozon_price"`
	MinPrice             string                                    `json:"min_price"`
	OldPrice             string                                    `json:"old_price"`
	PremiumPrice         string                                    `json:"premium_price"`
	Price                string                                    `json:"price"`
	RecommendedPrice     string                                    `json:"recommended_price"`
	RetailPrice          string                                    `json:"retail_price"`
	Vat                  string                                    `json:"vat"`
}

type PricesResponseResultItemPriceIndexesExternalIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type PricesResponseResultItemPriceIndexesOzonIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type PricesResponseResultItemPriceIndexesSelfMarketplacesIndexData struct {
	MinimalPrice         string  `json:"minimal_price"`
	MinimalPriceCurrency string  `json:"minimal_price_currency"`
	PriceIndexValue      float64 `json:"price_index_value"`
}

type PricesResponseResultItemPriceIndexes struct {
	ExternalIndexData         PricesResponseResultItemPriceIndexesExternalIndexData         `json:"external_index_data"`
	OzonIndexData             PricesResponseResultItemPriceIndexesOzonIndexData             `json:"ozon_index_data"`
	PriceIndex                PricesResponseResultItemPriceIndexesPriceIndex                `json:"price_index"`
	SelfMarketplacesIndexData PricesResponseResultItemPriceIndexesSelfMarketplacesIndexData `json:"self_marketplaces_index_data"`
}

type PricesResponseResultItem struct {
	Acquiring        int32                                     `json:"acquiring"`
	Commissions      PricesResponseResultItemCommissions       `json:"commissions"`
	MarketingActions *PricesResponseResultItemMarketingActions `json:"marketing_actions"`
	OfferID          string                                    `json:"offer_id"`
	Price            PricesResponseResultItemPrice             `json:"price"`
	PriceIndex       string                                    `json:"price_index"`
	PriceIndexes     PricesResponseResultItemPriceIndexes      `json:"price_indexes"`
	ProductID        int64                                     `json:"product_id"`
	VolumeWeight     float64                                   `json:"volume_weight"`
}

type PricesResponseResult struct {
	Items  []PricesResponseResultItem `json:"items"`
	LastID string                     `json:"last_id"`
	Total  int32                      `json:"total"`
}

type PricesResponse struct {
	Result PricesResponseResult `json:"result"`
}

type StocksRequestFilterWithQuant struct {
	Created bool `json:"created"`
	Exists  bool `json:"exists"`
}

type StocksRequestFilter struct {
	OfferID    []string                      `json:"offer_id"`
	ProductID  []string                      `json:"product_id"`
	Visibility StocksRequestFilterVisibility `json:"visibility"`
	WithQuant  StocksRequestFilterWithQuant  `json:"with_quant"`
}

type StocksRequest struct {
	Cursor string              `json:"cursor"`
	Filter StocksRequestFilter `json:"filter"`
	Limit  int32               `json:"limit"`
}

type StocksResponseItemStock struct {
	Present      int32                               `json:"present"`
	Reserved     int32                               `json:"reserved"`
	ShipmentType StocksResponseItemStockShipmentType `json:"shipment_type"`
	SKU          int64                               `json:"sku"`
	Type         string                              `json:"type"`
}

type StocksResponseItem struct {
	OfferID   string                    `json:"offer_id"`
	ProductID int64                     `json:"product_id"`
	Stocks    []StocksResponseItemStock `json:"stocks"`
}

type StocksResponse struct {
	Cursor string               `json:"cursor"`
	Items  []StocksResponseItem `json:"items"`
	Total  int32                `json:"total"`
}
