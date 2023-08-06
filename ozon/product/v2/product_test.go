package v2_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	v2 "github.com/diphantxm/ozon-api-client/ozon/product/v2"

	"github.com/diphantxm/ozon-api-client/internal/auth"
	"github.com/diphantxm/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestImport_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/product/import", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"items":[{"attributes":[{"complex_id":0,"id":5076,"values":[{"dictionary_value_id":971082156,"value":"Стойка для акустической системы"}]},{"complex_id":0,"id":9048,"values":[{"value":"Комплект защитных плёнок для X3 NFC. Темный хлопок"}]},{"complex_id":0,"id":8229,"values":[{"dictionary_value_id":95911,"value":"Комплект защитных плёнок для X3 NFC. Темный хлопок"}]},{"complex_id":0,"id":85,"values":[{"dictionary_value_id":5060050,"value":"Samsung"}]},{"complex_id":0,"id":10096,"values":[{"dictionary_value_id":61576,"value":"серый"}]}],"barcode":"112772873170","category_id":17033876,"color_image":"","complex_attributes":[],"currency_code":"RUB","depth":10,"dimension_unit":"mm","height":250,"images":[],"images360":[],"name":"Комплект защитных плёнок для X3 NFC. Темный хлопок","offer_id":"143210608","old_price":"1100","pdf_list":[],"premium_price":"900","price":"1000","primary_image":"","vat":"0.1","weight":100,"weight_unit":"g","width":150}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"task_id": 172549793
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Import(context.Background(), &v2.ImportRequest{
		Items: []v2.ImportRequestItem{
			{
				Attributes: []v2.ImportRequestItemAttribute{
					{
						ID: 5076,
						Values: []v2.ImportRequestItemAttributeValue{
							{
								DictionaryValueID: 971082156,
								Value:             "Стойка для акустической системы",
							},
						},
					},
					{
						ID: 9048,
						Values: []v2.ImportRequestItemAttributeValue{
							{
								Value: "Комплект защитных плёнок для X3 NFC. Темный хлопок",
							},
						},
					},
					{
						ID: 8229,
						Values: []v2.ImportRequestItemAttributeValue{
							{
								DictionaryValueID: 95911,
								Value:             "Комплект защитных плёнок для X3 NFC. Темный хлопок",
							},
						},
					},
					{
						ID: 85,
						Values: []v2.ImportRequestItemAttributeValue{
							{
								DictionaryValueID: 5060050,
								Value:             "Samsung",
							},
						},
					},
					{
						ID: 10096,
						Values: []v2.ImportRequestItemAttributeValue{
							{
								DictionaryValueID: 61576,
								Value:             "серый",
							},
						},
					},
				},
				Barcode:           "112772873170",
				CategoryID:        17033876,
				ComplexAttributes: make([]v2.ImportRequestItemComplexAttribute, 0),
				Depth:             10,
				DimensionUnit:     v2.ImportRequestItemDimensionUnitMm,
				Height:            250,
				Images:            make([]string, 0),
				Images360:         make([]string, 0),
				Name:              "Комплект защитных плёнок для X3 NFC. Темный хлопок",
				OfferID:           "143210608",
				OldPrice:          "1100",
				PDFList:           make([]v2.ImportRequestItemPDFItem, 0),
				PremiumPrice:      "900",
				Price:             "1000",
				Vat:               "0.1",
				Weight:            100,
				WeightUnit:        v2.ImportRequestItemWeightUnitG,
				Width:             150,
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.ImportResponse{
		Result: v2.ImportResponseResult{
			TaskID: 172549793,
		},
	}, resp)
}

func TestList_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/product/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"filter":{"offer_id":["136748"],"product_id":["223681945"],"visibility":"ALL"},"last_id":"","limit":100}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
					  		"result": {
								"items": [
						  			{
										"product_id": 223681945,
										"offer_id": "136748"
						  			}
								],
								"total": 1,
								"last_id": "bnVсbA=="
					  		}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &v2.ListRequest{
		Filter: v2.ListRequestFilter{
			OfferID: []string{
				"136748",
			},
			ProductID: []string{
				"223681945",
			},
		},
		Limit: 100,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.ListResponse{
		Result: v2.ListResponseResult{
			Items: []v2.ListResponseResultItem{
				{
					ProductID: 223681945,
					OfferID:   "136748",
				},
			},
			Total:  1,
			LastID: "bnVсbA==",
		},
	}, resp)
}

func TestInfo_Success(t *testing.T) {
	c := v2.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v2/product/info", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"offer_id":"","product_id":137208233,"sku":0}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  	"result": {
								"id": 137208233,
								"name": "Комплект защитных плёнок для X3 NFC. Темный хлопок",
								"offer_id": "143210586",
								"barcode": "",
								"barcodes": [
								  	"2335900005",
								  	"7533900005"
								],
								"buybox_price": "",
								"category_id": 17038062,
								"created_at": "2021-10-21T15:48:03.529178Z",
								"images": [
							  		"https://cdn1.ozone.ru/s3/multimedia-5/6088931525.jpg",
							  		"https://cdn1.ozone.ru/s3/multimedia-p/6088915813.jpg"
								],
								"has_discounted_item": true,
								"is_discounted": true,
								"discounted_stocks": {
								  	"coming": 0,
								  	"present": 0,
								  	"reserved": 0
								},
								"currency_code": "RUB",
								"marketing_price": "",
								"min_price": "",
								"old_price": "",
								"premium_price": "",
								"price": "590.0000",
								"recommended_price": "",
								"sources": [
								  	{
										"is_enabled": true,
										"sku": 522759607,
										"source": "fbo"
								  	},
								  	{
										"is_enabled": true,
										"sku": 522759608,
										"source": "fbs"
								  	}
								],
								"stocks": {
								  	"coming": 0,
								  	"present": 0,
								  	"reserved": 0
								},
								"errors": [],
								"updated_at": "2023-02-09T06:46:44.152Z",
								"vat": "0.0",
								"visible": false,
								"visibility_details": {
								  	"has_price": true,
								  	"has_stock": false,
								 	"active_product": false
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
								"commissions": [],
								"volume_weight": 0.1,
								"is_prepayment": false,
								"is_prepayment_allowed": true,
								"images360": [],
								"is_kgt": false,
								"color_image": "",
								"primary_image": "https://cdn1.ozone.ru/s3/multimedia-p/6088931545.jpg",
								"status": {
									"state": "imported",
									"state_failed": "imported",
									"moderate_status": "",
							  		"decline_reasons": [],
							  		"validation_state": "pending",
							  		"state_name": "Не продается",
								  	"state_description": "Не создан",
								  	"is_failed": true,
								  	"is_created": false,
								  	"state_tooltip": "",
							  		"item_errors": [
										{
											"code": "error_attribute_values_empty",
											"field": "attribute",
											"attribute_id": 9048,
											"state": "imported",
											"level": "error",
											"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
											"optional_description_elements": {},
											"attribute_name": "Название модели"
										},
										{
											"code": "error_attribute_values_empty",
											"field": "attribute",
											"attribute_id": 5076,
											"state": "imported",
											"level": "error",
											"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
											"optional_description_elements": {},
											"attribute_name": "Рекомендовано для"
										},
										{
											"code": "error_attribute_values_empty",
											"field": "attribute",
											"attribute_id": 8229,
											"state": "imported",
											"level": "error",
											"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
											"optional_description_elements": {},
											"attribute_name": "Тип"
										},
										{
											"code": "error_attribute_values_empty",
											"field": "attribute",
											"attribute_id": 85,
											"state": "imported",
											"level": "error",
											"description": "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
											"optional_description_elements": {},
											"attribute_name": "Бренд"
										}
							  		],
							  		"state_updated_at": "2021-10-21T15:48:03.927309Z"
								}
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v2/product",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Info(context.Background(), &v2.InfoRequest{
		ProductID: 137208233,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &v2.InfoResponse{
		Result: v2.InfoResponseResult{
			ID:      137208233,
			Name:    "Комплект защитных плёнок для X3 NFC. Темный хлопок",
			OfferID: "143210586",
			Barcode: "",
			Barcodes: []string{
				"2335900005",
				"7533900005",
			},
			BuyboxPrice: "",
			CategoryID:  17038062,
			CreatedAt:   test.TimeFromString(t, time.RFC3339, "2021-10-21T15:48:03.529178Z"),
			Images: []string{
				"https://cdn1.ozone.ru/s3/multimedia-5/6088931525.jpg",
				"https://cdn1.ozone.ru/s3/multimedia-p/6088915813.jpg",
			},
			HasDiscountedItem: true,
			IsDiscounted:      true,
			DiscountedStocks: v2.InfoResponseResultDiscountedStocks{
				Coming:   0,
				Present:  0,
				Reserved: 0,
			},
			CurrencyCode:     v2.InfoResponseResultCurrencyCodeRUB,
			MarketingPrice:   "",
			MinPrice:         "",
			OldPrice:         "",
			PremiumPrice:     "",
			Price:            "590.0000",
			RecommendedPrice: "",
			Sources: []v2.InfoResponseResultSource{
				{
					IsEnabled: true,
					SKU:       522759607,
					Source:    "fbo",
				},
				{
					IsEnabled: true,
					SKU:       522759608,
					Source:    "fbs",
				},
			},
			Stocks: v2.InfoResponseResultStocks{
				Coming:   0,
				Present:  0,
				Reserved: 0,
			},
			UpdatedAt: test.TimeFromString(t, time.RFC3339, "2023-02-09T06:46:44.152Z"),
			Vat:       "0.0",
			Visible:   false,
			VisibilityDetails: v2.InfoResponseResultVisibilityDetails{
				HasPrice:      true,
				HasStock:      false,
				ActiveProduct: false,
			},
			PriceIndexes: v2.InfoResponseResultPriceIndexes{
				ExternalIndexData: v2.InfoResponseResultPriceIndexesExternalIndexData{
					MinimalPrice:         "string",
					MinimalPriceCurrency: "string",
					PriceIndexValue:      0,
				},
				OzonIndexData: v2.InfoResponseResultPriceIndexesOzonIndexData{
					MinimalPrice:         "string",
					MinimalPriceCurrency: "string",
					PriceIndexValue:      0,
				},
				PriceIndex: v2.InfoResponseResultPriceIndexesPriceIndexWITHOUTINDEX,
				SelfMarketplacesIndexData: v2.InfoResponseResultPriceIndexesSelfMarketplacesIndexData{
					MinimalPrice:         "string",
					MinimalPriceCurrency: "string",
					PriceIndexValue:      0,
				},
			},
			Commissions:         make([]v2.InfoResponseResultCommission, 0),
			VolumeWeight:        0.1,
			IsPrepayment:        false,
			IsPrepaymentAllowed: true,
			Images360:           make([]string, 0),
			ColorImage:          "",
			PrimaryImage:        "https://cdn1.ozone.ru/s3/multimedia-p/6088931545.jpg",
			Status: v2.InfoResponseResultStatus{
				State:            "imported",
				StateFailed:      "imported",
				ModerateStatus:   "",
				DeclineReasons:   make([]string, 0),
				ValidationState:  "pending",
				StateName:        "Не продается",
				StateDescription: "Не создан",
				IsFailed:         true,
				IsCreated:        false,
				StateTooltip:     "",
				ItemErrors: []v2.InfoResponseResultStatusItemError{
					{
						Code:                        "error_attribute_values_empty",
						Field:                       "attribute",
						AttributeID:                 9048,
						State:                       "imported",
						Level:                       "error",
						Description:                 "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						OptionalDescriptionElements: v2.InfoResponseResultStatusItemErrorOptionalDescriptionElements{},
						AttributeName:               "Название модели",
					},
					{
						Code:                        "error_attribute_values_empty",
						Field:                       "attribute",
						AttributeID:                 5076,
						State:                       "imported",
						Level:                       "error",
						Description:                 "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						OptionalDescriptionElements: v2.InfoResponseResultStatusItemErrorOptionalDescriptionElements{},
						AttributeName:               "Рекомендовано для",
					},
					{
						Code:                        "error_attribute_values_empty",
						Field:                       "attribute",
						AttributeID:                 8229,
						State:                       "imported",
						Level:                       "error",
						Description:                 "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						OptionalDescriptionElements: v2.InfoResponseResultStatusItemErrorOptionalDescriptionElements{},
						AttributeName:               "Тип",
					},
					{
						Code:                        "error_attribute_values_empty",
						Field:                       "attribute",
						AttributeID:                 85,
						State:                       "imported",
						Level:                       "error",
						Description:                 "Не заполнен обязательный атрибут. Иногда мы обновляем обязательные атрибуты или добавляем новые. Отредактируйте товар или загрузите новый XLS-шаблон с актуальными атрибутами. ",
						OptionalDescriptionElements: v2.InfoResponseResultStatusItemErrorOptionalDescriptionElements{},
						AttributeName:               "Бренд",
					},
				},
				StateUpdatedAt: test.TimeFromString(t, time.RFC3339, "2021-10-21T15:48:03.927309Z"),
			},
		},
	}, resp)
}
