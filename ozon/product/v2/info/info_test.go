package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/andmetoo/ozon-api-client/ozon/product/v2/info"
	"github.com/stretchr/testify/require"
)

func TestList_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/product/info/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"offer_id":["23","010"],"product_id":[],"sku":[]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  	"result": {
								"items": [
							  		{
										"id": 78712196,
										"name": "Как выбрать детские музыкальные инструменты. Ксилофон, бубен, маракасы и другие инструменты для детей до 6 лет. Мастер-класс о раннем музыкальном развитии от Монтессори-педагога",
										"offer_id": "010",
										"barcode": "",
										"barcodes": [
								  			"2335900005",
								  			"7533900005"
										],
										"buybox_price": "",
										"category_id": 93726157,
										"created_at": "2021-06-03T03:40:05.871465Z",
										"images": [],
										"currency_code": "RUB",
										"marketing_price": "",
										"min_price": "",
										"old_price": "1000.0000",
										"premium_price": "590.0000",
										"price": "690.0000",
										"recommended_price": "",
										"sources": [
								  			{
												"is_enabled": true,
												"sku": 269628393,
												"source": "fbo"
								  			},
								  			{
												"is_enabled": true,
												"sku": 269628396,
												"source": "fbs"
								  			}
										],
										"has_discounted_item": true,
										"is_discounted": true,
										"discounted_stocks": {
								  			"coming": 0,
								  			"present": 0,
								  			"reserved": 0
										},
										"state": "",
										"stocks": {
								  			"coming": 0,
								  			"present": 13,
								  			"reserved": 0
										},
										"errors": [],
										"updated_at": "2023-02-09T06:46:44.152Z",
										"vat": "0.0",
										"visible": true,
										"visibility_details": {
											"has_price": false,
											"has_stock": true,
											"active_product": false,
											"reasons": {
												"field1": {
													"reasons": [
														{
															"id": 1, 
															"description": "descr"
														}
													]
												}
											}
										},
										"price_indexes": {
								  			"external_index_data": {
												"minimal_price": "string",
												"minimal_price_currency": "string",
												"price_index_value": 0
											},
											"ozon_index_data": {
												"minimal_price": "string",
												"minimal_price_currency": "string",
												"price_index_value": 0
											},
											"price_index": "WITHOUT_INDEX",
											"self_marketplaces_index_data": {
												"minimal_price": "string",
												"minimal_price_currency": "string",
												"price_index_value": 0
											}
										},
										"images360": [],
										"is_kgt": false,
										"color_image": "",
										"primary_image": "https://cdn1.ozone.ru/s3/multimedia-y/6077810038.jpg",
										"status": {
											"state": "price_sent",
											"state_failed": "",
											"moderate_status": "approved",
											"decline_reasons": [],
											"validation_state": "success",
											"state_name": "Продается",
											"state_description": "",
											"is_failed": false,
											"is_created": true,
											"state_tooltip": "",
											"item_errors": [],
											"state_updated_at": "2021-07-26T04:50:08.486697Z"
										}
							  		},
									{
										"id": 76723583,
										"name": "Онлайн-курс по дрессировке собак \"Собака: инструкция по применению. Одинокий волк\"",
										"offer_id": "23",
										"barcode": "",
										"buybox_price": "",
										"category_id": 90635895,
										"created_at": "2021-05-26T20:26:07.565586Z",
										"images": [],
										"marketing_price": "",
										"min_price": "",
										"old_price": "12200.0000",
										"premium_price": "5490.0000",
										"price": "6100.0000",
										"recommended_price": "",
										"sources": [
											{
												"is_enabled": true,
												"sku": 267684495,
												"source": "fbo"
											},
											{
												"is_enabled": true,
												"sku": 267684498,
												"source": "fbs"
											}
										],
										"state": "",
										"stocks": {
											"coming": 0,
											"present": 19,
											"reserved": 0
										},
										"errors": [],
										"updated_at": "2023-02-09T06:46:44.152Z",
										"vat": "0.0",
										"visible": true,
										"visibility_details": {
											"has_price": false,
											"has_stock": true,
											"active_product": false,
											"reasons": {}
										},
										"price_index": "0.00",
										"images360": [],
										"is_kgt": false,
										"color_image": "",
										"primary_image": "https://cdn1.ozone.ru/s3/multimedia-v/6062554531.jpg",
										"status": {
											"state": "price_sent",
											"state_failed": "",
											"moderate_status": "approved",
											"decline_reasons": [],
											"validation_state": "success",
											"state_name": "Продается",
											"state_description": "",
											"is_failed": false,
											"is_created": true,
											"state_tooltip": "",
											"item_errors": [],
											"state_updated_at": "2021-05-31T12:35:09.714641Z"
										}
									}
								]
						  	}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &info.ListRequest{
		OfferID: []string{
			"23",
			"010",
		},
		ProductID: make([]int64, 0),
		SKU:       make([]int64, 0),
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.ListResponse{
		Result: info.ListResponseResult{
			Items: []info.ListResponseResultItem{
				{
					ID:      78712196,
					Name:    "Как выбрать детские музыкальные инструменты. Ксилофон, бубен, маракасы и другие инструменты для детей до 6 лет. Мастер-класс о раннем музыкальном развитии от Монтессори-педагога",
					OfferID: "010",
					Barcode: "",
					Barcodes: []string{
						"2335900005",
						"7533900005",
					},
					BuyboxPrice:      "",
					CategoryID:       93726157,
					CreatedAt:        test.TimeFromString(t, time.RFC3339, "2021-06-03T03:40:05.871465Z"),
					Images:           make([]string, 0),
					CurrencyCode:     info.ListResponseResultItemCurrencyCodeRUB,
					MarketingPrice:   "",
					MinPrice:         "",
					OldPrice:         "1000.0000",
					PremiumPrice:     "590.0000",
					Price:            "690.0000",
					RecommendedPrice: "",
					Sources: []info.ListResponseResultItemSource{
						{
							IsEnabled: true,
							SKU:       269628393,
							Source:    "fbo",
						},
						{
							IsEnabled: true,
							SKU:       269628396,
							Source:    "fbs",
						},
					},
					HasDiscountedItem: true,
					IsDiscounted:      true,
					DiscountedStocks: info.ListResponseResultItemDiscountedStock{
						Coming:   0,
						Present:  0,
						Reserved: 0,
					},
					Stocks: info.ListResponseResultItemStocks{
						Coming:   0,
						Present:  13,
						Reserved: 0,
					},
					UpdatedAt: test.TimeFromString(t, time.RFC3339, "2023-02-09T06:46:44.152Z"),
					Vat:       "0.0",
					Visible:   true,
					VisibilityDetails: info.ListResponseResultItemVisibilityDetails{
						HasPrice:      false,
						HasStock:      true,
						ActiveProduct: false,
						Reasons: map[string]info.ListResponseResultItemVisibilityDetailsReasons{
							"field1": {
								Reasons: []info.ListResponseResultItemVisibilityDetailsReason{
									{
										Description: "descr",
										ID:          1,
									},
								},
							},
						},
					},
					PriceIndexes: info.ListResponseResultItemPriceIndexes{
						ExternalIndexData: info.ListResponseResultItemPriceIndexesExternalIndexData{
							MinimalPrice:         "string",
							MinimalPriceCurrency: "string",
							PriceIndexValue:      0,
						},
						OzonIndexData: info.ListResponseResultItemPriceIndexesOzonIndexData{
							MinimalPrice:         "string",
							MinimalPriceCurrency: "string",
							PriceIndexValue:      0,
						},
						PriceIndex: info.ListResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX,
						SelfMarketplacesIndexData: info.ListResponseResultItemPriceIndexesSelfMarketplacesIndexData{
							MinimalPrice:         "string",
							MinimalPriceCurrency: "string",
							PriceIndexValue:      0,
						},
					},
					Images360:    make([]string, 0),
					IsKGT:        false,
					ColorImage:   "",
					PrimaryImage: "https://cdn1.ozone.ru/s3/multimedia-y/6077810038.jpg",
					Status: info.ListResponseResultItemStatus{
						State:            "price_sent",
						StateFailed:      "",
						ModerateStatus:   "approved",
						DeclineReasons:   make([]string, 0),
						ValidationState:  "success",
						StateName:        "Продается",
						StateDescription: "",
						IsFailed:         false,
						IsCreated:        true,
						StateTooltip:     "",
						ItemErrors:       make([]info.ListResponseResultItemStatusItemError, 0),
						StateUpdatedAt:   test.TimeFromString(t, time.RFC3339, "2021-07-26T04:50:08.486697Z"),
					},
				},
				{
					ID:               76723583,
					Name:             "Онлайн-курс по дрессировке собак \"Собака: инструкция по применению. Одинокий волк\"",
					OfferID:          "23",
					Barcode:          "",
					BuyboxPrice:      "",
					CategoryID:       90635895,
					CreatedAt:        test.TimeFromString(t, time.RFC3339, "2021-05-26T20:26:07.565586Z"),
					Images:           make([]string, 0),
					MarketingPrice:   "",
					MinPrice:         "",
					OldPrice:         "12200.0000",
					PremiumPrice:     "5490.0000",
					Price:            "6100.0000",
					RecommendedPrice: "",
					Sources: []info.ListResponseResultItemSource{
						{
							IsEnabled: true,
							SKU:       267684495,
							Source:    "fbo",
						},
						{
							IsEnabled: true,
							SKU:       267684498,
							Source:    "fbs",
						},
					},
					Stocks: info.ListResponseResultItemStocks{
						Coming:   0,
						Present:  19,
						Reserved: 0,
					},
					UpdatedAt: test.TimeFromString(t, time.RFC3339, "2023-02-09T06:46:44.152Z"),
					Vat:       "0.0",
					Visible:   true,
					VisibilityDetails: info.ListResponseResultItemVisibilityDetails{
						HasPrice:      false,
						HasStock:      true,
						ActiveProduct: false,
						Reasons:       map[string]info.ListResponseResultItemVisibilityDetailsReasons{},
					},
					PriceIndex:   "0.00",
					Images360:    make([]string, 0),
					IsKGT:        false,
					ColorImage:   "",
					PrimaryImage: "https://cdn1.ozone.ru/s3/multimedia-v/6062554531.jpg",
					Status: info.ListResponseResultItemStatus{
						State:            "price_sent",
						StateFailed:      "",
						ModerateStatus:   "approved",
						DeclineReasons:   make([]string, 0),
						ValidationState:  "success",
						StateName:        "Продается",
						StateDescription: "",
						IsFailed:         false,
						IsCreated:        true,
						StateTooltip:     "",
						ItemErrors:       make([]info.ListResponseResultItemStatusItemError, 0),
						StateUpdatedAt:   test.TimeFromString(t, time.RFC3339, "2021-05-31T12:35:09.714641Z"),
					},
				},
			},
		},
	}, resp)
}
