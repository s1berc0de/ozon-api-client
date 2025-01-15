package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/andmetoo/ozon-api-client/ozon/product/v4/info"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestLimit_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v4/product/info/limit", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"daily_create": {
								"limit": 0,
								"reset_at": "2019-08-24T14:15:22Z",
								"usage": 0
							},
							"daily_update": {
								"limit": 0,
								"reset_at": "2019-08-24T14:15:22Z",
								"usage": 0
							},
							"total": {
								"limit": 0,
								"usage": 0
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v4/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Limit(context.Background(), &info.LimitRequest{})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.LimitResponse{
		DailyCreate: info.LimitResponseDailyCreate{
			Limit:   0,
			ResetAt: test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
			Usage:   0,
		},
		DailyUpdate: info.LimitResponseDailyUpdate{
			Limit:   0,
			ResetAt: test.TimeFromString(t, time.RFC3339, "2019-08-24T14:15:22Z"),
			Usage:   0,
		},
		Total: info.LimitResponseTotal{
			Limit: 0,
			Usage: 0,
		},
	}, resp)
}

func TestPrices_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v4/product/info/prices", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"offer_id":["356792"],"product_id":["243686911"],"visibility":"ALL"},"last_id":"","limit":100}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"items": [
									{
										"acquiring": 0,
										"product_id": 243686911,
										"offer_id": "356792",
										"price": {
											"currency_code": "RUB",
											"price": "499.0000",
											"old_price": "579.0000",
											"premium_price": "",
											"recommended_price": "",
											"retail_price": "",
											"vat": "0.200000",
											"min_ozon_price": "",
											"marketing_price": "",
											"marketing_seller_price": "",
											"auto_action_enabled": true
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
										"commissions": {
											"sales_percent": 15,
											"sales_percent_fbo": 15,
											"sales_percent_fbs": 0,
											"fbo_fulfillment_amount": 0,
											"fbo_direct_flow_trans_min_amount": 31,
											"fbo_direct_flow_trans_max_amount": 46.5,
											"fbo_deliv_to_customer_amount": 14.75,
											"fbo_return_flow_amount": 50,
											"fbo_return_flow_trans_min_amount": 21.7,
											"fbo_return_flow_trans_max_amount": 21.7,
											"fbs_first_mile_min_amount": 0,
											"fbs_first_mile_max_amount": 25,
											"fbs_direct_flow_trans_min_amount": 41,
											"fbs_direct_flow_trans_max_amount": 61.5,
											"fbs_deliv_to_customer_amount": 60,
											"fbs_return_flow_amount": 40,
											"fbs_return_flow_trans_min_amount": 41,
											"fbs_return_flow_trans_max_amount": 61.5
										},
										"marketing_actions": null,
										"volume_weight": 0
									}
								],
								"total": 1,
								"last_id": "ceVуbA=="
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v4/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Prices(context.Background(), &info.PricesRequest{
		Filter: info.PricesRequestFilter{
			OfferID: []string{
				"356792",
			},
			ProductID: []string{
				"243686911",
			},
			Visibility: info.PricesRequestFilterVisibilityALL,
		},
		LastID: "",
		Limit:  100,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.PricesResponse{
		Result: info.PricesResponseResult{
			Items: []info.PricesResponseResultItem{
				{
					Acquiring: 0,
					ProductID: 243686911,
					OfferID:   "356792",
					Price: info.PricesResponseResultItemPrice{
						CurrencyCode:         info.PricesResponseResultItemPriceCurrencyCodeRUB,
						Price:                "499.0000",
						OldPrice:             "579.0000",
						PremiumPrice:         "",
						RetailPrice:          "",
						Vat:                  "0.200000",
						MinOzonPrice:         "",
						MarketingPrice:       "",
						MarketingSellerPrice: "",
						AutoActionEnabled:    true,
					},
					PriceIndexes: info.PricesResponseResultItemPriceIndexes{
						ExternalIndexData: info.PricesResponseResultItemPriceIndexesExternalIndexData{
							MinimalPrice:         "string",
							MinimalPriceCurrency: "string",
							PriceIndexValue:      0,
						},
						OzonIndexData: info.PricesResponseResultItemPriceIndexesOzonIndexData{
							MinimalPrice:         "string",
							MinimalPriceCurrency: "string",
							PriceIndexValue:      0,
						},
						PriceIndex: info.PricesResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX,
						SelfMarketplacesIndexData: info.PricesResponseResultItemPriceIndexesSelfMarketplacesIndexData{
							MinimalPrice:         "string",
							MinimalPriceCurrency: "string",
							PriceIndexValue:      0,
						},
					},
					Commissions: info.PricesResponseResultItemCommissions{
						SalesPercent:                15,
						SalesPercentFBO:             15,
						SalesPercentFBS:             0,
						FBOFulfillmentAmount:        0,
						FBODirectFlowTransMinAmount: 31,
						FBODirectFlowTransMaxAmount: 46.5,
						FBODelivToCustomerAmount:    14.75,
						FBOReturnFlowAmount:         50,
						FBOReturnFlowTransMinAmount: 21.7,
						FBOReturnFlowTransMaxAmount: 21.7,
						FBSFirstMileMinAmount:       0,
						FBSFirstMileMaxAmount:       25,
						FBSDirectFlowTransMinAmount: 41,
						FBSDirectFlowTransMaxAmount: 61.5,
						FBSDelivToCustomerAmount:    60,
						FBSReturnFlowAmount:         40,
						FBSReturnFlowTransMinAmount: 41,
						FBSReturnFlowTransMaxAmount: 61.5,
					},
					MarketingActions: nil,
					VolumeWeight:     0,
				},
			},
			Total:  1,
			LastID: "ceVуbA==",
		},
	}, resp)
}

func TestStocks_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v4/product/info/stocks", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"cursor":"string","filter":{"offer_id":["string"],"product_id":["string"],"visibility":"ALL","with_quant":{"created":true,"exists":true}},"limit":0}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"cursor": "string",
							"items": [
								{
									"offer_id": "string",
									"product_id": 0,
									"stocks": [
										{
											"present": 0,
											"reserved": 0,
											"shipment_type": "SHIPMENT_TYPE_GENERAL",
											"sku": 0,
											"type": "string"
										}
									]
								}
							],
							"total": 0
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v4/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Stocks(context.Background(), &info.StocksRequest{
		Cursor: "string",
		Filter: info.StocksRequestFilter{
			OfferID: []string{
				"string",
			},
			ProductID: []string{
				"string",
			},
			Visibility: info.StocksRequestFilterVisibilityALL,
			WithQuant: info.StocksRequestFilterWithQuant{
				Created: true,
				Exists:  true,
			},
		},
		Limit: 0,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.StocksResponse{
		Cursor: "string",
		Items: []info.StocksResponseItem{
			{
				OfferID:   "string",
				ProductID: 0,
				Stocks: []info.StocksResponseItemStock{
					{
						Present:      0,
						Reserved:     0,
						ShipmentType: info.StocksResponseItemStockShipmentTypeSHIPMENTTYPEGENERAL,
						SKU:          0,
						Type:         "string",
					},
				},
			},
		},
		Total: 0,
	}, resp)
}
